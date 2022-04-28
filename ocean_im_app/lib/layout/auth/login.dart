import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/router/page_builder.dart';
import 'package:ocean_im/router/page_routes.dart';
import 'package:ocean_im/util/api/base_api.dart';
import 'package:ocean_im/util/db/hive_db_util.dart';
import 'package:ocean_im/util/error/errot_handler.dart';
import 'package:ocean_im/util/net/http_util.dart';
// import 'package:shared_preferences/shared_preferences.dart';

class Login extends StatefulWidget {
  final BundleArguments bundleArguments;
  const Login({Key? key, required this.bundleArguments}) : super(key: key);

  @override
  // State<StatefulWidget> createState() {
  //   return _LoginState();
  // }
  _LoginState createState() => _LoginState();
}

class _LoginState extends State<Login> {
  final userAccountController = TextEditingController();
  final passwordController = TextEditingController();
  HttpUtil httpUtil = HttpUtil();
  bool showpass = true;
  bool hasAccount = true;
  // ConnectivityResult _connectionStatus = ConnectivityResult.none;
  // final Connectivity _connectivity = Connectivity();
  // late StreamSubscription<ConnectivityResult> _connectivitySubscription;

  @override
  void initState() {
    super.initState();
    Global.lastChooseMsgToUid = "";
    // initConnectivity();
    // _connectivitySubscription =
    //     _connectivity.onConnectivityChanged.listen(_updateConnectionStatus);
  }

  @override
  void dispose() {
    userAccountController.dispose();
    passwordController.dispose();
    // _connectivitySubscription.cancel();
    super.dispose();
  }

  // Future<void> initConnectivity() async {
  //   late ConnectivityResult result;
  //   // Platform messages may fail, so we use a try/catch PlatformException.
  //   try {
  //     result = await _connectivity.checkConnectivity();
  //   } on PlatformException catch (e) {
  //     debugPrint('Couldn\'t check connectivity status' + e.toString());
  //     return;
  //   }
  //   if (!mounted) {
  //     return Future.value(null);
  //   }
  //   return _updateConnectionStatus(result);
  // }

