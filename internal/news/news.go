// Package news описывает событие новости, которое ходит через Kafka.
package news

import "time"

const TopicIncoming = "news.incoming"

type News struct {
	ID          string    `json:"news_id"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	PublishedAt time.Time `json:"published_at"`
}
