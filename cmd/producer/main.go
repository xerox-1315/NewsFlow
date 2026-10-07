package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/YOUR_USERNAME/newsflow/internal/news"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    news.TopicIncoming,
		Balancer: &kafka.Hash{}, // одна категория -> одна партиция
	}
	defer w.Close()

	n := news.News{
		ID:          "demo-1",
		Category:    "tech",
		Title:       "Вышел новый процессор",
		Body:        "Тестовая новость для проверки Kafka",
		PublishedAt: time.Now().UTC(),
	}
	payload, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := w.WriteMessages(ctx, kafka.Message{Key: []byte(n.Category), Value: payload}); err != nil {
		log.Fatalf("send failed: %v", err)
	}
	log.Println("новость отправлена в", news.TopicIncoming)
}
