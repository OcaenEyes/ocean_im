import 'package:flutter/foundation.dart';
import 'package:ocean_im/model/chat_message_model.dart';

class MessageProvider extends ChangeNotifier {
  Map<String, List<Detail_list>> allMsgs = {};

  Map<String, List<Detail_list>> get all => allMsgs;

  setAllMsg(Map<String, List<Detail_list>> tmp) {
    allMsgs = tmp;
    notifyListeners();
  }
}
