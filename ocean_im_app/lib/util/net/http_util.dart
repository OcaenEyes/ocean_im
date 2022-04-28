import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:ocean_im/util/net/dio_util.dart';

typedef Success<T> = Function(T data);
typedef Fail<T> = Function(T data);

class HttpUtil {
  //GET
  void getRequest<T>(
    baseUrl,
    path,
    parameters, {
    Success? success,
    Fail? fail,
  }) {
    DioUtils _dioUtils = DioUtils();
    _dioUtils.get(Method.get, baseUrl + path, parameters, success: (result) {
      if (result['code'] == 0) {
        success!(result);
      } else {
        EasyLoading.showToast(result['message'].toString(),
            duration: const Duration(milliseconds: 500));
      }
    }, fail: (e) {
      fail!(e);
    });
  }

  //POST
  void postRequest<T>(
    baseUrl,
    path,
    parameters, {
    Success? success,
    Fail? fail,
  }) {
    DioUtils _dioUtils = DioUtils();
    _dioUtils.request(Method.post, baseUrl + path, parameters,
        success: (result) {
      if (result['code'] == 0) {
        success!(result);
      } else {
        EasyLoading.showToast(result['message'],
            duration: const Duration(milliseconds: 500));
      }
    }, fail: (e) {
      fail!(e);
    });
  }
}
