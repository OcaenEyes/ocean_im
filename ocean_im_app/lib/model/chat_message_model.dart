import 'package:hive_flutter/adapters.dart';
import 'package:json_annotation/json_annotation.dart';

part 'chat_message_model.g.dart';

@JsonSerializable()
class ChatMessageModel extends Object {
  @JsonKey(name: 'chats')
  List<Chats> chats;

  @JsonKey(name: 'code')
  int code;

  @JsonKey(name: 'message')
  String message;

  ChatMessageModel(
    this.chats,
    this.code,
    this.message,
  );

  factory ChatMessageModel.fromJson(Map<String, dynamic> srcJson) =>
      _$ChatMessageModelFromJson(srcJson);

  Map<String, dynamic> toJson() => _$ChatMessageModelToJson(this);
}

@JsonSerializable()
@HiveType(typeId: 1)
class Chats extends Object {
  @JsonKey(name: 'detail_list')
  @HiveField(1)
  List<Detail_list> detailList;

  @JsonKey(name: 'to_user')
  @HiveField(0)
  String toUser;

  Chats(
    this.detailList,
    this.toUser,
  );

  factory Chats.fromJson(Map<String, dynamic> srcJson) =>
      _$ChatsFromJson(srcJson);

  Map<String, dynamic> toJson() => _$ChatsToJson(this);
}

@JsonSerializable()
@HiveType(typeId: 3)
class Detail_list extends Object {
  @JsonKey(name: 'msg_id')
  @HiveField(0)
  String msgId;

  @JsonKey(name: 'send_id')
  @HiveField(1)
  String sendId;

  @JsonKey(name: 'receive_id')
  @HiveField(2)
  String receiveId;

  @JsonKey(name: 'send_time')
  @HiveField(3)
  String sendTime;

  @JsonKey(name: 'sign_type')
  @HiveField(4)
  int signType;

  @JsonKey(name: 'msg_type')
  @HiveField(5)
  int msgType;

  @JsonKey(name: 'content_type')
  @HiveField(6)
  int contentType;

  @JsonKey(name: 'content')
  @HiveField(7)
  String content;

  Detail_list(
    this.msgId,
    this.sendId,
    this.receiveId,
    this.sendTime,
    this.signType,
    this.msgType,
    this.contentType,
    this.content,
  );

  factory Detail_list.fromJson(Map<String, dynamic> srcJson) =>
      _$Detail_listFromJson(srcJson);

  Map<String, dynamic> toJson() => _$Detail_listToJson(this);
}
