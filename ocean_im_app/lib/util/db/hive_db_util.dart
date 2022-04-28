import 'package:hive_flutter/hive_flutter.dart';
import 'package:ocean_im/model/chat_message_model.dart';
import 'package:ocean_im/model/friend_model.dart';
import 'package:ocean_im/model/friend_receive_req_model.dart';
import 'package:ocean_im/model/tmp_friend_model.dart';
import 'package:ocean_im/model/user_info_model.dart';
import 'package:ocean_im/util/directory_util/file_util.dart';

class HiveDbUtil {
  static HiveDbUtil hiveDbUtil = HiveDbUtil();
  static late Box userClient;
  static late Box friendsClient;
  static late Box chatsClient;
  static late Box friendReqReceiveClient;
  static late Box tmpFriendsClient;

  static Future<void> install(String id) async {
    var dbPath = '${await FileUtil.fileUtil.localDocumentPath}' "/oceanim/" +
        id +
        "/databases";
    FileUtil.fileUtil.createFileFolder(dbPath);
    Hive.init(dbPath);
    if (!Hive.isAdapterRegistered(0)) {
      Hive.registerAdapter(UserInfoAdapter());
    }
    if (!Hive.isAdapterRegistered(2)) {
      Hive.registerAdapter(FriendsAdapter());
    }
    if (!Hive.isAdapterRegistered(1)) {
      Hive.registerAdapter(ChatsAdapter());
    }
    if (!Hive.isAdapterRegistered(3)) {
      Hive.registerAdapter(DetaillistAdapter());
    }
    if (!Hive.isAdapterRegistered(4)) {
      Hive.registerAdapter(ReqreceiveAdapter());
    }
    if (!Hive.isAdapterRegistered(5)) {
      Hive.registerAdapter(TmpFriendModelAdapter());
    }
  }

  static Future<HiveDbUtil> getInstance() async {
    userClient = await Hive.openBox<UserInfo>("userbox");
    friendsClient = await Hive.openBox<Friends>("friendbox");
    chatsClient = await Hive.openBox<Chats>("chatbox");
    friendReqReceiveClient = await Hive.openBox<Req_receive>("friendreqbox");
    tmpFriendsClient = await Hive.openBox<TmpFriendModel>("tmpfriendbox");

    await Hive.initFlutter();
    return hiveDbUtil;
  }

  static Future closeInstance() async {
    userClient.close();
    friendsClient.close();
    chatsClient.close();
    friendReqReceiveClient.close();
    tmpFriendsClient.close();
  }
}
