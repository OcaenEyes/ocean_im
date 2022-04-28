import 'package:hive/hive.dart';
import 'package:json_annotation/json_annotation.dart';

part 'user_info_model.g.dart';

@JsonSerializable()
class UserInfoModel extends Object {
  @JsonKey(name: 'code')
  int code;

  @JsonKey(name: 'message')
  String message;

  @JsonKey(name: 'userInfo')
  UserInfo userInfo;

  UserInfoModel(
    this.code,
    this.message,
    this.userInfo,
  );

  factory UserInfoModel.fromJson(Map<String, dynamic> srcJson) =>
      _$UserInfoModelFromJson(srcJson);

  Map<String, dynamic> toJson() => _$UserInfoModelToJson(this);
}

@JsonSerializable()
@HiveType(typeId: 0)
class UserInfo extends Object {
  @JsonKey(name: 'uid')
  @HiveField(0)
  String uid;

  @JsonKey(name: 'nickname')
  @HiveField(1)
  String nickname;

  @JsonKey(name: 'phone')
  @HiveField(2)
  String phone;

  @JsonKey(name: 'email')
  @HiveField(3)
  String email;

  @JsonKey(name: 'birth')
  @HiveField(4)
  String birth;

  @JsonKey(name: 'head_img')
  @HiveField(5)
  String headImg;

  @JsonKey(name: 'gender')
  @HiveField(6)
  int gender;

  @JsonKey(name: 'ex')
  @HiveField(7)
  String ex;

  UserInfo(
    this.uid,
    this.nickname,
    this.phone,
    this.email,
    this.birth,
    this.headImg,
    this.gender,
    this.ex,
  );

  factory UserInfo.fromJson(Map<String, dynamic> srcJson) =>
      _$UserInfoFromJson(srcJson);

  Map<String, dynamic> toJson() => _$UserInfoToJson(this);
}
