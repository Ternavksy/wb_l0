import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:animate_do/animate_do.dart';
import 'models/order.dart';

void main() {
  runApp(const OrderLookupApp());
}

class OrderLookupApp extends StatelessWidget {
  const OrderLookupApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'WB_L0',
      theme: ThemeData(
        primaryColor: const Color(0xFF8A00FF), 
        scaffoldBackgroundColor: Colors.transparent,
        fontFamily: 'Roboto',
        textTheme: const TextTheme(
          bodyMedium: TextStyle(color: Colors.white),
        ),
        colorScheme: const ColorScheme.light(
          primary: Color(0xFF8A00FF), 
          secondary: Color(0xFFE6007A), 
          surface: Color(0xFFF5F5F5), 
          onPrimary: Colors.white,
          onSecondary: Colors.white,
          onSurface: Color(0xFF2D2D2D), 
        ),
      ),
      home: const OrderLookupScreen(),
    );
  }
}

class OrderLookupScreen extends StatefulWidget {
  const OrderLookupScreen({super.key});

  @override
  _OrderLookupScreenState createState() => _OrderLookupScreenState();
}

class _OrderLookupScreenState extends State<OrderLookupScreen> {
  final TextEditingController _controller = TextEditingController();
  Order? _order;
  String _errorMessage = '';
  bool _isLoading = false;

