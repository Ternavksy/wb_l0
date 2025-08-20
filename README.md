# **WB_L0**

**Проект** для поиска заказов, включающий Go-бэкенд (Kafka, PostgreSQL, кэш) и мобильное приложение на Flutter.

## **Структура проекта**

- **backend/**: Go-микросервис для обработки сообщений Kafka, хранения в PostgreSQL и HTTP API.
- **mobile/**: Flutter-приложение для запроса и отображения данных о заказах.
- **videos/**: Видео с демонстрацией работы.

## **Основные функции**

### **Бэкенд**

- Получение сообщений из Kafka (топик `orders`).
- Сохранение заказов в PostgreSQL.
- Кэширование заказов в памяти.
- HTTP API (`GET /order/{id}`).

### **Мобильное приложение**

- UI для ввода `order_uid` и отображения данных заказа.

### **Тесты**

- Юнит-тесты для Go (кэш, БД, consumer, server).
- Юнит- и виджет-тесты для Flutter (модель, UI).

## **Установка и запуск**

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/Ternavksy/wb_l0.git
   cd wb_l0


2. Настройка бэкенда:
    ```bash
    cd backend
    go mod tidy
    docker-compose up -d
    psql -U postgres -f schema.sql
    go run main.go


3. Настройка мобильного приложения:
    ```bash
    cd mobile
    flutter pub get
    flutter run



## **Демо-видео**

- Демо бэкенда (videos/backend.mp4). Также ниже прикрепил ссылку на гугл диск.
- Демо мобильного приложения (videos/mobile.mp4)

## **API**

- GET /order/{id}: Возвращает данные заказа в JSON.curl http://localhost:8080/order/test-order-145



## **Примечания**

Для эмулятора Android используйте http://10.0.2.2:8080, для физического устройства — IP в локальной сети (например, http://192.168.1.33:8080).

## **Видео демонстрации работы**
- https://drive.google.com/drive/folders/1Q6C89I0aNOo2JXgQ3J6jGi3VEsMIb3qO?hl=ru