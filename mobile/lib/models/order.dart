class Delivery {
  final String name;
  final String phone;
  final String zip;
  final String city;
  final String address;
  final String region;
  final String email;

  Delivery({
    required this.name,
    required this.phone,
    required this.zip,
    required this.city,
    required this.address,
    required this.region,
    required this.email,
  });

  factory Delivery.fromJson(Map<String, dynamic> json) {
    return Delivery(
      name: json['name'],
      phone: json['phone'],
      zip: json['zip'],
      city: json['city'],
      address: json['address'],
      region: json['region'],
      email: json['email'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'phone': phone,
      'zip': zip,
      'city': city,
      'address': address,
      'region': region,
      'email': email,
    };
  }
}

class Payment {
  final String transaction;
  final String? requestId;
  final String currency;
  final String provider;
  final int amount;
  final int paymentDt;
  final String bank;
  final int deliveryCost;
  final int goodsTotal;
  final int customFee;

  Payment({
    required this.transaction,
    this.requestId,
    required this.currency,
    required this.provider,
    required this.amount,
    required this.paymentDt,
    required this.bank,
    required this.deliveryCost,
    required this.goodsTotal,
    required this.customFee,
  });

  factory Payment.fromJson(Map<String, dynamic> json) {
    return Payment(
      transaction: json['transaction'],
      requestId: json['request_id'],
      currency: json['currency'],
      provider: json['provider'],
      amount: json['amount'],
      paymentDt: json['payment_dt'],
      bank: json['bank'],
      deliveryCost: json['delivery_cost'],
      goodsTotal: json['goods_total'],
      customFee: json['custom_fee'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'transaction': transaction,
      'request_id': requestId,
      'currency': currency,
      'provider': provider,
      'amount': amount,
      'payment_dt': paymentDt,
      'bank': bank,
      'delivery_cost': deliveryCost,
      'goods_total': goodsTotal,
      'custom_fee': customFee,
    };
  }
}

class Item {
  final int chrtId;
  final String trackNumber;
  final int price;
  final String rid;
  final String name;
  final int sale;
  final String size;
  final int totalPrice;
  final int nmId;
  final String brand;
  final int status;

  Item({
    required this.chrtId,
    required this.trackNumber,
    required this.price,
    required this.rid,
    required this.name,
    required this.sale,
    required this.size,
    required this.totalPrice,
    required this.nmId,
    required this.brand,
    required this.status,
  });

  factory Item.fromJson(Map<String, dynamic> json) {
    return Item(
      chrtId: json['chrt_id'],
      trackNumber: json['track_number'],
      price: json['price'],
      rid: json['rid'],
      name: json['name'],
      sale: json['sale'],
      size: json['size'],
      totalPrice: json['total_price'],
      nmId: json['nm_id'],
      brand: json['brand'],
      status: json['status'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'chrt_id': chrtId,
      'track_number': trackNumber,
      'price': price,
      'rid': rid,
      'name': name,
      'sale': sale,
      'size': size,
      'total_price': totalPrice,
      'nm_id': nmId,
      'brand': brand,
      'status': status,
    };
  }
}

class Order {
  final String orderUid;
  final String trackNumber;
  final String entry;
  final Delivery delivery;
  final Payment payment;
  final List<Item> items;
  final String locale;
  final String? internalSignature;
  final String customerId;
  final String deliveryService;
  final String shardkey;
  final int smId;
  final String dateCreated;
  final String oofShard;

  Order({
    required this.orderUid,
    required this.trackNumber,
    required this.entry,
    required this.delivery,
    required this.payment,
    required this.items,
    required this.locale,
    this.internalSignature,
    required this.customerId,
    required this.deliveryService,
    required this.shardkey,
    required this.smId,
    required this.dateCreated,
    required this.oofShard,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    return Order(
      orderUid: json['order_uid'],
      trackNumber: json['track_number'],
      entry: json['entry'],
      delivery: Delivery.fromJson(json['delivery']),
      payment: Payment.fromJson(json['payment']),
      items: (json['items'] as List).map((item) => Item.fromJson(item)).toList(),
      locale: json['locale'],
      internalSignature: json['internal_signature'],
      customerId: json['customer_id'],
      deliveryService: json['delivery_service'],
      shardkey: json['shardkey'],
      smId: json['sm_id'],
      dateCreated: json['date_created'],
      oofShard: json['oof_shard'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'order_uid': orderUid,
      'track_number': trackNumber,
      'entry': entry,
      'delivery': delivery.toJson(),
      'payment': payment.toJson(),
      'items': items.map((item) => item.toJson()).toList(),
      'locale': locale,
      'internal_signature': internalSignature,
      'customer_id': customerId,
      'delivery_service': deliveryService,
      'shardkey': shardkey,
      'sm_id': smId,
      'date_created': dateCreated,
      'oof_shard': oofShard,
    };
  }
}