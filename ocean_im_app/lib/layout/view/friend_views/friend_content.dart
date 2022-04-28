import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/model/friend_model.dart';
import 'package:ocean_im/util/api/base_api.dart';
import 'package:ocean_im/util/db/hive_db_util.dart';
import 'package:ocean_im/util/net/http_util.dart';

class FriendContent extends StatefulWidget {
  final Friends? friend;
  const FriendContent({Key? key, required this.friend}) : super(key: key);
  @override
  _FriendContentState createState() => _FriendContentState();
}

class _FriendContentState extends State<FriendContent> {
  TextEditingController remarkTextEditingController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return widget.friend != null
        ? Column(
            children: [
              Container(
                height: 56,
                alignment: Alignment.center,
                width: MediaQuery.of(context).size.width - 356,
                color: const Color.fromARGB(135, 240, 240, 240),
                child: const Text(
                  "好友信息",
                  style: TextStyle(fontSize: 18, fontWeight: FontWeight.w500),
                ),
              ),
              Container(
                width: MediaQuery.of(context).size.width - 356,
                height: MediaQuery.of(context).size.height - 66,
                color: const Color.fromARGB(135, 240, 240, 240),
                child: Center(
                    child: Stack(
                  children: [
                    Card(
                        shadowColor: Colors.black38,
                        elevation: 20,
                        color: Colors.white,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(10),
                        ),
                        child: Container(
                          padding: const EdgeInsets.all(10),
                          width: 320,
                          height: 360,
                          child: Column(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                              crossAxisAlignment: CrossAxisAlignment.center,
                              children: [
                                Container(
                                  alignment: Alignment.centerLeft,
                                  width: 46,
                                  height: 46,
                                  margin: const EdgeInsets.only(right: 10),
                                  decoration: BoxDecoration(
                                      borderRadius: BorderRadius.circular(8),
                                      image: const DecorationImage(
                                          image:
                                              AssetImage("assets/images/3.jpg"),
                                          fit: BoxFit.fill)),
                                ),
                                Text("用户ID:" + widget.friend!.uid),
                                Text("昵称:" + widget.friend!.nickname),
                                GestureDetector(
                                  onTap: () {
                                    debugPrint("修改备注");
                                    showDialog(
                                        context: context,
                                        builder: (context) => AlertDialog(
                                              title: const Text("修改备注",
                                                  style:
                                                      TextStyle(fontSize: 18)),
                                              content: TextField(
                                                controller:
                                                    remarkTextEditingController,
                                                cursorColor: Colors.black,
                                              ),
                                              actions: [
                                                TextButton(
                                                    onPressed: () {
                                                      changeRemark();
                                                      Navigator.pop(context);
                                                    },
                                                    child: const Text("修改",
                                                        style: TextStyle(
                                                            color:
                                                                Colors.black))),
                                                TextButton(
                                                    onPressed: () {
                                                      Navigator.pop(context);
                                                    },
                                                    child: const Text("取消",
                                                        style: TextStyle(
                                                            color:
                                                                Colors.black))),
                                              ],
                                            ));
                                  },
                                  child: widget.friend!.remark != ""
                                      ? Text("备注:" + widget.friend!.remark)
                                      : Text(
                                          "备注:点击此处修改备注",
                                          style: TextStyle(
                                              color: Colors.green[200]),
                                        ),
                                ),
                                Text("手机:" + widget.friend!.phone),
                                Text("性别:" +
                                    (widget.friend!.gender == 0
                                        ? "未知"
                                        : (widget.friend!.gender == 1
                                            ? "男"
                                            : "女"))),
                                Text("签名:" + widget.friend!.ex),
                                GestureDetector(
                                  onTap: sendToFriend,
                                  child: Container(
                                    width: 200,
                                    height: 48,
                                    decoration: BoxDecoration(
                                        borderRadius:
                                            BorderRadius.circular(4.0),
                                        color: Colors.black,
                                        boxShadow: const [
                                          BoxShadow(
                                              color: Color(0x99FFFF00),
                                              offset: Offset(5.0, 5.0),
                                              blurRadius: 10.0,
                                              spreadRadius: 2.0),
                                          BoxShadow(
                                              color: Color(0x9900FF00),
                                              offset: Offset(1.0, 1.0)),
                                          BoxShadow(color: Color(0xFF0000FF))
                                        ]),
                                    alignment: Alignment.center,
                                    child: const Text(
                                      "发送信息",
                                      style: TextStyle(
                                        color: Colors.white,
                                      ),
                                    ),
                                  ),
                                )
                              ]),
                        )),
                    Positioned(
                        top: 10,
                        right: 10,
                        child: GestureDetector(
                          onTap: () {
                            showDialog(
                                context: context,
                                builder: (context) => AlertDialog(
                                      title: const Text("确认要删除好友吗？",
                                          style: TextStyle(fontSize: 18)),
                                      actions: [
                                        TextButton(
                                            onPressed: () {
                                              deleteFriend();
                                              Navigator.pop(context);
                                            },
                                            child: const Text("确认",
                                                style: TextStyle(
                                                    color: Colors.black))),
                                        TextButton(
                                            onPressed: () {
                                              Navigator.pop(context);
                                            },
                                            child: const Text("取消",
                                                style: TextStyle(
                                                    color: Colors.black))),
                                      ],
                                    ));
                          },
                          child: const Icon(
                            Icons.delete,
                            color: Colors.black45,
                          ),
                        )),
                  ],
                )),
              ),
            ],
          )
        : Container();
  }

  dynamic sendToFriend() {
    Global.lastChooseMsgToUid = widget.friend!.uid;
    Global.pageController.jumpToPage(0);
    // if (!context.read<MessageProvider>().all.keys.contains(widget.uid)) {
    //   Global.allMsgs[widget.uid] = [];
    //   context.read<MessageProvider>().setAllMsg(Global.allMsgs);
    // }

    // if (!HiveDbUtil.chatsClient.containsKey(widget.friend!.uid)) {
    //   List<Detail_list> tmp = [];
    //   Chats temp = Chats(tmp, widget.friend!.uid);
    //   HiveDbUtil.chatsClient.put(widget.friend!.uid, temp);
    // }
    if (Global.talkList.value.contains(widget.friend!.uid)) {
      Global.talkList.value.remove(widget.friend!.uid);
      Global.talkList.value.insert(0, widget.friend!.uid);
    } else {
      Global.talkList.value.insert(0, widget.friend!.uid);
    }
  }

  dynamic changeRemark() {
    if (remarkTextEditingController.text == "") {
      EasyLoading.showToast("未输入备注",
          duration: const Duration(milliseconds: 200));
      return;
    }
    HttpUtil httpUtil = HttpUtil();
    httpUtil.postRequest(
        BaseUrl.friendBaseUrl,
        api[ApiType.updateRemark],
        {
          "owner_id": Global.userInfo.uid,
          "friend_id": widget.friend!.uid,
          "comment": remarkTextEditingController.text,
        },
        success: (data) {
          Friends tmp = HiveDbUtil.friendsClient.get(widget.friend!.uid);
          tmp.remark = remarkTextEditingController.text;
          HiveDbUtil.friendsClient.put(widget.friend!.uid, tmp);
        },
        fail: (data) => {});
  }

  dynamic deleteFriend() {
    HttpUtil httpUtil = HttpUtil();
    httpUtil.postRequest(
        BaseUrl.friendBaseUrl,
        api[ApiType.deleteFriend],
        {
          "owner_id": Global.userInfo.uid,
          "friend_id": widget.friend!.uid,
        },
        success: (data) {
          HiveDbUtil.friendsClient.delete(widget.friend!.uid);
          HiveDbUtil.chatsClient.delete(widget.friend!.uid);
          HiveDbUtil.friendReqReceiveClient.delete(widget.friend!.uid);
          HiveDbUtil.tmpFriendsClient.delete(widget.friend!.uid);
          Global.talkList.value.remove(widget.friend!.uid);
          Global.lastChooseMsgToUid = "";
        },
        fail: (data) => {});
  }
}
