# newsflow

Сервис рассылки новостей на Go и Kafka. Новости приходят через Kafka, пользователь сам выбирает категории и каналы доставки (email / push / sms).

## Архитектура

```
news.incoming -> Dispatcher -> news.to_deliver -> Delivery Worker -> notification_history
                     |
                PostgreSQL (подписки пользователей)
Preferences API (Gin) управляет подписками
```

## Стек
Go, Apache Kafka (KRaft), PostgreSQL, Gin, Docker, GitHub Actions.

## Быстрый старт
```bash
make up        # поднять Kafka
make consumer  # терминал 1
make producer  # терминал 2
```

## Roadmap
- [x] Kafka в Docker, producer/consumer
- [ ] Схема PostgreSQL + миграции
- [ ] Preferences API (Gin)
- [ ] Dispatcher (fan-out по подписчикам категории)
- [ ] Delivery Worker (моки каналов, история)
- [ ] Retry + dead-letter топик
- [ ] Dockerfile для сервисов, CI/CD, Kubernetes
- [ ] Метрики Prometheus + Grafana