  // Future<void> _updateConnectionStatus(ConnectivityResult result) async {
  //   setState(() {
  //     _connectionStatus = result;
  //   });
  // }

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.black,
      child: Center(
        child: Card(
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
            clipBehavior: Clip.antiAlias,
            shadowColor: Colors.white,
            elevation: 20,
            child: SizedBox(
              width: MediaQuery.of(context).size.width * 0.68,
              height: MediaQuery.of(context).size.width * 0.48,
              child: Row(
                children: [
                  Image(
                    image: const AssetImage(
                      "assets/images/1.jpg",
                    ),
                    width: MediaQuery.of(context).size.width * 0.34,
                    height: MediaQuery.of(context).size.width * 0.48,
                    fit: BoxFit.fitHeight,
                  ),
                  Container(
                    width: MediaQuery.of(context).size.width * 0.34,
                    // height: 200,
                    color: const Color.fromARGB(95, 238, 238, 238),
                    child: Column(
                      children: [
                        const Icon(
                          Icons.image,
                          size: 56,
                        ),
                        hasAccount ? loginWidget() : registerWidget(),
                      ],
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      crossAxisAlignment: CrossAxisAlignment.center,
                    ),
                  ),
                ],
                mainAxisAlignment: MainAxisAlignment.start,
                crossAxisAlignment: CrossAxisAlignment.center,
              ),
            )),
      ),
    );
  }

  dynamic changeShowpass() {
    setState(() {
      showpass = !showpass;
    });
  }

  dynamic registerWidget() {
    return SizedBox(
        width: MediaQuery.of(context).size.width * 0.21,
        child: Column(children: [
          TextField(
            cursorColor: Colors.black,
            controller: userAccountController,
            keyboardType: TextInputType.phone,
            inputFormatters: <TextInputFormatter>[
              FilteringTextInputFormatter.digitsOnly,
              LengthLimitingTextInputFormatter(11),
            ],
            decoration: const InputDecoration(
                icon: Icon(Icons.person),
                helperText: "",
                hintText: "手机",
                hintStyle: TextStyle(color: Colors.black45),
                floatingLabelBehavior: FloatingLabelBehavior.auto),
          ),
          TextField(
            cursorColor: Colors.black,
            controller: passwordController,
            obscureText: showpass,
            inputFormatters: <TextInputFormatter>[
              FilteringTextInputFormatter(RegExp("^[a-z0-9A-Z]+"), allow: true),
              LengthLimitingTextInputFormatter(15),
            ],
            decoration: InputDecoration(
                icon: const Icon(Icons.lock),
                helperText: "",
                hintText: "密码",
                hintStyle: const TextStyle(color: Colors.black45),
                floatingLabelBehavior: FloatingLabelBehavior.auto,
                focusColor: Colors.black,
                suffixIcon: GestureDetector(
                    onTap: changeShowpass,
                    child: const Icon(
                      Icons.remove_red_eye,
                      size: 14,
                    ))),
          ),
          const SizedBox(
            height: 24,
          ),
          GestureDetector(
            onTap: regist,
            child: Container(
              width: MediaQuery.of(context).size.width * 0.19,
              height: 48,
              decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(4.0),
                  color: Colors.black,
                  boxShadow: const [
                    BoxShadow(
                        color: Color(0x99FFFF00),
                        offset: Offset(5.0, 5.0),
                        blurRadius: 10.0,
                        spreadRadius: 2.0),
                    BoxShadow(
                        color: Color(0x9900FF00), offset: Offset(1.0, 1.0)),
                    BoxShadow(color: Color(0xFF0000FF))
                  ]),
              alignment: Alignment.center,
              child: const Text(
                "注册并登录",
                style: TextStyle(
                  color: Colors.white,
                ),
              ),
            ),
          ),
          GestureDetector(
            onTap: () {
              setState(() {
                hasAccount = true;
              });
            },
            child: Container(
              width: MediaQuery.of(context).size.width * 0.19,
              height: 40,
              alignment: Alignment.centerRight,
              child: const Text(
                "已有账号，去登录",
                style: TextStyle(
                  color: Colors.black,
                ),
              ),
            ),
          )
        ]));
  }

  dynamic loginWidget() {
    return SizedBox(
        width: MediaQuery.of(context).size.width * 0.21,
        child: Column(children: [
          TextField(
            cursorColor: Colors.black,
            controller: userAccountController,
            keyboardType: TextInputType.phone,
            inputFormatters: <TextInputFormatter>[
              FilteringTextInputFormatter.digitsOnly,
              LengthLimitingTextInputFormatter(11),
            ],
            decoration: const InputDecoration(
                icon: Icon(Icons.person),
                helperText: "",
                hintText: "手机",
                hintStyle: TextStyle(color: Colors.black45),
                floatingLabelBehavior: FloatingLabelBehavior.auto),
          ),
          TextField(
            cursorColor: Colors.black,
            controller: passwordController,
            obscureText: showpass,
            inputFormatters: <TextInputFormatter>[
              FilteringTextInputFormatter(RegExp("^[a-z0-9A-Z]+"), allow: true),
              LengthLimitingTextInputFormatter(15),
            ],
            decoration: InputDecoration(
                icon: const Icon(Icons.lock),
                helperText: "",
                hintText: "密码",
                hintStyle: const TextStyle(color: Colors.black45),
                floatingLabelBehavior: FloatingLabelBehavior.auto,
                focusColor: Colors.black,
                suffixIcon: GestureDetector(
                    onTap: changeShowpass,
                    child: const Icon(
                      Icons.remove_red_eye,
                      size: 14,
                    ))),
          ),
          const SizedBox(
            height: 24,
          ),
          GestureDetector(
            onTap: login,
            child: Container(
              width: MediaQuery.of(context).size.width * 0.19,
              height: 48,
              decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(4.0),
                  color: Colors.black,
                  boxShadow: const [
                    BoxShadow(
                        color: Color(0x99FFFF00),
                        offset: Offset(5.0, 5.0),
                        blurRadius: 10.0,
                        spreadRadius: 2.0),
                    BoxShadow(
                        color: Color(0x9900FF00), offset: Offset(1.0, 1.0)),
                    BoxShadow(color: Color(0xFF0000FF))
                  ]),
              alignment: Alignment.center,
              child: const Text(
                "登录",
                style: TextStyle(
                  color: Colors.white,
                ),
              ),
            ),
          ),
          GestureDetector(
            onTap: () {
              setState(() {
                hasAccount = false;
              });
            },
            child: Container(
              width: MediaQuery.of(context).size.width * 0.19,
              height: 40,
              alignment: Alignment.centerRight,
              child: const Text(
                "没有账号？",
                style: TextStyle(
                  color: Colors.black,
                ),
              ),
            ),
          )
        ]));
  }

  dynamic regist() {
    if (userAccountController.text == "" || passwordController.text == "") {
      EasyLoading.showToast("账号或密码为空",
          duration: const Duration(milliseconds: 500));
    }
    Map<String, dynamic> _map = {};
    _map["phone"] = userAccountController.text;
    _map["password"] = passwordController.text;

    httpUtil.postRequest(BaseUrl.userBaseUrl, api[ApiType.registUrl], _map,
        success: (res) {
      // debugPrint(res["token"]);
      initdb(_map["phone"]).then((value) => {
            EasyLoading.showToast("注册并登录成功",
                duration: const Duration(milliseconds: 500)),
            getUserInfo(res["token"]),
            Global.token = res["token"]
          });
    }, fail: (res) {
      // debugPrint(res.toString());
      EasyLoading.showToast(res["message"],
          duration: const Duration(seconds: 3));
    });
  }

  dynamic login() {
    if (userAccountController.text == "" || passwordController.text == "") {
      EasyLoading.showToast("账号或密码为空",
          duration: const Duration(milliseconds: 500));
      return;
    }
    Map<String, dynamic> _map = {};
    _map["phone"] = userAccountController.text;
    _map["password"] = passwordController.text;

    httpUtil.postRequest(BaseUrl.userBaseUrl, api[ApiType.loginUrl], _map,
        success: (res) {
      // debugPrint(res["token"]);
      initdb(_map["phone"]).then((value) => {
            EasyLoading.showToast("登录成功",
                duration: const Duration(milliseconds: 500)),
            getUserInfo(res["token"]),
            Global.token = res["token"]
          });
    }, fail: (res) {
      // debugPrint(res.toString());
      EasyLoading.showToast(res.message, duration: const Duration(seconds: 3));
    });
  }

  Future initdb(String id) async {
    if (mounted) {
      await HiveDbUtil.install(id);
      await HiveDbUtil.getInstance();
    }
  }

  dynamic getUserInfo(token) {
    httpUtil.postRequest(
        BaseUrl.userBaseUrl, api[ApiType.getUserInfo], {"token": token},
        success: (data) => {
              debugPrint(data.toString()),
              // saveUserInfo(data["userInfo"])
              // userInfo = UserInfo.fromJson(data["userInfo"]),
              // saveUserFile(userInfo),
              initGlobalAndPushMain(data["userInfo"]),
            },
        fail: (data) => {debugPrint(data.toString())});
  }

  // dynamic saveUserInfo(user) async {
  //   SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
  //   sharedPreferences.setString("userInfo", json.encode(user.toString()));
  // }

  dynamic initGlobalAndPushMain(Map<String, dynamic> map) {
    Global.init(map).then((value) => Navigator.pushNamedAndRemoveUntil(
        context, PageName.mainView.toString(), (route) => false,
        arguments: BundleArguments()));
  }
}
