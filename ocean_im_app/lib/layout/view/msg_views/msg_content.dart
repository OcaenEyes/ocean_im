import 'package:demoji/demoji.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/model/chat_message_model.dart';
import 'package:ocean_im/util/db/hive_db_util.dart';

class MsgContent extends StatefulWidget {
  final String uid;
  final List<Detail_list>? msgList;
  final TextEditingController st;
  const MsgContent(
      {Key? key, required this.uid, required this.msgList, required this.st})
      : super(key: key);

  @override
  _MsgContentState createState() => _MsgContentState();
}

class _MsgContentState extends State<MsgContent> {
  ScrollController scrollController = ScrollController();
  // bool showBottomButton = false;
  bool isShowEmoji = false;
  @override
  void initState() {
    widget.st.clear();
    super.initState();
    // scrollController.addListener(() {
    //   debugPrint("scrollController.offset.toString()");
    //   debugPrint(scrollController.offset.toString());

    //   debugPrint("scrollController.position.maxScrollExtent.toString()");
    //   debugPrint(scrollController.position.maxScrollExtent.toString());

    //   if (scrollController.hasClients) {
    //     if (scrollController.offset ==
    //         scrollController.position.maxScrollExtent) {
    //       showBottomButton = false;
    //     }
    //   } else if (scrollController.offset <
    //       scrollController.position.maxScrollExtent - 10) {
    //     showBottomButton = true;
    //   }
    // });
  }

