import 'package:flutter/material.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/layout/view/msg_views/msg_content.dart';
import 'package:ocean_im/model/chat_message_model.dart';
import 'package:ocean_im/model/friend_model.dart';
import 'package:ocean_im/util/db/hive_db_util.dart';

class MsgView extends StatefulWidget {
  const MsgView({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return _MsgViewState();
  }
}

class _MsgViewState extends State<MsgView> {
  ScrollController scrollController = ScrollController();
  TextEditingController sendTextEditingController = TextEditingController();
  bool isDrag = false;
  @override
  void initState() {
    super.initState();
    sendTextEditingController.clear();
  }

  @override
  void dispose() {
    scrollController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return ValueListenableBuilder(
        valueListenable: Hive.box<Chats>("chatbox").listenable(),
        builder: (context, Box<Chats> chatsBox, _) {
          return Row(
            children: [
              chatsBox.isOpen
                  ? (chatsBox.keys.isNotEmpty
                      ? Expanded(
                          child: ValueListenableBuilder(
                          valueListenable: Global.talkList,
                          builder: ((cpntext, List v, _) {
                            return v.isNotEmpty
                                ? Container(
                                    width: 300,
                                    color: const Color.fromARGB(
                                        136, 245, 244, 244),
                                    child: Align(
                                      alignment: Alignment.topCenter,
                                      child: ReorderableListView.builder(
                                        buildDefaultDragHandles: isDrag,

                                        onReorder:
                                            (int oldIndex, int newIndex) {
                                          debugPrint("oldIndex.toString()");
                                          debugPrint(oldIndex.toString());
                                          debugPrint("newIndex.toString()");
                                          debugPrint(newIndex.toString());
                                          if (oldIndex < newIndex) {
                                            newIndex -= 1;
                                          }

                                          String tmpStr = Global.talkList.value
                                              .removeAt(oldIndex);
                                          Global.talkList.value
                                              .insert(newIndex, tmpStr);
                                          setState(() {
                                            isDrag = false;
                                          });
                                        },
                                        // reverse: true,
                                        shrinkWrap: true,
                                        scrollDirection: Axis.vertical,
                                        scrollController: scrollController,

                                        itemBuilder: (BuildContext context,
                                            int position) {
                                          // return ValueListenableBuilder(
                                          //   valueListenable: Global.talkList,
                                          //   builder: ((cpntext, List v, _) {
                                          return getRow(
                                              context,
                                              chatsBox.get(v[position]),
                                              position);
                                          //   }),
                                          // );
                                        },
                                        itemCount: chatsBox.length,
                                      ),
                                    ),
                                  )
                                : Container(
                                    width: 300,
                                    alignment: Alignment.topCenter,
                                    child: const Text("em..暂无消息呢"),
                                    color: const Color.fromARGB(
                                        137, 228, 228, 228),
                                  );
                          }),
                        ))
                      : Container(
                          width: 300,
                          alignment: Alignment.topCenter,
                          padding: const EdgeInsets.fromLTRB(0, 10, 0, 0),
                          child: const Text("em..暂无消息呢"),
                          color: const Color.fromARGB(137, 228, 228, 228),
                        ))
                  : Container(
                      width: 300,
                      alignment: Alignment.topCenter,
                      child: const Text("em..暂无消息呢"),
                      padding: const EdgeInsets.fromLTRB(0, 10, 0, 0),
                      color: const Color.fromARGB(137, 228, 228, 228),
                    ),
              Container(
                color: Colors.white30,
                width: MediaQuery.of(context).size.width - 356,
                child: Global.lastChooseMsgToUid != ""
                    ? MsgContent(
                        st: sendTextEditingController,
                        uid: Global.lastChooseMsgToUid,
                        msgList:
                            chatsBox.get(Global.lastChooseMsgToUid)?.detailList)
                    : null,
              )
            ],
          );
        });
  }

