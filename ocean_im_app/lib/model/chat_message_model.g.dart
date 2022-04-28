// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'chat_message_model.dart';

// **************************************************************************
// TypeAdapterGenerator
// **************************************************************************

class ChatsAdapter extends TypeAdapter<Chats> {
  @override
  final int typeId = 1;

  @override
  Chats read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return Chats(
      (fields[1] as List).cast<Detail_list>(),
      fields[0] as String,
    );
  }

  @override
  void write(BinaryWriter writer, Chats obj) {
    writer
      ..writeByte(2)
      ..writeByte(1)
      ..write(obj.detailList)
      ..writeByte(0)
      ..write(obj.toUser);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is ChatsAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}

class DetaillistAdapter extends TypeAdapter<Detail_list> {
  @override
  final int typeId = 3;

  @override
  Detail_list read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return Detail_list(
      fields[0] as String,
      fields[1] as String,
      fields[2] as String,
      fields[3] as String,
      fields[4] as int,
      fields[5] as int,
      fields[6] as int,
      fields[7] as String,
    );
  }

  @override
  void write(BinaryWriter writer, Detail_list obj) {
    writer
      ..writeByte(8)
      ..writeByte(0)
      ..write(obj.msgId)
      ..writeByte(1)
      ..write(obj.sendId)
      ..writeByte(2)
      ..write(obj.receiveId)
      ..writeByte(3)
      ..write(obj.sendTime)
      ..writeByte(4)
      ..write(obj.signType)
      ..writeByte(5)
      ..write(obj.msgType)
      ..writeByte(6)
      ..write(obj.contentType)
      ..writeByte(7)
      ..write(obj.content);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is DetaillistAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ChatMessageModel _$ChatMessageModelFromJson(Map<String, dynamic> json) =>
    ChatMessageModel(
      (json['chats'] as List<dynamic>)
          .map((e) => Chats.fromJson(e as Map<String, dynamic>))
          .toList(),
      json['code'] as int,
      json['message'] as String,
    );

Map<String, dynamic> _$ChatMessageModelToJson(ChatMessageModel instance) =>
    <String, dynamic>{
      'chats': instance.chats,
      'code': instance.code,
      'message': instance.message,
    };

Chats _$ChatsFromJson(Map<String, dynamic> json) => Chats(
      (json['detail_list'] as List<dynamic>)
          .map((e) => Detail_list.fromJson(e as Map<String, dynamic>))
          .toList(),
      json['to_user'] as String,
    );

Map<String, dynamic> _$ChatsToJson(Chats instance) => <String, dynamic>{
      'detail_list': instance.detailList,
      'to_user': instance.toUser,
    };

Detail_list _$Detail_listFromJson(Map<String, dynamic> json) => Detail_list(
      json['msg_id'] as String,
      json['send_id'] as String,
      json['receive_id'] as String,
      json['send_time'] as String,
      json['sign_type'] as int,
      json['msg_type'] as int,
      json['content_type'] as int,
      json['content'] as String,
    );

Map<String, dynamic> _$Detail_listToJson(Detail_list instance) =>
    <String, dynamic>{
      'msg_id': instance.msgId,
      'send_id': instance.sendId,
      'receive_id': instance.receiveId,
      'send_time': instance.sendTime,
      'sign_type': instance.signType,
      'msg_type': instance.msgType,
      'content_type': instance.contentType,
      'content': instance.content,
    };
