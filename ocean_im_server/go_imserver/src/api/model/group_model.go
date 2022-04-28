package model

import (
	"fmt"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/common/db"
	"time"
)

type Group struct {
	GroupId      string    `db:"group_id" json:"group_id"`
	GroupName    string    `db:"group_name" json:"group_name"`
	Introduction string    `db:"introduction" json:"introduction"`
	Notification string    `db:"notification" json:"notification"`
	FaceImg      string    `db:"face_img" json:"face_img"`
	CreateTime   time.Time `db:"create_time" json:"create_time"`
	UpdateTime   time.Time `db:"update_time" json:"update_time"`
	Ex           string    `db:"ex" json:"ex"`
}

type GroupTmp struct {
	GroupId      string   ` json:"group_id"`
	GroupName    string   `json:"group_name"`
	Introduction string   ` json:"introduction"`
	Notification string   ` json:"notification"`
	FaceImg      string   ` json:"face_img"`
	Ex           string   `json:"ex"`
	CreatorUid   string   `json:"creator_uid"`
	MemberUid    []string `json:"member_uid"`
}

type GroupMember struct {
	GroupId            string    `db:"group_id" json:"group_id"`
	UserId             string    `db:"user_id" json:"user_id"`
	NickName           string    `db:"nick_name" json:"nick_name"`
	AdministratorLevel int       `db:"administrator_level" json:"administrator_level"`
	JoinTime           time.Time `db:"join_time" json:"join_time"`
	UserGroupFaceUrl   string    `db:"user_group_face_url" json:"user_group_face_url"`
}

type GroupRequest struct {
	ID               string    `db:"id" json:"id"`
	GroupId          string    `db:"group_id" json:"group_id"`
	FromUserId       string    `db:"from_user_id" json:"from_user_id"`
	ToUserId         string    `db:"to_user_id" json:"to_user_id"`
	Flag             int       `db:"flag" json:"flag"`
	ReqMessage       string    `db:"req_message" json:"req_message"`
	HandledMsg       string    `db:"handled_msg" json:"handled_msg"`
	CreateTime       time.Time `db:"create_time" json:"create_time"`
	FromUserNickName string    `db:"from_user_nick_name" json:"from_user_nick_name"`
	ToUserNickName   string    `db:"to_user_nick_name" json:"to_user_nick_name"`
	FromUserHeadImg  string    `db:"from_user_head_img" json:"from_user_head_img"`
	ToUserHeadImg    string    `db:"to_user_head_img" json:"to_user_head_img"`
	HandledUser      string    `db:"handled_user" json:"handled_user"`
}

func CheckGroupIsExit(groupId string) bool {
	db.InitMysqlDB()
	var tmpId string
	_ = db.Db.QueryRow("select group_id from group where group_id=?", groupId).Scan(&tmpId)
	defer db.Db.Close()
	if tmpId != "" {
		return true
	} else {
		return false
	}
}

func SaveGroup(group *Group) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("INSERT INTO group(group_id,group_name,introduction,notification,face_img,create_time,update_time,ex) values (?,?,?,?,?,?,?,?)",
		&group.GroupId,
		&group.GroupName,
		&group.Introduction,
		&group.Notification,
		&group.FaceImg,
		&group.CreateTime,
		&group.UpdateTime,
		&group.Ex,
	)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func UpdateGroup() {

}

func CheckIsInGroup(userId string, groupId string) bool {
	db.InitMysqlDB()
	var member GroupMember
	_ = db.Db.QueryRow("select user_id,group_id from group_member where group_id=? and user_id=?", groupId, userId).Scan(
		&member.UserId,
		&member.GroupId)
	defer db.Db.Close()
	if member.UserId != "" {
		return true
	} else {
		return false
	}
}

func JoinGroup(member *GroupMember) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("INSERT INTO group_member(group_id,user_id,nick_name,user_group_face_url,administrator_level,join_time) values (?,?,?,?,?,?)",
		&member.GroupId,
		&member.UserId,
		&member.NickName,
		&member.UserGroupFaceUrl,
		&member.AdministratorLevel,
		&member.JoinTime,
	)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func DeleteGroupMember(userId string, groupId string) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("DELETE from group_member where group_id=? and user_id=?", groupId, userId)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func SaveGroupRequest(request *GroupRequest) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("insert into group_request(id,group_id,from_user_id,from_user_nick_name,from_user_head_img,to_user_id,to_user_nick_name,to_user_head_img,flag,req_message,handled_user,handled_msg,create_time) values (?,?,?,?,?,?,?,?,?,?,?,?,?)",
		&request.ID,
		&request.GroupId,
		&request.FromUserId,
		&request.FromUserNickName,
		&request.FromUserHeadImg,
		&request.ToUserId,
		&request.ToUserNickName,
		&request.ToUserHeadImg,
		&request.Flag,
		&request.ReqMessage,
		&request.HandledUser,
		&request.HandledMsg,
		&request.CreateTime,
	)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func UpdateGroupRequest(request *GroupRequest) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("update group_request set flag=? and handled_user=? and handled_msg=? where from_user_id=? and group_id=? ",
		&request.Flag,
		&request.HandledUser,
		&request.HandledMsg,
		&request.FromUserId,
		&request.GroupId)

	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func DeleteGroupRequest(request *GroupRequest) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("delete from group_request  where from_user_id=? and group_id=? ",
		&request.FromUserId,
		&request.GroupId)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}
