import 'package:flutter/material.dart';
import 'package:ocean_im/model/user_info_model.dart';
import 'package:ocean_im/util/directory_util/file_util.dart';
import 'package:ocean_im/util/ws/init_ws.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class Global {
  static Global global = Global._instance();
  Global._instance();

  static late PageController pageController;
  static String lastChooseMsgToUid = "";
  static late UserInfo userInfo;
  static String token = "";
  static late FileUtil fileUtil;
  static late String userPath;
  static late IO.Socket socket;
  static late ValueNotifier<int> countNoRead;
  static late ValueNotifier<List<String>> talkList;
  static late ValueNotifier<int> countFriendReqs;
  static List<MaterialColor> get themes => <MaterialColor>[
        Colors.blue,
        Colors.cyan,
        Colors.teal,
        Colors.green,
        Colors.red,
      ];

  static Future init(Map<String, dynamic> map) async {
    userInfo = UserInfo.fromJson(map);
    fileUtil = FileUtil.fileUtil;
    countNoRead = ValueNotifier<int>(0);
    countFriendReqs = ValueNotifier<int>(0);
    talkList = ValueNotifier<List<String>>([]);
    // userPath = '${await fileUtil.localDocumentPath}' "/oceanim/" +
    //     userInfo.uid +
    //     "/data";
    // fileUtil.createFile(userPath + "/info.res");
    // fileUtil.writeFile(userPath + "/info.res", userInfo.toJson().toString());
    socket = initWs();
    Global.socket.connect();
    Global.socket
        .emit("register", {"uid": Global.userInfo.uid, "platform": "windows"});
    Global.socket.on("error", (data) => debugPrint("error:" + data.toString()));
    Global.socket
        .onReconnectError((data) => debugPrint("重连失败:" + data.toString()));
  }
}
