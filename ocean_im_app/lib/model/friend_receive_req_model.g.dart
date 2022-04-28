// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'friend_receive_req_model.dart';

// **************************************************************************
// TypeAdapterGenerator
// **************************************************************************

class ReqreceiveAdapter extends TypeAdapter<Req_receive> {
  @override
  final int typeId = 4;

  @override
  Req_receive read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return Req_receive(
      fields[0] as String,
      fields[1] as String,
      fields[2] as int,
      fields[3] as String,
      fields[4] as String,
    );
  }

  @override
  void write(BinaryWriter writer, Req_receive obj) {
    writer
      ..writeByte(5)
      ..writeByte(0)
      ..write(obj.reqId)
      ..writeByte(1)
      ..write(obj.userId)
      ..writeByte(2)
      ..write(obj.flag)
      ..writeByte(3)
      ..write(obj.reqMessage)
      ..writeByte(4)
      ..write(obj.createTime);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is ReqreceiveAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

FriendReceiveReqModel _$FriendReceiveReqModelFromJson(
        Map<String, dynamic> json) =>
    FriendReceiveReqModel(
      json['code'] as int,
      json['message'] as String,
      (json['req_receive'] as List<dynamic>)
          .map((e) => Req_receive.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$FriendReceiveReqModelToJson(
        FriendReceiveReqModel instance) =>
    <String, dynamic>{
      'code': instance.code,
      'message': instance.message,
      'req_receive': instance.reqReceive,
    };

Req_receive _$Req_receiveFromJson(Map<String, dynamic> json) => Req_receive(
      json['req_id'] as String,
      json['user_id'] as String,
      json['flag'] as int,
      json['req_message'] as String,
      json['create_time'] as String,
    );

Map<String, dynamic> _$Req_receiveToJson(Req_receive instance) =>
    <String, dynamic>{
      'req_id': instance.reqId,
      'user_id': instance.userId,
      'flag': instance.flag,
      'req_message': instance.reqMessage,
      'create_time': instance.createTime,
    };
