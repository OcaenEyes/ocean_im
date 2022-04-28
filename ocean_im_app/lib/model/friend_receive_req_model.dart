import 'package:hive_flutter/hive_flutter.dart';
import 'package:json_annotation/json_annotation.dart';

part 'friend_receive_req_model.g.dart';

@JsonSerializable()
class FriendReceiveReqModel extends Object {
  @JsonKey(name: 'code')
  int code;

  @JsonKey(name: 'message')
  String message;

  @JsonKey(name: 'req_receive')
  List<Req_receive> reqReceive;

  FriendReceiveReqModel(
    this.code,
    this.message,
    this.reqReceive,
  );

  factory FriendReceiveReqModel.fromJson(Map<String, dynamic> srcJson) =>
      _$FriendReceiveReqModelFromJson(srcJson);

  Map<String, dynamic> toJson() => _$FriendReceiveReqModelToJson(this);
}

@JsonSerializable()
@HiveType(typeId: 4)
class Req_receive extends Object {
  @JsonKey(name: 'req_id')
  @HiveField(0)
  String reqId;

  @JsonKey(name: 'user_id')
  @HiveField(1)
  String userId;

  @JsonKey(name: 'flag')
  @HiveField(2)
  int flag;

  @JsonKey(name: 'req_message')
  @HiveField(3)
  String reqMessage;

  @JsonKey(name: 'create_time')
  @HiveField(4)
  String createTime;

  Req_receive(
    this.reqId,
    this.userId,
    this.flag,
    this.reqMessage,
    this.createTime,
  );

  factory Req_receive.fromJson(Map<String, dynamic> srcJson) =>
      _$Req_receiveFromJson(srcJson);

  Map<String, dynamic> toJson() => _$Req_receiveToJson(this);
}
