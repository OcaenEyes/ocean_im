import 'package:hive_flutter/adapters.dart';
import 'package:json_annotation/json_annotation.dart';

part 'tmp_friend_model.g.dart';

@HiveType(typeId: 5)
@JsonSerializable()
class TmpFriendModel extends Object {
  @HiveField(0)
  @JsonKey(name: 'uid')
  String uid;

  @HiveField(1)
  @JsonKey(name: 'remark')
  String remark;

  @HiveField(2)
  @JsonKey(name: 'nickname')
  String nickname;

  @HiveField(3)
  @JsonKey(name: 'phone')
  String phone;

  @HiveField(4)
  @JsonKey(name: 'email')
  String email;

  @HiveField(5)
  @JsonKey(name: 'birth')
  String birth;

  @HiveField(6)
  @JsonKey(name: 'head_img')
  String headImg;

  @HiveField(7)
  @JsonKey(name: 'gender')
  int gender;

  @HiveField(8)
  @JsonKey(name: 'ex')
  String ex;

  TmpFriendModel(
    this.uid,
    this.remark,
    this.nickname,
    this.phone,
    this.email,
    this.birth,
    this.headImg,
    this.gender,
    this.ex,
  );

  factory TmpFriendModel.fromJson(Map<String, dynamic> srcJson) =>
      _$TmpFriendModelFromJson(srcJson);

  Map<String, dynamic> toJson() => _$TmpFriendModelToJson(this);
}
