import 'package:flutter/material.dart';
import 'package:ocean_im/model/user_info_model.dart';

class UserInfoNotifier extends ChangeNotifier {
  late UserInfo userInfo;
  UserInfo? get info => userInfo;
  void setInfo(info) {
    userInfo = UserInfo.fromJson(info);
    notifyListeners();
  }
}
