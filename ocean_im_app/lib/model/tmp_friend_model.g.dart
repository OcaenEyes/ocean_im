// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'tmp_friend_model.dart';

// **************************************************************************
// TypeAdapterGenerator
// **************************************************************************

class TmpFriendModelAdapter extends TypeAdapter<TmpFriendModel> {
  @override
  final int typeId = 5;

  @override
  TmpFriendModel read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return TmpFriendModel(
      fields[0] as String,
      fields[1] as String,
      fields[2] as String,
      fields[3] as String,
      fields[4] as String,
      fields[5] as String,
      fields[6] as String,
      fields[7] as int,
      fields[8] as String,
    );
  }

  @override
  void write(BinaryWriter writer, TmpFriendModel obj) {
    writer
      ..writeByte(9)
      ..writeByte(0)
      ..write(obj.uid)
      ..writeByte(1)
      ..write(obj.remark)
      ..writeByte(2)
      ..write(obj.nickname)
      ..writeByte(3)
      ..write(obj.phone)
      ..writeByte(4)
      ..write(obj.email)
      ..writeByte(5)
      ..write(obj.birth)
      ..writeByte(6)
      ..write(obj.headImg)
      ..writeByte(7)
      ..write(obj.gender)
      ..writeByte(8)
      ..write(obj.ex);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is TmpFriendModelAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

TmpFriendModel _$TmpFriendModelFromJson(Map<String, dynamic> json) =>
    TmpFriendModel(
      json['uid'] as String,
      json['remark'] as String,
      json['nickname'] as String,
      json['phone'] as String,
      json['email'] as String,
      json['birth'] as String,
      json['head_img'] as String,
      json['gender'] as int,
      json['ex'] as String,
    );

Map<String, dynamic> _$TmpFriendModelToJson(TmpFriendModel instance) =>
    <String, dynamic>{
      'uid': instance.uid,
      'remark': instance.remark,
      'nickname': instance.nickname,
      'phone': instance.phone,
      'email': instance.email,
      'birth': instance.birth,
      'head_img': instance.headImg,
      'gender': instance.gender,
      'ex': instance.ex,
    };
