import 'package:flutter/material.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/layout/view/friend_views/friend_content.dart';
import 'package:ocean_im/layout/view/friend_views/friend_tool.dart';
import 'package:ocean_im/model/friend_model.dart';

class FriendView extends StatefulWidget {
  const FriendView({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return _FriendViewState();
  }
}

class _FriendViewState extends State<FriendView> {
  List<Friends> friends = [];
  late ScrollController scrollController;
  String userid = "";
  bool showFriendTool = false;
  @override
  void initState() {
    super.initState();
    if (!Global.socket.connected) {
      Global.socket.connect();
      Global.socket.emit(
          "register", {"uid": Global.userInfo.uid, "platform": "windows"});
    }
    scrollController = ScrollController();
  }

  @override
  void dispose() {
    super.dispose();
    scrollController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return ValueListenableBuilder(
        valueListenable: Hive.box<Friends>("friendbox").listenable(),
        builder: (context, Box<Friends> friends, _) {
          return Row(
            children: [
              Container(
                width: 300,
                color: const Color.fromARGB(136, 245, 244, 244),
                child: Column(children: [
                  ListTile(
                    onTap: () {
                      setState(() {
                        userid = "";
                        showFriendTool = true;
                      });
                    },
                    hoverColor: Colors.green[50],
                    leading: const Icon(Icons.person_add),
                    title: const Text(
                      "添加好友",
                      style: TextStyle(fontSize: 14),
                    ),
                    trailing: mounted
                        ? ValueListenableBuilder(
                            valueListenable: Global.countFriendReqs,
                            builder: ((context, int value, _) {
                              return Container(
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
                                        borderRadius: BorderRadius.circular(6))
                                    : null,
                              );
                            }))
                        : const Text(""),
                  ),
                  friends.keys.isNotEmpty
                      ? SizedBox(
                          child: ListView.builder(
                            controller: scrollController,
                            itemBuilder: (BuildContext context, int position) {
                              return getRow(
                                  context, friends.values.toList()[position]);
                            },
                            itemCount: friends.length,
                          ),
                          width: 300,
                          height: MediaQuery.of(context).size.height - 50,
                          // color: const Color.fromARGB(137, 228, 228, 228),
                        )
                      : Container(
                          width: 300,
                          alignment: Alignment.topCenter,
                          child: const Text("em..没有朋友呢"),
                          // color: const Color.fromARGB(137, 228, 228, 228),
                        )
                ]),
              ),
              Container(
                color: Colors.white30,
                width: MediaQuery.of(context).size.width - 356,
                child: userid != ""
                    ? FriendContent(friend: friends.get(userid))
                    : (showFriendTool ? const FriendTool() : null),
              )
            ],
          );
        });
  }

  dynamic getRow(BuildContext context, Friends friend) {
    return Container(
      child: ListTile(
        onTap: () {
          setState(() {
            userid = friend.uid;
            showFriendTool = false;
          });
        },
        leading: Container(
          width: 34,
          height: 34,
          decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(6),
              image: const DecorationImage(
                  image: AssetImage("assets/images/3.jpg"), fit: BoxFit.fill)),
        ),
        title: Text(friend.remark != "" ? friend.remark : friend.nickname),
        // subtitle: Text(friend.uid),
        // focusColor: Colors.blue[100],
        // selectedColor: Colors.blue[100],
        hoverColor: Colors.green[50],
      ),
      color: userid == friend.uid ? Colors.green[100] : Colors.transparent,
    );
  }
}
