// 03-notification-system: 完整的实时通知系统
// 演示如何构建一个生产级的用户通知系统
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/tmaxmax/go-sse"
)

// Notification 通知结构
type Notification struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

// NotificationServer 通知服务器
type NotificationServer struct {
	sse     *sse.Server
	storage *NotificationStorage
}

// NotificationStorage 内存存储
type NotificationStorage struct {
	mu    sync.RWMutex
	items map[string]*Notification
}

func NewNotificationStorage() *NotificationStorage {
	return &NotificationStorage{
		items: make(map[string]*Notification),
	}
}

func (ns *NotificationStorage) Add(n *Notification) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.items[n.ID] = n
}

func (ns *NotificationStorage) Get(id string) (*Notification, bool) {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	n, ok := ns.items[id]
	return n, ok
}

func NewNotificationServer() *NotificationServer {
	// 创建重放器：保留最近 1000 条通知
	replayer, err := sse.NewFiniteReplayer(1000, true)
	if err != nil {
		panic(err)
	}

	joe := &sse.Joe{
		Replayer: replayer,
	}

	storage := NewNotificationStorage()

	s := &NotificationServer{
		sse: &sse.Server{
			Provider: joe,
			OnSession: func(session *sse.Session) (sse.Subscription, bool) {
				// 从请求中获取用户信息
				userID := session.Req.URL.Query().Get("user_id")
				if userID == "" {
					// 未授权用户
					return sse.Subscription{}, false
				}

				log.Printf("用户 %s 连接，LastEventID: %s", userID, session.LastEventID)

				// 订阅用户专属主题和全局通知
				return sse.Subscription{
					Topics:      []string{"global", "user:" + userID},
					LastEventID: session.LastEventID,
					Client:      session.Res,
				}, true
			},
		},
		storage: storage,
	}

	// 启动模拟通知
	go s.simulateNotifications()

	return s
}

