CREATE TABLE orders(
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50),
    "entry" VARCHAR(50), -- поставил в двойных ковычках, так как слово зарезервированно
    customer_id VARCHAR(50),
    date_created TIMESTAMP,
    locale VARCHAR(10),
    internal_signature VARCHAR(50),
    delivery_service VARCHAR(50),
    shardkey VARCHAR(10),
    sm_id INTEGER,
    oof_shard VARCHAR(10)
);

CREATE TABLE delivery(
    order_uid VARCHAR(50) PRIMARY KEY,
    email VARCHAR(100),
    region VARCHAR(100),
    "address" VARCHAR(200),
    city VARCHAR(50),
    zip VARCHAR(20),
    phone VARCHAR(20),
    "name" VARCHAR(100),
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);

CREATE TABLE payment(
    order_uid VARCHAR(50) PRIMARY KEY,
    transaction VARCHAR(50),
    request_id VARCHAR(50),
    currency VARCHAR(10),
    provider VARCHAR(50),
    amount INTEGER,
    payment_dt BIGINT,
    bank VARCHAR(50),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);

CREATE TABLE items(
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(50),
    chrt_id INTEGER,
    track_number VARCHAR(50),
    price INTEGER,
    rid VARCHAR(50),
    name VARCHAR(100),
    sale INTEGER,
    size VARCHAR(10),
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR(100),
    status INTEGER,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);

--Тестовые данные

INSERT INTO orders (order_uid, track_number, entry, customer_id, date_created, locale, delivery_service, shardkey, sm_id, oof_shard)
VALUES ('b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 'test', '2021-11-26T06:22:19Z', 'en', 'meest', '9', 99, '1');

INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
VALUES ('b563feb7b2b84b6test', 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');

INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES ('b563feb7b2b84b6test', 'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0);

INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES ('b563feb7b2b84b6test', 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);