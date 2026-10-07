package main

import (
	"context"
	"encoding/json"
	"log"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"

	"github.com/YOUR_USERNAME/newsflow/internal/news"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   news.TopicIncoming,
		GroupID: "newsflow-dispatcher",
	})
	defer r.Close()

	log.Println("consumer запущен, ждём новости...")
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println("остановка:", err)
			return
		}
		var n news.News
		if err := json.Unmarshal(msg.Value, &n); err != nil {
			log.Printf("битое сообщение offset=%d: %v", msg.Offset, err)
			continue
		}
		log.Printf("новость [%s] %q (partition=%d offset=%d)", n.Category, n.Title, msg.Partition, msg.Offset)
	}
}