  dynamic getRow(BuildContext context, Chats? tmp, int i) {
    int noRead = 0;
    for (var element in tmp!.detailList) {
      if (element.signType == 0 &&
          element.sendId != Global.userInfo.uid &&
          Global.lastChooseMsgToUid != tmp.toUser) {
        noRead += 1;
      }
    }

    return ValueListenableBuilder(
        key: ValueKey(i),
        valueListenable: Hive.box<Friends>("friendbox").listenable(),
        builder: (context, Box<Friends> friendbox, _) {
          return friendbox.containsKey(tmp.toUser)
              ? Container(
                  child: ListTile(
                    onLongPress: () {
                      setState(() {
                        isDrag = true;
                      });
                    },
                    onTap: () {
                      // debugPrint(tmp.keys.toList().toString());

                      setState(() {
                        Global.lastChooseMsgToUid = tmp.toUser;
                        sendTextEditingController.clear();
                      });

                      List<Detail_list> tempp = [];
                      for (var element in tmp.detailList) {
                        if (element.signType == 0 &&
                            element.sendId != Global.userInfo.uid) {
                          element.signType = 1;
                          tempp.add(element);
                          Global.countNoRead.value -= 1;
                          Global.socket.emit("readMsg", element.msgId);
                        } else if (element.signType == 0 &&
                            element.sendId == Global.userInfo.uid &&
                            element.receiveId == Global.userInfo.uid) {
                          element.signType = 1;
                          tempp.add(element);
                          Global.countNoRead.value -= 1;
                          Global.socket.emit("readMsg", element.msgId);
                        } else {
                          tempp.add(element);
                        }
                      }
                      Chats ttmp = Chats(tempp, tmp.toUser);
                      HiveDbUtil.chatsClient.put(tmp.toUser, ttmp);
                    },
                    leading: Stack(
                      children: [
                        Container(
                          width: 40,
                          height: 40,
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
                                (noRead > 0) ? noRead.toString() : "",
                                style: const TextStyle(
                                  color: Colors.white,
                                  fontSize: 10,
                                ),
                              ),
                              decoration: (noRead > 0)
                                  ? BoxDecoration(
                                      color: Colors.red[400],
                                      borderRadius: BorderRadius.circular(6))
                                  : null,
                            ))
                      ],
                    ),
                    title: Text(friendbox.get(tmp.toUser)!.remark != ""
                        ? friendbox.get(tmp.toUser)!.remark
                        : friendbox.get(tmp.toUser)!.nickname),
                    subtitle: Text(
                        tmp.detailList.isNotEmpty & tmp.detailList.isNotEmpty
                            ? tmp.detailList.first.content.trim()
                            : "",
                        style: const TextStyle(
                            color: Colors.black54,
                            fontSize: 12,
                            letterSpacing: 1.5,
                            overflow: TextOverflow.ellipsis)),
                    trailing: Text(
                      tmp.detailList.isNotEmpty & tmp.detailList.isNotEmpty
                          ? DateTime.parse(tmp.detailList.first.sendTime)
                                  .month
                                  .toString() +
                              "/" +
                              DateTime.parse(tmp.detailList.first.sendTime)
                                  .day
                                  .toString()
                          : "",
                      style:
                          const TextStyle(color: Colors.black38, fontSize: 11),
                    ),
                    hoverColor: Colors.green[50],
                    // selectedColor: Colors.blue[50],
                    // focusColor: Colors.blue[50],
                  ),
                  color: Global.lastChooseMsgToUid == tmp.toUser
                      ? Colors.green[100]
                      : Colors.transparent,
                )
              : Container(
                  width: 300,
                  padding: const EdgeInsets.fromLTRB(0, 10, 0, 0),
                  alignment: Alignment.topCenter,
                  child: const Text("em..暂无消息呢"),
                  // color: const Color.fromARGB(137, 228, 228, 228),
                );
        });
  }
}