  Future<void> _fetchOrder(String orderId) async {
    setState(() {
      _isLoading = true;
      _order = null;
      _errorMessage = '';
    });

    try {
      final response = await http.get(Uri.parse('http://10.0.2.2:8080/order/$orderId'));
      if (response.statusCode == 200) {
        setState(() {
          _order = Order.fromJson(jsonDecode(response.body));
        });
      } else {
        setState(() {
          _errorMessage = 'Order not found';
        });
      }
    } catch (e) {
      setState(() {
        _errorMessage = 'Error: $e';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.bottomRight,
            colors: [Color(0xFF8A00FF), Color(0xFFE6007A)],
          ),
        ),
        child: SafeArea(
          child: Padding(
            padding: const EdgeInsets.all(20.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                FadeInDown(
                  duration: const Duration(milliseconds: 800),
                  child: const Text(
                    'WB L0',
                    style: TextStyle(
                      fontSize: 32,
                      fontWeight: FontWeight.bold,
                      color: Colors.white,
                    ),
                  ),
                ),
                const SizedBox(height: 20),
                FadeInUp(
                  duration: const Duration(milliseconds: 800),
                  child: _buildSearchField(),
                ),
                const SizedBox(height: 20),
                if (_isLoading)
                  Center(
                    child: CircularProgressIndicator(
                      valueColor: AlwaysStoppedAnimation<Color>(
                          Theme.of(context).colorScheme.secondary), 
                    ),
                  )
                else if (_errorMessage.isNotEmpty)
                  FadeIn(
                    duration: const Duration(milliseconds: 600),
                    child: Center(
                      child: Text(
                        _errorMessage,
                        style: const TextStyle(
                          color: Color(0xFFFF4D4D), 
                          fontSize: 18,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                    ),
                  )
                else if (_order != null)
                  Expanded(
                    child: FadeInUp(
                      duration: const Duration(milliseconds: 800),
                      child: _buildOrderDetails(),
                    ),
                  )
                else
                  const Expanded(
                    child: Center(
                      child: Text(
                        'Enter an order ID to search',
                        style: TextStyle(
                          color: Colors.white70,
                          fontSize: 18,
                        ),
                      ),
                    ),
                  ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildSearchField() {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white.withOpacity(0.15), 
        borderRadius: BorderRadius.circular(15),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.2),
            blurRadius: 10,
            offset: const Offset(5, 5),
          ),
          BoxShadow(
            color: Colors.white.withOpacity(0.1),
            blurRadius: 10,
            offset: const Offset(-5, -5),
          ),
        ],
      ),
      child: TextField(
        controller: _controller,
        style: const TextStyle(color: Colors.white),
        decoration: InputDecoration(
          hintText: 'Enter order UID',
          hintStyle: TextStyle(color: Colors.white.withOpacity(0.6)),
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(15),
            borderSide: BorderSide.none,
          ),
          contentPadding: const EdgeInsets.symmetric(horizontal: 20, vertical: 15),
          suffixIcon: IconButton(
            icon: Icon(Icons.search, color: Theme.of(context).colorScheme.secondary),
            onPressed: () => _fetchOrder(_controller.text),
          ),
        ),
        onSubmitted: (_) => _fetchOrder(_controller.text),
      ),
    );
  }

  Widget _buildOrderDetails() {
    return ListView(
      children: [
        _buildSectionCard('Order Details', _buildOrderInfo()),
        _buildSectionCard('Delivery Information', _buildDeliveryInfo()),
        _buildSectionCard('Payment Details', _buildPaymentInfo()),
        _buildSectionCard('Items', _buildItemsList()),
      ],
    );
  }

  Widget _buildSectionCard(String title, Widget content) {
    return FadeInUp(
      duration: const Duration(milliseconds: 600),
      child: Card(
        color: Colors.white.withOpacity(0.2), 
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(15)),
        elevation: 0,
        child: ExpansionTile(
          title: Text(
            title,
            style: TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
              color: Theme.of(context).colorScheme.onPrimary, 
            ),
          ),
          children: [
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: content,
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildOrderInfo() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        _buildInfoRow('Order UID', _order!.orderUid),
        _buildInfoRow('Track Number', _order!.trackNumber),
        _buildInfoRow('Entry', _order!.entry),
        _buildInfoRow('Customer ID', _order!.customerId),
        _buildInfoRow('Delivery Service', _order!.deliveryService),
        _buildInfoRow('Date Created', _order!.dateCreated),
      ],
    );
  }

  Widget _buildDeliveryInfo() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        _buildInfoRow('Name', _order!.delivery.name),
        _buildInfoRow('Phone', _order!.delivery.phone),
        _buildInfoRow('Email', _order!.delivery.email),
        _buildInfoRow('Address', _order!.delivery.address),
        _buildInfoRow('City', _order!.delivery.city),
        _buildInfoRow('Region', _order!.delivery.region),
        _buildInfoRow('Zip', _order!.delivery.zip),
      ],
    );
  }

  Widget _buildPaymentInfo() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        _buildInfoRow('Transaction', _order!.payment.transaction),
        _buildInfoRow('Provider', _order!.payment.provider),
        _buildInfoRow('Currency', _order!.payment.currency),
        _buildInfoRow('Amount', '\$${_order!.payment.amount / 100}'),
        _buildInfoRow('Bank', _order!.payment.bank),
        _buildInfoRow('Delivery Cost', '\$${_order!.payment.deliveryCost / 100}'),
        _buildInfoRow('Goods Total', '\$${_order!.payment.goodsTotal / 100}'),
      ],
    );
  }

  Widget _buildItemsList() {
    return Column(
      children: _order!.items.asMap().entries.map((entry) {
        final item = entry.value;
        return Padding(
          padding: const EdgeInsets.only(bottom: 8.0),
          child: Card(
            color: Colors.white.withOpacity(0.15), 
            shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
            child: ListTile(
              title: Text(
                item.name,
                style: TextStyle(
                  color: Theme.of(context).colorScheme.onPrimary, 
                  fontWeight: FontWeight.w600,
                ),
              ),
              subtitle: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text('Brand: ${item.brand}',
                      style: const TextStyle(color: Colors.white70)),
                  Text('Price: \$${item.price / 100}',
                      style: const TextStyle(color: Colors.white70)),
                  Text('Total: \$${item.totalPrice / 100}',
                      style: const TextStyle(color: Colors.white70)),
                ],
              ),
            ),
          ),
        );
      }).toList(),
    );
  }

  Widget _buildInfoRow(String label, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(
            label,
            style: const TextStyle(color: Colors.white70, fontSize: 16),
          ),
          Flexible(
            child: Text(
              value,
              style: TextStyle(
                color: Theme.of(context).colorScheme.onPrimary,
                fontSize: 16,
                fontWeight: FontWeight.w500,
              ),
              textAlign: TextAlign.right,
            ),
          ),
        ],
      ),
    );
  }
}