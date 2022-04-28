import 'package:hive/hive.dart';
import 'package:json_annotation/json_annotation.dart';

part 'friend_model.g.dart';

@JsonSerializable()
class FriendModel extends Object {
  @JsonKey(name: 'code')
  int code;

  @JsonKey(name: 'friends')
  List<Friends> friends;

  @JsonKey(name: 'message')
  String message;

  FriendModel(
    this.code,
    this.friends,
    this.message,
  );

  factory FriendModel.fromJson(Map<String, dynamic> srcJson) =>
      _$FriendModelFromJson(srcJson);

  Map<String, dynamic> toJson() => _$FriendModelToJson(this);
}

@HiveType(typeId: 2)
@JsonSerializable()
class Friends extends Object {
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

  Friends(
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

  factory Friends.fromJson(Map<String, dynamic> srcJson) =>
      _$FriendsFromJson(srcJson);

  Map<String, dynamic> toJson() => _$FriendsToJson(this);
}
