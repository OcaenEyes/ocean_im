import 'dart:async';

import 'package:flutter/material.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/layout/view/friend_views/friend_view.dart';
import 'package:ocean_im/layout/view/msg_views/msg_view.dart';
import 'package:ocean_im/model/chat_message_model.dart';
import 'package:ocean_im/model/friend_model.dart';
import 'package:ocean_im/model/friend_receive_req_model.dart';
import 'package:ocean_im/model/tmp_friend_model.dart';
import 'package:ocean_im/model/user_info_model.dart';
import 'package:ocean_im/router/page_builder.dart';
import 'package:ocean_im/router/page_routes.dart';
import 'package:ocean_im/util/api/base_api.dart';
import 'package:ocean_im/util/db/hive_db_util.dart';
import 'package:ocean_im/util/net/http_util.dart';
// import 'package:shared_preferences/shared_preferences.dart';
import 'package:socket_io_client/socket_io_client.dart';

class MainView extends StatefulWidget {
  final BundleArguments bundleArguments;
  const MainView({Key? key, required this.bundleArguments}) : super(key: key);

  @override
  _MainViewState createState() => _MainViewState();
}

class _MainViewState extends State<MainView> {
  late UserInfoModel userInfoModel;
  late int currentIndex;
  List<Widget> viewList = [];
  bool shouldCancel = false;
  late String userId;
  @override
  void initState() {
    viewList
      ..add(const MsgView())
      ..add(const FriendView());
    Global.pageController = PageController(initialPage: 0, viewportFraction: 1);
    if (!Global.socket.connected) {
      Global.socket.connect();
      Global.socket.emit(
          "register", {"uid": Global.userInfo.uid, "platform": "windows"});
    }

    Global.socket.on("singChat", (data) => {dueSocktMsg(data)});
    Global.socket.on("heartRes", (data) => {dueHeartRes(data)});
    Global.socket.onDisconnect((data) => {dueDisconn(data)});
    readFriendsample();
    readChatSample();
    getAllFriendReq();
    heatlisten();
    super.initState();
    Global.socket.on("receive_friend_req", (data) => {dueReceiveFriReq(data)});
    Global.socket.on("receive_due_friend", (data) => {dueFriReqRes(data)});
    if (mounted) {
      currentIndex = 0;
    }
  }

