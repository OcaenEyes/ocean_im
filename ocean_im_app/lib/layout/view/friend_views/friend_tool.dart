import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:hive_flutter/hive_flutter.dart';

import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/model/friend_model.dart';
import 'package:ocean_im/model/friend_receive_req_model.dart';
import 'package:ocean_im/model/tmp_friend_model.dart';
import 'package:ocean_im/model/user_info_model.dart';
import 'package:ocean_im/util/api/base_api.dart';
import 'package:ocean_im/util/db/hive_db_util.dart';
import 'package:ocean_im/util/net/http_util.dart';

class FriendTool extends StatefulWidget {
  const FriendTool({Key? key}) : super(key: key);
  @override
  _FriendToolState createState() => _FriendToolState();
}

class _FriendToolState extends State<FriendTool> {
  TextEditingController searchUserPhoneController = TextEditingController();
  TextEditingController reqCommentController = TextEditingController();

  UserInfo searchFriendInfo = UserInfo("", "", "", "", "", "", 0, "");
  @override
  Widget build(BuildContext context) {
    return Container(
        width: MediaQuery.of(context).size.width - 356,
        height: MediaQuery.of(context).size.height,
        padding: const EdgeInsets.all(10),
        // color: const Color.fromARGB(255, 223, 223, 223),
        child: Column(
          children: [
            Container(
              decoration:
                  BoxDecoration(borderRadius: BorderRadius.circular(8.0)),
              width: (MediaQuery.of(context).size.width - 356) / 1.4,
              height: 48,
              child: TextField(
                cursorColor: Colors.black26,
                controller: searchUserPhoneController,
                decoration: InputDecoration(
                  hintText: "搜索添加好友，可以输入对方的手机号进行检索",
                  suffixIcon: GestureDetector(
                      onTap: getUserByPhone,
                      child: const Icon(
                        Icons.search_rounded,
                        color: Colors.black26,
                      )),
                  enabledBorder: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(5.0),
                    borderSide: const BorderSide(
                      width: 0.1,
                      color: Colors.black26,
                    ),
                  ),
                  focusedBorder: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(5.0),
                    borderSide: const BorderSide(
                      width: 0.1,
                      color: Colors.black26,
                    ),
                  ),
                ),
              ),
            ),
            searchFriendInfo.uid != ""
                ? Container(
                    margin: const EdgeInsets.fromLTRB(0, 10, 0, 10),
                    width: (MediaQuery.of(context).size.width - 356) / 1.4,
                    height: 130,
                    decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(8),
                        color: const Color.fromARGB(135, 240, 240, 240)),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceAround,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        Container(
                          width: 66,
                          height: 66,
                          decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(8),
                              image: const DecorationImage(
                                  image: AssetImage("assets/images/2.jpg"),
                                  fit: BoxFit.cover)),
                        ),
                        Stack(
                          children: [
                            SizedBox(
                              width: (MediaQuery.of(context).size.width - 356) /
                                      1.4 -
                                  120,
                              child: Column(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceAround,
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(
                                    "昵称:" + (searchFriendInfo.nickname),
                                    style: const TextStyle(
                                        letterSpacing: 1.2,
                                        fontWeight: FontWeight.w500),
                                  ),
                                  Text(
                                    "账号:" + (searchFriendInfo.uid),
                                    style: const TextStyle(
                                        letterSpacing: 1.2,
                                        fontWeight: FontWeight.w500),
                                  ),
                                  Text(
                                    "邮箱:" + (searchFriendInfo.email),
                                    style: const TextStyle(
                                        letterSpacing: 1.2,
                                        fontWeight: FontWeight.w500),
                                  ),
                                  Text(
                                    "性别:" +
                                        (searchFriendInfo.gender == 0
                                            ? "未知"
                                            : (searchFriendInfo.gender == 1
                                                ? "男"
                                                : "女")),
                                    style: const TextStyle(
                                        letterSpacing: 1.2,
                                        fontWeight: FontWeight.w500),
                                  ),
                                  Text(
                                    "签名:" + (searchFriendInfo.ex),
                                    style: const TextStyle(
                                        letterSpacing: 1,
                                        fontWeight: FontWeight.w500),
                                  ),
                                ],
                              ),
                            ),
                            Positioned(
                                right: 10,
                                top: 40,
                                child: HiveDbUtil.friendsClient
                                        .containsKey(searchFriendInfo.uid)
                                    ? IconButton(
                                        onPressed: sendToFriend,
                                        icon: const Icon(
                                          Icons.send,
                                          color: Colors.green,
                                        ),
                                      )
                                    : IconButton(
                                        onPressed: () {
                                          showDialog(
                                              context: context,
                                              builder: (context) => AlertDialog(
                                                    title: const Text(
                                                      "添加请求",
                                                      style: TextStyle(
                                                          fontSize: 18),
                                                    ),
                                                    content: TextField(
                                                      cursorColor: Colors.black,
                                                      controller:
                                                          reqCommentController,
                                                    ),
                                                    actions: [
                                                      TextButton(
                                                          onPressed: () {
                                                            sendFriendReq();
                                                            Navigator.pop(
                                                                context);
                                                          },
                                                          child: const Text(
                                                            "发送",
                                                            style: TextStyle(
                                                                color: Colors
                                                                    .black),
                                                          )),
                                                      TextButton(
                                                          onPressed: () {
                                                            Navigator.pop(
                                                                context);
                                                          },
                                                          child: const Text(
                                                            "取消",
                                                            style: TextStyle(
                                                                color: Colors
                                                                    .black),
                                                          ))
                                                    ],
                                                  ));
                                        },
                                        icon: const Icon(
                                          Icons.add,
                                          color: Colors.amber,
                                        ),
                                      ))
                          ],
                        )
                      ],
                    ),
                  )
                : Container(),
            Container(
              width: (MediaQuery.of(context).size.width - 356) / 1.4,
              alignment: Alignment.centerLeft,
              // color: Colors.amber,
              height: 48,
              child: const Text(
                "下面是收到的好友请求",
                style: TextStyle(fontWeight: FontWeight.w400),
              ),
            ),
            ValueListenableBuilder(
              valueListenable:
                  Hive.box<Req_receive>("friendreqbox").listenable(),
              builder: (context, Box<Req_receive> friendreq, _) {
                return Card(
                    shadowColor: Colors.black38,
                    elevation: 20,
                    color: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                    child: friendreq.isNotEmpty
                        ? Container(
                            width:
                                (MediaQuery.of(context).size.width - 356) / 1.4,
                            margin: const EdgeInsets.fromLTRB(0, 0, 0, 4),
                            height: searchFriendInfo.uid != ""
                                ? MediaQuery.of(context).size.height - 278
                                : MediaQuery.of(context).size.height - 128,
                            // decoration: BoxDecoration(
                            //     border: Border.all(
                            //         color: Colors.black26, width: 0.1),
                            //     borderRadius: BorderRadius.circular(6)),
                            child: ListView.builder(
                                itemCount: friendreq.length,
                                itemBuilder:
                                    (BuildContext context, int position) {
                                  return getRow(
                                      friendreq.values.toList()[position]);
                                }),
                          )
                        : Container(
                            width: 300,
                            alignment: Alignment.topCenter,
                            child: const Text("em..也暂无请求呢"),
                          ));
              },
            ),
          ],
        ));
  }

  dynamic getRow(Req_receive reqfrined) {
    return ListTile(
        onTap: () {
          reqfrined.flag == 0
              ? showDialog(
                  context: context,
                  builder: (context) => AlertDialog(
                        title: const Text(
                          "来处理一下好友请求",
                          style: TextStyle(fontSize: 18),
                        ),
                        content: Text(reqfrined.reqMessage != ""
                            ? reqfrined.reqMessage
                            : "可以加我为好友吗? ^-^"),
                        actions: [
                          TextButton(
                              onPressed: () {
                                debugPrint("点击了忽略");
                                dueFriendReq(4, reqfrined);
                                Navigator.pop(context);
                              },
                              child: const Text('忽略',
                                  style: TextStyle(color: Colors.black))),
                          TextButton(
                              onPressed: () {
                                debugPrint("点击了拒绝");
                                dueFriendReq(2, reqfrined);
                                Navigator.pop(context);
                              },
                              child: const Text('拒绝',
                                  style: TextStyle(color: Colors.black))),
                          TextButton(
                              onPressed: () {
                                debugPrint("点击了接受");
                                dueFriendReq(1, reqfrined);
                                Navigator.pop(context);
                              },
                              child: const Text('接受',
                                  style: TextStyle(color: Colors.black))),
                        ],
                      ))
              : "";
        },
        hoverColor: Colors.green,
        leading: Stack(
          children: [
            Container(
              width: 34,
              height: 34,
              decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(6),
                  image: const DecorationImage(
                      image: AssetImage("assets/images/3.jpg"),
                      fit: BoxFit.fill)),
            ),
            Positioned(
                top: 0,
                right: 0,
                child: Container(
                  width: 16,
                  height: 16,
                  alignment: Alignment.center,
                  child: Text(
                    (reqfrined.flag == 0) ? "1" : "",
                    style: const TextStyle(
                      color: Colors.white,
                      fontSize: 10,
                    ),
                  ),
                  decoration: (reqfrined.flag == 0)
                      ? BoxDecoration(
                          color: Colors.red[400],
                          borderRadius: BorderRadius.circular(6))
                      : null,
                ))
          ],
        ),
        title: Text(
          reqfrined.reqId,
          style: const TextStyle(fontSize: 14),
        ),
        trailing: Text(reqfrined.flag == 0
            ? "待处理"
            : (reqfrined.flag == 1
                ? "已接受"
                : (reqfrined.flag == 2
                    ? "已拒绝"
                    : (reqfrined.flag == 4 ? "已忽略" : "")))));
  }

  dynamic getUserByPhone() {
    if (searchUserPhoneController.text == "") {
      EasyLoading.showToast("请输入要搜索的用户信息",
          duration: const Duration(milliseconds: 500));
      return;
    }

    HttpUtil httpUtil = HttpUtil();
    httpUtil.postRequest(BaseUrl.userBaseUrl, api[ApiType.getUserInfoByPhone],
        {"phone": searchUserPhoneController.text},
        success: (data) =>
            {debugPrint(data.toString()), dugSearchRes(data["userInfo"])},
        fail: (data) => {
              EasyLoading.showToast(data["message"],
                  duration: const Duration(milliseconds: 500))
            });
  }

  dynamic dugSearchRes(Map<String, dynamic> tmp) {
    setState(() {
      searchFriendInfo = UserInfo.fromJson(tmp);
    });
  }

  dynamic sendToFriend() {
    Global.lastChooseMsgToUid = searchFriendInfo.uid;
    Global.pageController.jumpToPage(0);
    if (Global.talkList.value.contains(searchFriendInfo.uid)) {
      Global.talkList.value.remove(searchFriendInfo.uid);
      Global.talkList.value.insert(0, searchFriendInfo.uid);
    } else {
      Global.talkList.value.insert(0, searchFriendInfo.uid);
    }
    setState(() {
      searchUserPhoneController.clear();
      searchFriendInfo = UserInfo("", "", "", "", "", "", 0, "");
    });
  }

  dynamic sendFriendReq() {
    HttpUtil httpUtil = HttpUtil();
    httpUtil.postRequest(
        BaseUrl.friendBaseUrl,
        api[ApiType.sendFriendReqUrl],
        {
          "req_id": Global.userInfo.uid,
          "user_id": searchFriendInfo.uid,
          "req_message": reqCommentController.text
        },
        success: (data) => {debugPrint(data.toString())},
        fail: (data) => {
              EasyLoading.showToast(data["message"],
                  duration: const Duration(milliseconds: 500))
            });

    setState(() {
      searchUserPhoneController.clear();
      searchFriendInfo = UserInfo("", "", "", "", "", "", 0, "");
    });
  }

  dynamic dueFriendReq(int i, Req_receive reqf) {
    HttpUtil httpUtil = HttpUtil();

    httpUtil.postRequest(BaseUrl.friendBaseUrl, api[ApiType.dueFriendReqUrl], {
      "req_id": reqf.reqId,
      "user_id": reqf.userId,
      "flag": i,
    }, success: ((data) {
      debugPrint(data.toString());
      if (mounted) {
        Global.countFriendReqs.value -= 1;
        Req_receive tmp = HiveDbUtil.friendReqReceiveClient.get(reqf.reqId);
        tmp.flag = i;
        HiveDbUtil.friendReqReceiveClient.put(reqf.reqId, tmp);
      }
      if (i == 1) {
        TmpFriendModel tmpf = HiveDbUtil.tmpFriendsClient.get(reqf.reqId);
        Friends f = Friends(tmpf.uid, tmpf.remark, tmpf.nickname, tmpf.phone,
            tmpf.email, tmpf.birth, tmpf.headImg, tmpf.gender, tmpf.ex);
        HiveDbUtil.friendsClient.put(reqf.reqId, f);
      }
    }), fail: ((data) {}));
  }
}
