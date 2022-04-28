import 'package:fluro/fluro.dart';
import 'package:flutter/cupertino.dart';

typedef HandlerFunc = Widget Function(
    BuildContext? context, Map<String, List<String>> parameters);

typedef PageBuilderFunc = Widget Function(BundleArguments bundleArguments);

class BundleArguments {
  final Map<String, dynamic> _map = {};
  _setValue(var k, var v) => _map[k] = v;

  _getValue(String k) {
    if (!_map.containsKey(k)) {
      throw Exception("你使用的$k 不存在");
    }
    return _map[k];
  }

  putInt(String k, int v) => _map[k] = v;
  putString(String k, String v) => _setValue(k, v);
  putBool(String k, bool v) => _setValue(k, v);
  putList<V>(String k, List<V> v) => _setValue(k, v);
  putMap<K, V>(String k, Map<K, V> v) => _setValue(k, v);

  int getInt(String k) => _getValue(k) as int;
  String getString(String k) => _getValue(k) as String;
  bool getBool(String k) => _getValue(k) as bool;
  List getList(String k) => _getValue(k) as List;
  Map getMap(String k) => _getValue(k) as Map;

  @override
  String toString() {
    return _map.toString();
  }
}

class PageBuilder {
  PageBuilderFunc pageBuilderFunc;
  late HandlerFunc handlerFunc;

  PageBuilder({required this.pageBuilderFunc}) {
    handlerFunc = (context, _) => pageBuilderFunc(
        ModalRoute.of(context!)?.settings.arguments as BundleArguments);
  }
  Handler getHandler() {
    return Handler(handlerFunc: handlerFunc);
  }
}