func (s *NotificationServer) simulateNotifications() {
	types := []string{"info", "warning", "success", "error"}
	titles := []string{
		"系统更新",
		"新消息",
		"任务完成",
		"安全提醒",
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	counter := 0
	for range ticker.C {
		counter++
		notif := &Notification{
			ID:      fmt.Sprintf("notif-%d", counter),
			Type:    types[counter%len(types)],
			Title:   titles[counter%len(titles)],
			Content: fmt.Sprintf("这是第 %d 条通知", counter),
			Time:    time.Now(),
		}

		s.storage.Add(notif)
		s.Broadcast(notif, "global")
		log.Printf("[广播] %s: %s", notif.Title, notif.Content)
	}
}

func (s *NotificationServer) Broadcast(notif *Notification, topic string) {
	data, _ := json.Marshal(notif)

	msg := &sse.Message{}
	msg.ID = sse.ID(notif.ID)
	msg.Type = sse.Type(notif.Type)
	msg.AppendData(string(data))

	s.sse.Publish(msg, topic)
}

func (s *NotificationServer) SendToUser(userID string, notif *Notification) {
	data, _ := json.Marshal(notif)

	msg := &sse.Message{}
	msg.ID = sse.ID(notif.ID)
	msg.Type = sse.Type(notif.Type)
	msg.AppendData(string(data))

	s.sse.Publish(msg, "user:"+userID)
	log.Printf("[用户通知] %s: %s", userID, notif.Title)
}

func (s *NotificationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.sse.ServeHTTP(w, r)
}

func main() {
	server := NewNotificationServer()

	// 设置路由
	mux := http.NewServeMux()
	mux.Handle("/events", server)
	mux.HandleFunc("/notify", handleNotify(server))
	mux.HandleFunc("/", indexHandler)

	log.Println("========================================")
	log.Println("   实时通知系统启动在 :8082")
	log.Println("========================================")
	log.Println("")
	log.Println("连接地址:")
	log.Println("  curl -N 'http://localhost:8082/events?user_id=alice'")
	log.Println("")
	log.Println("发送个人通知:")
	log.Println("  curl -X POST http://localhost:8082/notify \\")
	log.Println("    -H 'Content-Type: application/json' \\")
	log.Println("    -d '{\"user_id\":\"alice\",\"title\":\"新消息\",\"content\":\"你好，Alice！\"}'")
	log.Println("")
	log.Fatal(http.ListenAndServe(":8082", mux))
}

func handleNotify(server *NotificationServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			UserID  string `json:"user_id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		notif := &Notification{
			ID:      fmt.Sprintf("notif-%d", time.Now().UnixNano()),
			Type:    "info",
			Title:   req.Title,
			Content: req.Content,
			Time:    time.Now(),
		}

		server.SendToUser(req.UserID, notif)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"id":      notif.ID,
		})
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>SSE 通知系统演示</title>
    <style>
        body { font-family: system-ui; max-width: 800px; margin: 50px auto; padding: 20px; }
        .container { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; }
        .panel { border: 1px solid #ddd; padding: 20px; border-radius: 8px; }
        .panel h2 { margin-top: 0; }
        textarea { width: 100%; padding: 10px; margin-bottom: 10px; }
        button { padding: 10px 20px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
        button:hover { background: #0056b3; }
        #notifications { height: 400px; overflow-y: auto; border: 1px solid #ddd; padding: 10px; }
        .notification { padding: 10px; margin-bottom: 10px; border-radius: 4px; }
        .notification.info { background: #d1ecf1; }
        .notification.warning { background: #fff3cd; }
        .notification.success { background: #d4edda; }
        .notification.error { background: #f8d7da; }
    </style>
</head>
<body>
    <h1>🔔 SSE 实时通知系统演示</h1>
    <div class="container">
        <div class="panel">
            <h2>发送通知</h2>
            <input type="text" id="userId" placeholder="用户ID (如: alice)" value="alice">
            <input type="text" id="title" placeholder="标题" value="测试通知">
            <textarea id="content" rows="3" placeholder="内容">这是一条测试通知！</textarea>
            <button onclick="sendNotification()">发送通知</button>
        </div>
        <div class="panel">
            <h2>接收通知</h2>
            <input type="text" id="myUserId" placeholder="你的用户ID" value="alice">
            <button onclick="connect()">连接</button>
            <button onclick="disconnect()">断开</button>
            <div id="notifications"></div>
        </div>
    </div>
    <script>
        let eventSource = null;

        function connect() {
            const userId = document.getElementById('myUserId').value;
            const url = '/events?user_id=' + encodeURIComponent(userId);
            eventSource = new EventSource(url);

            eventSource.onmessage = (event) => {
                const data = JSON.parse(event.data);
                addNotification(data);
            };

            eventSource.onerror = (error) => {
                console.error('连接错误:', error);
            };

            document.getElementById('notifications').innerHTML += '<p class="notification info">已连接到服务器</p>';
        }

        function disconnect() {
            if (eventSource) {
                eventSource.close();
                eventSource = null;
                document.getElementById('notifications').innerHTML += '<p class="notification warning">已断开连接</p>';
            }
        }

        function sendNotification() {
            const userId = document.getElementById('userId').value;
            const title = document.getElementById('title').value;
            const content = document.getElementById('content').value;

            fetch('/notify', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ user_id: userId, title, content })
            }).then(r => r.json()).then(data => {
                alert('通知已发送: ' + data.id);
            });
        }

        function addNotification(data) {
            const div = document.createElement('div');
            div.className = 'notification ' + data.type;
            div.innerHTML = '<strong>' + data.title + '</strong><br>' + data.content + '<br><small>' + new Date(data.time).toLocaleString() + '</small>';
            document.getElementById('notifications').prepend(div);
        }

        // 自动连接
        setTimeout(connect, 500);
    </script>
</body>
</html>`))
}
