import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:flutter/cupertino.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/util/error/errot_handler.dart';

const int connectTimeout = 15000; //15s
const int receiveTimeout = 15000;
const int sendTimeout = 10000;

typedef Success<T> = Function(T data);
typedef Fail<T> = Function(T data);

class DioUtils {
  late Dio _dio;
  // 创建 dio 实例对象
  Dio createInstance() {
    var options = BaseOptions(
      /// 请求的Content-Type，默认值是"application/json; charset=utf-8".
      /// 如果您想以"application/x-www-form-urlencoded"格式编码请求数据,
      /// 可以设置此选项为 `Headers.formUrlEncodedContentType`,  这样[Dio]就会自动编码请求体.
      /// contentType: Headers.formUrlEncodedContentType, // 适用于post form表单提交
      responseType: ResponseType.json,
      validateStatus: (status) {
        // 不使用http状态码判断状态，使用AdapterInterceptor来处理（适用于标准REST风格）
        return true;
      },
      headers: httpHeaders,
      connectTimeout: connectTimeout,
      receiveTimeout: receiveTimeout,
      sendTimeout: sendTimeout,
    );
    _dio = Dio(options);
    return _dio;
  }

  Future request<T>(Method method, String path, dynamic params,
      {Success? success, Fail? fail}) async {
    try {
      // 没有网络
      Dio _dio = createInstance();
      _dio.options.headers = {
        'Accept': 'application/json,*/*',
        'Content-Type': 'application/json',
        'Token': Global.token
      };
      Response response = await _dio.request(path,
          data: params, options: Options(method: methodValues[method]));
      // debugPrint(response.data.toString());
      if (success != null) {
        success(response.data);
      }
    } on DioError catch (e) {
      final NetError netError = ExceptionHandle.handleException(e);
      _onError(NetError(netError.code, netError.message), fail!);
    }
  }

  Future get<T>(Method method, String path, dynamic params,
      {Success? success, Fail? fail}) async {
    try {
      // 没有网络
      Dio _dio = createInstance();
      Response response = await _dio.get(path,
          queryParameters: params,
          options: Options(method: methodValues[method]));
      // debugPrint(response.data.toString());
      if (success != null) {
        success(response.data);
      }
    } on DioError catch (e) {
      final NetError netError = ExceptionHandle.handleException(e);
      _onError(NetError(netError.code, netError.message), fail!);
    }
  }
}

/// 自定义Header
Map<String, dynamic> httpHeaders = {
  'Accept': 'application/json,*/*',
  'Content-Type': 'application/json',
  'Token': Global.token
};

void _onError(NetError _netError, Fail fail) {
  // debugPrint('接口请求异常： code: $code, msg: $msg');
  fail(_netError);
}

Map<String, dynamic> parseData(String data) {
  return json.decode(data) as Map<String, dynamic>;
}

enum Method { get, post, delete, put, patch, head }
//使用：methodValues[Method.post]
const methodValues = {
  Method.get: "GET",
  Method.post: "POST",
  Method.delete: "DELETE",
  Method.put: "PUT",
  Method.patch: "PATCH",
  Method.head: "HEAD",
};
