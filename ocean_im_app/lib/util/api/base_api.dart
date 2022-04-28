import 'package:ocean_im/util/env.dart';

class BaseUrl {
  static const userBaseUrl = devBaseUrl + "/user/" + apiVesion;
  static const friendBaseUrl = devBaseUrl + "/friend/" + apiVesion;
  static const msgBaseUrl = devBaseUrl + "/msg/" + apiVesion;
}

enum ApiType {
  registUrl,
  loginUrl,
  sendFriendReqUrl,
  dueFriendReqUrl,
  getUserInfo,
  getUserInfoByUid,
  getUserInfoByPhone,
  getAllFriends,
  getAllMsg,
  getAllFriendReqs,
  updateRemark,
  deleteFriend
}

const api = {
  ApiType.registUrl: "/register",
  ApiType.loginUrl: "/login",
  ApiType.sendFriendReqUrl: "/sendFriendReq",
  ApiType.dueFriendReqUrl: "/dueFriendReq",
  ApiType.getUserInfo: "/getUserInfo",
  ApiType.getUserInfoByUid: "/getUserInfoByUid",
  ApiType.getUserInfoByPhone: "/getUserInfoByPhone",
  ApiType.getAllFriends: "/getAllFriends",
  ApiType.getAllMsg: "/getAllMsg",
  ApiType.getAllFriendReqs: "/getAllFriendReqs",
  ApiType.deleteFriend: "/deleteFriend",
  ApiType.updateRemark: "/updateFriendRemark"
};
