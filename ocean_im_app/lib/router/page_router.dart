import 'package:fluro/fluro.dart';
import 'package:ocean_im/router/page_routes.dart';

class PageRouter {
  static final router = FluroRouter();
  static setupRoutes() {
    pageRoutes.forEach((key, value) {
      router.define(key.toString(), handler: value.getHandler());
    });
  }
}
