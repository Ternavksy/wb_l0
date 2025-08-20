Сервис поиска заказов (WB_L0)
Демо-проект для поиска заказов, включающий Go-бэкенд (Kafka, PostgreSQL, кэш) и мобильное приложение на Flutter
Структура проекта

backend/: Go-микросервис для обработки сообщений Kafka, хранения в PostgreSQL и HTTP API
mobile/: Flutter-приложение для запроса и отображения данных о заказах
videos/: Видео с демонстрацией работы

Основные функции
Бэкенд

Получение сообщений из Kafka (топик orders)
Сохранение заказов в PostgreSQL
Кэширование заказов в памяти
HTTP API (GET /order/{id})

Мобильное приложение

UI для ввода order_uid и отображения данных заказа

Тесты

Юнит-тесты для бэка (кэш, БД, server)

Установка и запуск

Клонировать репозиторий:
git clone https://github.com/Ternavksy/wb_l0.git
cd wb_l0


Настройка бэкенда:
cd backend
go mod tidy
docker-compose up -d
psql -U postgres -f schema.sql
go run main.go


Настройка мобильного приложения:
cd mobile
flutter pub get
flutter run



Демо-видео

Демо бэкенда (videos/backend, также можно посмотреть в гугл-диске)
Демо мобильного приложения (videos/mobile)
гугл диск: https://drive.google.com/drive/folders/1Q6C89I0aNOo2JXgQ3J6jGi3VEsMIb3qO?usp=drive_link

API

GET /order/{id}: Возвращает данные заказа в JSON.curl http://localhost:8080/order/test-order-145



Примечания

Для эмулятора Android используйте http://10.0.2.2:8080, для физического устройства — IP в локальной сети (например, http://192.168.1.33:8080)