  @override
  void dispose() {
    scrollController.dispose();
    widget.st.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Global.lastChooseMsgToUid != ""
        ? Column(
            children: [
              Container(
                height: 62,
                alignment: Alignment.centerLeft,
                padding: const EdgeInsets.fromLTRB(14, 0, 0, 0),
                width: MediaQuery.of(context).size.width - 356,
                color: Colors.white,
                child: Text(
                  HiveDbUtil.friendsClient.containsKey(widget.uid)
                      ? (HiveDbUtil.friendsClient.get(widget.uid).remark != ""
                          ? HiveDbUtil.friendsClient.get(widget.uid).remark
                          : HiveDbUtil.friendsClient.get(widget.uid).nickname)
                      : "",
                  style: const TextStyle(
                      fontSize: 18, fontWeight: FontWeight.w600),
                ),
              ),
              Container(
                  height: MediaQuery.of(context).size.height - 246,
                  width: MediaQuery.of(context).size.width - 356,
                  color: const Color.fromARGB(255, 235, 235, 235),
                  child: Stack(
                    children: [
                      widget.uid != "" ? msglist(widget.msgList) : Container(),
                      // Positioned(
                      //     bottom: 0,
                      //     right: (MediaQuery.of(context).size.width - 380) / 2,
                      //     child: Container(
                      //       child: showBottomButton
                      //           ? TextButton(
                      //               onPressed: () {
                      //                 scrollController.jumpTo(
                      //                     scrollController.position.maxScrollExtent);
                      //                 showBottomButton = false;
                      //               },
                      //               child: Row(children: const [
                      //                 Icon(
                      //                   Icons.arrow_drop_down,
                      //                   color: Colors.black54,
                      //                 ),
                      //                 Text(
                      //                   "回到底部",
                      //                   style: TextStyle(
                      //                       fontSize: 10, color: Colors.black54),
                      //                 )
                      //               ]),
                      //             )
                      //           : null,
                      //     ))
                      Positioned(
                          bottom: 0,
                          left: 4,
                          child: isShowEmoji ? emojiWidget() : Container()),
                    ],
                  )),
              Container(
                height: 184,
                width: MediaQuery.of(context).size.width - 356,
                color: Colors.white,
                child: Column(children: [
                  Container(
                    color: Colors.orange[50],
                    child: Row(children: [
                      SizedBox(
                        width: 40,
                        height: 40,
                        child: GestureDetector(
                            onTap: () {
                              debugPrint(Demoji.smile);
                              setState(() {
                                isShowEmoji = !isShowEmoji;
                              });
                            },
                            child: const Icon(Icons.emoji_emotions)),
                      )
                    ]),
                  ),
                  Stack(children: [
                    TextField(
                      cursorColor: Colors.black,
                      cursorWidth: 0.2,
                      autofocus: true,
                      controller: widget.st,
                      keyboardType: TextInputType.multiline,
                      enableInteractiveSelection: true,
                      maxLines: 6,
                      minLines: 6,
                      decoration: const InputDecoration(
                          border: InputBorder.none,
                          focusedBorder: InputBorder.none),
                    ),
                    Positioned(
                      bottom: 10,
                      right: 20,
                      child: GestureDetector(
                        onTap: () {
                          // debugPrint(textEditingController.text);
                          sendMsg();
                        },
                        child: Container(
                          width: 56,
                          height: 36,
                          alignment: Alignment.center,
                          decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(2),
                              color: const Color.fromARGB(31, 131, 131, 131)),
                          child: const Text(
                            "发送",
                            style: TextStyle(
                                color: Colors.black,
                                fontSize: 16,
                                fontWeight: FontWeight.w500),
                          ),
                        ),
                      ),
                    )
                  ])
                ]),
              ),
            ],
          )
        : Container();
  }

  dynamic emojiWidget() {
    List tmp = [
      Demoji.one_hundred,
      Demoji.grinning,
      Demoji.grimacing,
      Demoji.grin,
      Demoji.joy,
      Demoji.smiley,
      Demoji.smile,
      Demoji.sweat_smile,
      Demoji.laughing,
      Demoji.innocent,
      Demoji.upside_down_face,
      Demoji.yum,
      Demoji.kissing,
      Demoji.kissing_heart,
      Demoji.kissing_closed_eyes,
      Demoji.clown_face,
      Demoji.sunglasses,
      Demoji.stuck_out_tongue,
      Demoji.cowboy_hat_face,
      Demoji.no_mouth,
      Demoji.hand_over_mouth,
      Demoji.symbols_over_mouth,
      Demoji.worried,
      Demoji.angry,
      Demoji.rage,
      Demoji.pensive,
      Demoji.confused,
      Demoji.pleading,
      Demoji.triumph,
      Demoji.cold_sweat,
      Demoji.fearful,
      Demoji.frowning,
      Demoji.cry,
      Demoji.cold,
      Demoji.hot,
      Demoji.zipper_mouth_face,
      Demoji.nauseated_face,
      Demoji.mask,
      Demoji.vomiting,
      Demoji.zzz,
      Demoji.sleeping,
      Demoji.skull
    ];
    return Container(
      width: 336,
      height: 300,
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(8),
          color: const Color.fromARGB(255, 224, 224, 224)),
      child: GridView.builder(
        gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
            maxCrossAxisExtent: 48),
        itemBuilder: (BuildContext context, int index) {
          return emojiitem(tmp[index]);
        },
        itemCount: tmp.length,
      ),
    );
  }

  dynamic emojiitem(emoji) {
    return GestureDetector(
      onTap: () {
        isShowEmoji = false;
        widget.st.text = widget.st.text + emoji;
      },
      child: Container(
          alignment: Alignment.center,
          child: Text(
            emoji,
            style: const TextStyle(fontSize: 24),
          )),
    );
  }

  dynamic msglist(List<Detail_list>? tmp) {
    // if (tmp != null) {
    //   // Timer(
    //   //     const Duration(milliseconds: 500),
    //   //     () => scrollController.positions.isNotEmpty
    //   //         ? scrollController
    //   //             .jumpTo(scrollController.position.maxScrollExtent)
    //   //         : null);
    //   debugPrint("scrollController.offset.toString()");
    //   debugPrint(scrollController.offset.toString());

    //   debugPrint("scrollController.position.maxScrollExtent.toString()");
    //   debugPrint(scrollController.position.maxScrollExtent.toString());

    //   if (scrollController.hasClients) {
    //     if (scrollController.offset ==
    //         scrollController.position.maxScrollExtent) {
    //       setState(() {
    //         showBottomButton = false;
    //       });
    //     } else if (scrollController.offset <
    //         scrollController.position.maxScrollExtent - 10) {
    //       setState(() {
    //         showBottomButton = true;
    //       });
    //     }
    //   }
    // } else {
    //   tmp = [];
    // }
    tmp ??= [];

    List<Detail_list> tempp = [];
    for (var element in tmp) {
      if (element.signType == 0 && element.sendId != Global.userInfo.uid) {
        element.signType = 1;
        tempp.add(element);
        Global.socket.emit("readMsg", element.msgId);
      } else {
        tempp.add(element);
      }
    }
    Chats ttmp = Chats(tempp, Global.lastChooseMsgToUid);
    HiveDbUtil.chatsClient.put(Global.lastChooseMsgToUid, ttmp);
    return ListView.builder(
      reverse: true,
      controller: scrollController,
      scrollDirection: Axis.vertical,
      shrinkWrap: true,
      itemBuilder: (context, index) => getRow(context, tmp![index]),
      itemCount: tmp.length,
    );
  }

  dynamic getRow(BuildContext context, Detail_list? detail) {
    return Container(
      margin: const EdgeInsets.fromLTRB(10, 8, 10, 5),
      child: detail!.sendId == Global.userInfo.uid
          ? Row(
              mainAxisAlignment: MainAxisAlignment.end,
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Container(
                  margin: const EdgeInsets.fromLTRB(0, 0, 6, 0),
                  padding: const EdgeInsets.all(8.0),
                  alignment: Alignment.center,
                  decoration: BoxDecoration(
                      color: Colors.green[200],
                      borderRadius: BorderRadius.circular(4)),
                  child: ConstrainedBox(
                      child: Text(
                        detail.content.trim(),
                        style: const TextStyle(fontSize: 16),
                      ),
                      constraints: BoxConstraints(
                          maxWidth: MediaQuery.of(context).size.width * 0.34)),
                ),
                // Text(
                //     HiveDbUtil.friendsClient.containsKey(detail.sendId)
                //         ? HiveDbUtil.friendsClient.get(detail.sendId).nickname
                //         : "",
                //     style: const TextStyle(fontSize: 10))
              ],
            )
          : Row(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                // Text(
                //     HiveDbUtil.friendsClient.containsKey(detail.sendId)
                //         ? HiveDbUtil.friendsClient.get(detail.sendId).nickname
                //         : "",
                //     style: const TextStyle(fontSize: 10)),
                Container(
                  margin: const EdgeInsets.fromLTRB(6, 0, 0, 0),
                  padding: const EdgeInsets.all(8.0),
                  alignment: Alignment.center,
                  decoration: BoxDecoration(
                      color: Colors.amber[200],
                      borderRadius: BorderRadius.circular(4)),
                  child: ConstrainedBox(
                      child: Text(
                        detail.content.trim(),
                        style: const TextStyle(fontSize: 16),
                      ),
                      constraints: BoxConstraints(
                          maxWidth: MediaQuery.of(context).size.width * 0.34)),
                ),
              ],
            ),
    );
  }

  dynamic sendMsg() {
    Map<String, dynamic> msg = {};
    Map<String, dynamic> msgData = {};
    msg["type"] = 0;
    msgData["content"] = widget.st.text;
    msgData["send_id"] = Global.userInfo.uid;
    msgData["receive_id"] = widget.uid;
    msgData["platform"] = "windows";
    msgData["group_id"] = "";
    msgData["content_type"] = 3;
    msg["data"] = msgData;
    // if (!Global.socket.connected) {
    //   Global.socket.connect();
    //   Global.socket.emit(
    //       "register", {"uid": Global.userInfo.uid, "platform": "windows"});
    // }
    Global.socket.emitWithAck("send", msg, ack: () {});
    Global.socket.on("toClient",
        (data) => {debugPrint("toclient了消息"), debugPrint(data.toString())});

    Map<String, dynamic> mapTmp = {};
    mapTmp["send_id"] = msgData["send_id"];
    mapTmp["receive_id"] = msgData["receive_id"];
    mapTmp["content_type"] = msgData["content_type"];
    mapTmp["content"] = msgData["content"];
    mapTmp["msg_id"] = "";
    if (widget.uid == Global.userInfo.uid) {
      mapTmp["sign_type"] = 1;
    } else {
      mapTmp["sign_type"] = 0;
    }

    mapTmp["msg_type"] = 3;
    mapTmp["send_time"] = DateTime.now().toString();
    Detail_list detail = Detail_list.fromJson(mapTmp);

    if (HiveDbUtil.chatsClient.containsKey(detail.receiveId)) {
      Chats chats = HiveDbUtil.chatsClient.get(detail.receiveId);
      chats.detailList.insert(0, detail);
      HiveDbUtil.chatsClient.put(detail.receiveId, chats);
    } else {
      List<Detail_list> tmpList = [];
      Chats chats = Chats(tmpList, detail.receiveId);
      tmpList.add(detail);
      HiveDbUtil.chatsClient.put(detail.receiveId, chats);
    }

    setState(() {
      widget.st.clear();
    });

    Global.socket.on("sendRes", (data) {
      debugPrint("这次收到了消息");
      debugPrint(data.toString());
      if (data["code"] == "30002") {
        EasyLoading.showToast(data["send_res"],
            duration: const Duration(seconds: 5));
        // HiveDbUtil.friendsClient.delete(Global.lastChooseMsgToUid);
        // HiveDbUtil.chatsClient.delete(Global.lastChooseMsgToUid);
        // HiveDbUtil.friendReqReceiveClient.delete(Global.lastChooseMsgToUid);
        // Global.talkList.value.remove(Global.lastChooseMsgToUid);
        // Global.lastChooseMsgToUid = "";
      }
    });

    // final _channel = WebSocketChannel.connect(
    //   Uri.parse('ws://127.0.0.1:8080/ws'),
    // );
    // _channel.sink.add(textEditingController.text);
  }
}