  @override
  void dispose() {
    super.dispose();
    Global.pageController.dispose();
    Global.socket.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [mainAside(), mainContent()],
    );
  }

  dynamic heatlisten() {
    Timer.periodic(const Duration(seconds: 5), (timer) {
      Global.socket.emit("heartBeat", {
        "uid": Global.userInfo.uid,
        "last_act_time": DateTime.now().toString()
      });
      // debugPrint("定时器再运行了" + DateTime.now().toString());
      if (shouldCancel) {
        timer.cancel();
      }
    });
  }

  dynamic dueHeartRes(Map<String, dynamic> tmp) {
    if (tmp["status"] == "heart success") {
      // debugPrint("还在登录中");
      // debugPrint(tmp.toString());
    } else {
      // debugPrint(tmp.toString());
      // debugPrint("没有登录了");
      // debugPrint(tmp.toString());
      if (mounted) {
        setState(() {
          shouldCancel = true;
        });
      }
    }
  }

  dynamic dueDisconn(dynamic tmp) {
    debugPrint("断开链接了:" + tmp.toString());
    Global.socket.connect();
    Global.socket
        .emit("register", {"uid": Global.userInfo.uid, "platform": "windows"});
    // if (mounted) {
    //   if (!Global.socket.connected) {
    //     setState(() {
    //       shouldCancel = true;
    //     });
    //   }
    // }
  }

  dynamic readChatSample() {
    HttpUtil httpUtil = HttpUtil();
    late List<Chats> tmp;
    late List list;
    httpUtil.postRequest(
        BaseUrl.msgBaseUrl, api[ApiType.getAllMsg], {"token": Global.token},
        success: (data) => {
              list = data["chats"] ?? [],
              // debugPrint(list.toString()),
              tmp = list.map((e) => Chats.fromJson(e)).toList(),
              dueMsg(tmp)
            },
        fail: (data) => {});
  }

  dynamic dueMsg(List<Chats> tmp) {
    for (var element in tmp) {
      // if (HiveDbUtil.chatsClient.containsKey(element.toUser)) {
      //   Chats tmp = HiveDbUtil.chatsClient.get(element.toUser);
      //   tmp.detailList.addAll(element.detailList);
      //   HiveDbUtil.chatsClient.put(element.toUser, tmp);
      // } else {
      //   HiveDbUtil.chatsClient.put(element.toUser, element);
      // }
      for (var item in element.detailList) {
        if (item.signType == 0 &&
            item.sendId != Global.userInfo.uid &&
            item.sendId != Global.lastChooseMsgToUid) {
          Global.countNoRead.value += 1;
          debugPrint("登录以后当前的未读数量");
          debugPrint(Global.countNoRead.value.toString());
        }
      }
      Global.talkList.value.add(element.toUser);
      HiveDbUtil.chatsClient.put(element.toUser, element);
    }
    for (Chats item in HiveDbUtil.chatsClient.values) {
      if (item.detailList.isEmpty) {
        HiveDbUtil.chatsClient.delete(item.toUser);
      }
    }
  }

  dynamic readFriendsample() {
    late List list;
    late List<Friends> tmp;
    HttpUtil httpUtil = HttpUtil();
    httpUtil.postRequest(BaseUrl.friendBaseUrl, api[ApiType.getAllFriends],
        {"token": Global.token},
        success: (data) => {
              list = data["friends"] ?? [],
              // debugPrint(list.toString()),
              tmp = list.map((e) => Friends.fromJson(e)).toList(),
              dueFriend(tmp),
            },
        fail: (data) => {});

    // rootBundle.loadString("lib/test_data/friend.json").then((value) {
    //   var a = json.decode(value);
    //   List b = a["friends"];
    //   tmp = b.map((e) => Friends.fromJson(e)).toList();
    //   setState(() {
    //     friends = tmp;
    //   });
    // });
  }

  dynamic dueFriend(List<Friends> tmp) {
    // List<String> tmppp = [];
    // tmppp = tmp.map((e) => e.uid).toList();
    for (var item in tmp) {
      HiveDbUtil.friendsClient.put(item.uid, item);
    }
    // for (var item in HiveDbUtil.friendsClient.keys) {
    //   if (!tmppp.contains(item)) {
    //     HiveDbUtil.friendsClient.delete(item);
    //   }
    // }
  }

  dynamic getAllFriendReq() {
    HttpUtil httpUtil = HttpUtil();
    late List list;
    late Req_receive friendreq;
    httpUtil.postRequest(
        BaseUrl.friendBaseUrl,
        api[ApiType.getAllFriendReqs],
        {
          "token": Global.token,
        },
        success: (data) => {
              debugPrint(data.toString()),
              if (data["req_receive"] != null)
                {
                  list = data["req_receive"] ?? [],
                  for (var item in list)
                    {
                      friendreq = Req_receive.fromJson(item),
                      if (friendreq.flag == 0)
                        {Global.countFriendReqs.value += 1},
                      HiveDbUtil.friendReqReceiveClient
                          .put(friendreq.reqId, friendreq)
                    }
                }
            },
        fail: (data) => {});
  }

  dynamic mainAside() {
    return Container(
      width: 56,
      color: const Color.fromARGB(255, 27, 27, 27),
      child:
          Column(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
        Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              const SizedBox(height: 40),
              SizedBox(
                child: GestureDetector(
                  onTap: head,
                  child: Container(
                    width: 36,
                    height: 36,
                    decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(6),
                        color: Colors.white,
                        image: const DecorationImage(
                            image: AssetImage("assets/images/2.jpg"),
                            fit: BoxFit.cover)),
                  ),
                ),
              ),
              const SizedBox(height: 20),
              SizedBox(
                height: 56,
                width: 56,
                child: Stack(
                  alignment: Alignment.center,
                  children: [
                    GestureDetector(
                        onTap: msg,
                        child: const Icon(
                          Icons.mail,
                          size: 20,
                          color: Colors.white,
                        )),
                    mounted
                        ? ValueListenableBuilder(
                            valueListenable: Global.countNoRead,
                            builder: ((context, int value, _) {
                              return Positioned(
                                  top: 8,
                                  right: 6,
                                  child: Container(
                                    width: 16,
                                    height: 16,
                                    alignment: Alignment.center,
                                    child: Text(
                                      (value > 0) ? value.toString() : "",
                                      style: const TextStyle(
                                        color: Colors.white,
                                        fontSize: 10,
                                      ),
                                    ),
                                    decoration: (value > 0)
                                        ? BoxDecoration(
                                            color: Colors.red[400],
                                            borderRadius:
                                                BorderRadius.circular(6))
                                        : null,
                                  ));
                            }),
                          )
                        : Container()
                  ],
                ),
              ),
              SizedBox(
                  height: 56,
                  width: 56,
                  child: Stack(alignment: Alignment.center, children: [
                    GestureDetector(
                        onTap: friend,
                        child: const Icon(
                          Icons.group,
                          size: 20,
                          color: Colors.white,
                        )),
                    mounted
                        ? ValueListenableBuilder(
                            valueListenable: Global.countFriendReqs,
                            builder: ((context, int value, _) {
                              return Positioned(
                                  top: 8,
                                  right: 6,
                                  child: Container(
                                    width: 16,
                                    height: 16,
                                    alignment: Alignment.center,
                                    child: Text(
                                      (value > 0) ? value.toString() : "",
                                      style: const TextStyle(
                                        color: Colors.white,
                                        fontSize: 10,
                                      ),
                                    ),
                                    decoration: (value > 0)
                                        ? BoxDecoration(
                                            color: Colors.red[400],
                                            borderRadius:
                                                BorderRadius.circular(6))
                                        : null,
                                  ));
                            }),
                          )
                        : Container(),
                  ])),
              // SizedBox(
              //   height: 56,
              //   child: GestureDetector(
              //     onTap: circle,
              //     child: const Icon(
              //       Icons.camera_enhance,
              //       size: 20,
              //       color: Colors.white,
              //     ),
              //   ),
              // ),
              SizedBox(
                height: MediaQuery.of(context).size.height - 268,
              ),
              SizedBox(
                height: 40,
                child: GestureDetector(
                    onTap: logout,
                    child: const Icon(
                      Icons.logout,
                      color: Colors.white,
                      size: 20,
                    )),
              ),
              const SizedBox(
                height: 20,
              )
            ]),
      ]),
    );
  }

  dynamic mainContent() {
    return SizedBox(
      width: MediaQuery.of(context).size.width - 56,
      height: MediaQuery.of(context).size.height,
      child: PageView.builder(
        // pageSnapping: false,
        scrollDirection: Axis.vertical,
        controller: Global.pageController,
        physics: const NeverScrollableScrollPhysics(),
        itemCount: viewList.length,
        itemBuilder: (BuildContext _, int index) {
          return viewList[index];
        },
      ),
    );
  }

  dynamic head() {
    debugPrint("头像被点击");
    showMenu(
        context: context,
        position: RelativeRect.fromLTRB(
            60, 10, MediaQuery.of(context).size.width - 80, 10),
        shape:
            RoundedRectangleBorder(borderRadius: BorderRadius.circular(10.0)),
        color: Colors.amber[50],
        items: userInfoMenu());
  }

  dynamic userInfoMenu() {
    List<UserInfo> tmpL = [Global.userInfo];
    return tmpL
        .map((e) => PopupMenuItem(
              child: SizedBox(
                width: 140,
                height: 200,
                child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceAround,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const Text("用户信息"),
                      Text("昵称:" + e.nickname),
                      Text("账号:" + e.uid),
                      Text("手机:" + e.phone),
                      Text("邮箱:" + (e.email != "" ? e.email : "暂无")),
                      Text("性别:" +
                          (e.gender == 1
                              ? "男"
                              : e.gender == 2
                                  ? "女"
                                  : "未知 ")),
                    ]),
              ),
            ))
        .toList();
  }

  dynamic msg() {
    debugPrint("消息被点击");
    if (Global.pageController.hasClients) {
      setState(() {
        currentIndex = 0;
        Global.pageController.jumpToPage(currentIndex);
        // pageController.animateToPage(
        //     currentIndex * MediaQuery.of(context).size.height.toInt(),
        //     duration: const Duration(milliseconds: 400),
        //     curve: Curves.easeInOut);
      });
    }
  }

  dynamic friend() {
    debugPrint("朋友被点击");
    if (Global.pageController.hasClients) {
      setState(() {
        currentIndex = 1;
        // pageController.animateToPage(
        //     currentIndex * MediaQuery.of(context).size.height.toInt(),
        //     duration: const Duration(milliseconds: 400),
        //     curve: Curves.easeInOut);
        Global.pageController.jumpToPage(currentIndex);
      });
    }
  }

  dynamic circle() {
    debugPrint("朋友圈被点击");
    if (Global.pageController.hasClients) {
      setState(() {
        currentIndex = 2;
        Global.pageController.jumpToPage(currentIndex);
      });
    }
  }

  dynamic logout() async {
    debugPrint("退出被点击");
    Global.socket.emit("logout", Global.userInfo.uid);
    Global.socket.close();
    setState(() {
      shouldCancel = true;
    });
    HiveDbUtil.closeInstance().then((value) => Navigator.popAndPushNamed(
        context, PageName.login.toString(),
        arguments: BundleArguments()));
  }

  // dynamic removeLocalData() async {
  //   SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
  //   sharedPreferences.clear();
  // }

  dynamic dueSocktMsg(Map<String, dynamic> tmp) {
    debugPrint("收到了单聊");
    debugPrint(tmp.toString());

    Detail_list detail = Detail_list.fromJson(tmp);
    // debugPrint("detail.toString()");
    // debugPrint(detail.toString());
    if (HiveDbUtil.chatsClient.containsKey(detail.sendId)) {
      Chats chats = HiveDbUtil.chatsClient.get(detail.sendId);
      chats.detailList.insert(0, detail);
      HiveDbUtil.chatsClient.put(detail.sendId, chats);
    } else {
      List<Detail_list> tmpList = [];
      Chats chats = Chats(tmpList, detail.sendId);
      tmpList.add(detail);
      HiveDbUtil.chatsClient.put(detail.sendId, chats);
    }
    if (Global.talkList.value.contains(detail.sendId)) {
      Global.talkList.value.remove(detail.sendId);
      Global.talkList.value.insert(0, detail.sendId);
    } else {
      Global.talkList.value.insert(0, detail.sendId);
    }

    if (detail.sendId != Global.lastChooseMsgToUid) {
      Global.countNoRead.value += 1;
      debugPrint("收到消息以后当前的未读数量");
      debugPrint(Global.countNoRead.value.toString());
    }
  }

  dynamic dueReceiveFriReq(Map<String, dynamic> tmp) {
    debugPrint("收到好友请求:" + tmp.toString());
    Req_receive reqf = Req_receive.fromJson(tmp["friendR"]);
    TmpFriendModel tmpf = TmpFriendModel.fromJson(tmp["tmpInfo"]);
    HiveDbUtil.tmpFriendsClient.put(tmpf.uid, tmpf);
    if (HiveDbUtil.friendReqReceiveClient.containsKey(reqf.reqId)) {
      Req_receive r = HiveDbUtil.friendReqReceiveClient.get(reqf.reqId);
      if (r.flag != 0) {
        if (mounted) {
          Global.countFriendReqs.value += 1;
        }
        HiveDbUtil.friendReqReceiveClient.put(reqf.reqId, reqf);
      }
    } else {
      if (mounted) {
        Global.countFriendReqs.value += 1;
      }
      HiveDbUtil.friendReqReceiveClient.put(reqf.reqId, reqf);
    }
  }

  dynamic dueFriReqRes(Map<String, dynamic> tmp) {
    debugPrint("对方添加的结果:" + tmp.toString());
    Detail_list detail = Detail_list.fromJson(tmp["msg"]);
    Friends f = Friends.fromJson(tmp["friend"]);
    HiveDbUtil.friendsClient.put(detail.sendId, f);
    if (HiveDbUtil.chatsClient.containsKey(detail.sendId)) {
      Chats chats = HiveDbUtil.chatsClient.get(detail.sendId);
      chats.detailList.insert(0, detail);
      HiveDbUtil.chatsClient.put(detail.sendId, chats);
    } else {
      List<Detail_list> tmpList = [];
      Chats chats = Chats(tmpList, detail.sendId);
      tmpList.add(detail);
      HiveDbUtil.chatsClient.put(detail.sendId, chats);
    }
    if (Global.talkList.value.contains(detail.sendId)) {
      Global.talkList.value.remove(detail.sendId);
      Global.talkList.value.insert(0, detail.sendId);
    } else {
      Global.talkList.value.insert(0, detail.sendId);
    }

    if (detail.sendId != Global.lastChooseMsgToUid) {
      Global.countNoRead.value += 1;
      debugPrint("收到消息以后当前的未读数量");
      debugPrint(Global.countNoRead.value.toString());
    }
    if (HiveDbUtil.friendReqReceiveClient.containsKey(f.uid) &&
        HiveDbUtil.friendReqReceiveClient.get(f.uid).flag == 0) {
      Req_receive tmpr = HiveDbUtil.friendReqReceiveClient.get(f.uid);
      tmpr.flag = 1;
      HiveDbUtil.friendReqReceiveClient.put(f.uid, tmpr);
      Global.countFriendReqs.value -= 1;
    }
  }
}
