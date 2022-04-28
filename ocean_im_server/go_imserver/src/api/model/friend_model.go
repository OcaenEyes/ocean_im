package model

import (
	"fmt"
	"oceanim/go_server/src/api/enum"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/common/db"
	"time"
)

type Friend struct {
	OwnerId    string    `db:"owner_id" json:"owner_id"`
	FriendId   string    `db:"friend_id" json:"friend_id"`
	Comment    string    `db:"comment" json:"comment"`
	FriendFlag int       `db:"friend_flag" json:"friend_flag"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}

type FriendInfo struct {
	Uid      string `db:"uid" json:"uid"`
	Remark   string `db:"remark" json:"remark"`
	NickName string `db:"nick_name" json:"nickname"`
	Phone    string `db:"phone" json:"phone"`
	Email    string `db:"email" json:"email"`
	Birth    string `db:"birth" json:"birth"`
	HeadImg  string `db:"head_img" json:"head_img"`
	Gender   int  `db:"gender" json:"gender"`
	Ex       string `db:"ex" json:"ex"`
}

type FriendRequest struct {
	ReqId      string    `db:"req_id" json:"req_id"`
	UserId     string    `db:"user_id" json:"user_id"`
	Flag       int       `db:"flag" json:"flag"`
	ReqMessage string    `db:"req_message" json:"req_message"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}

func CheckIsFriend(friend *Friend) string {
	db.InitMysqlDB()
	var f1 Friend
	var f2 Friend
	err1 := db.Db.QueryRow("select owner_id,friend_id from friend where owner_id=? and friend_id=?",
		&friend.OwnerId,
		&friend.FriendId,
	).Scan(
		&f1.OwnerId,
		&f1.FriendId)
	err2 := db.Db.QueryRow("select owner_id,friend_id from friend where owner_id=? and friend_id=?",
		&friend.OwnerId,
		&friend.FriendId,
	).Scan(&f2.OwnerId,
		&f2.FriendId)
	db.Db.Close()

	if fmt.Sprint(err1)=="sql: no rows in result set" && fmt.Sprint(err2) =="sql: no rows in result set" {
		return "false"
	}else if err1 == nil  || err1 ==nil{
		return "true"
	}else {
		return "error"
	}
}

func SaveFriendRequest(request *FriendRequest) error {
	db.InitMysqlDB()
	var friendRequest FriendRequest
	err := db.Db.QueryRow("SELECT req_id,user_id,flag,create_time from friend_request where req_id=? and user_id=?", request.ReqId, request.UserId).Scan(
		&friendRequest.ReqId,
		&friendRequest.UserId,
		&friendRequest.Flag,
		&friendRequest.CreateTime)
	if err != nil {
		if fmt.Sprint(err) == "sql: no rows in result set" {
			_, err = db.Db.Exec("INSERT INTO friend_request(req_id,user_id,flag,req_message,create_time) values(?,?,?,?,?)",
				&request.ReqId,
				&request.UserId,
				&request.Flag,
				&request.ReqMessage,
				&request.CreateTime)
			if err != nil {
				fmt.Println(err)
				return errorcode.DataErr
			}
		}
	} else {
		_, err := db.Db.Exec("UPDATE friend_request set flag=? and create_time =? where req_id=? and user_id=?",
			&request.Flag,
			&request.CreateTime,
			&request.ReqId,
			&request.UserId,
		)
		if err != nil {
			fmt.Println(err)
			return errorcode.DataErr
		}
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func DeleteFriendReq(request *FriendRequest) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("DELETE from friend_request where req_id=? and user_id=?", &request.ReqId, &request.UserId)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func UpdateFriendRequest(request *FriendRequest) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("UPDATE friend_request set flag=? and create_time =? where req_id=? and user_id=?",
		&request.Flag,
		&request.CreateTime,
		&request.ReqId,
		&request.UserId,
	)
	_, err = db.Db.Exec("UPDATE friend_request set flag=? and create_time =? where req_id=? and user_id=?",
		&request.Flag,
		&request.CreateTime,
		&request.UserId,
		&request.ReqId,
	)

	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func QueryFriendRequestsSend(uid string) ([]FriendRequest, error) {
	db.InitMysqlDB()
	var friendReqs []FriendRequest
	rows, err := db.Db.Query("select req_id,user_id,flag,req_message,create_time from friend_request where req_id=?", uid)
	if err != nil {
		fmt.Println(err)
		if fmt.Sprint(err) == "sql: no rows in result set" {
			return nil, errorcode.FriendReqSendsNo
		}
		return nil, errorcode.DataErr
	}
	defer db.Db.Close()
	for rows.Next() {
		var req FriendRequest
		err = rows.Scan(&req.ReqId, &req.UserId, &req.Flag, &req.ReqMessage, &req.CreateTime)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.DataErr
		}
		friendReqs = append(friendReqs, req)
	}
	return friendReqs, errorcode.Ok
}

func QueryFriendRequestsReceive(uid string) ([]FriendRequest, error) {
	db.InitMysqlDB()
	var friendReqs []FriendRequest
	rows, err := db.Db.Query("select req_id,user_id,flag,req_message,create_time from friend_request where user_id=?", uid)
	if err != nil {
		fmt.Println(err)
		if fmt.Sprint(err) == "sql: no rows in result set" {
			return nil, errorcode.FriendReqReceiveNo
		}
		return nil, errorcode.DataErr
	}

	defer db.Db.Close()
	for rows.Next() {
		var req FriendRequest
		err = rows.Scan(&req.ReqId, &req.UserId, &req.Flag, &req.ReqMessage, &req.CreateTime)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.DataErr
		}
		friendReqs = append(friendReqs, req)
	}
	return friendReqs, errorcode.Ok
}

func SaveFriend(friend *Friend) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("INSERT INTO friend(owner_id,friend_id,comment,friend_flag,create_time) values(?,?,?,?,?),(?,?,?,?,?)",
		&friend.OwnerId, &friend.FriendId, &friend.Comment, &friend.FriendFlag, &friend.CreateTime,
		&friend.FriendId, &friend.OwnerId, &friend.Comment, &friend.FriendFlag, &friend.CreateTime)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	_, err = db.Db.Exec("update friend_request set flag =? where req_id=? and user_id=?",
		enum.Accept.Flag,
		&friend.FriendId,
		&friend.OwnerId)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return errorcode.Ok
}

func QueryFriends(uid string) ([]Friend, error) {
	db.InitMysqlDB()
	var friends []Friend
	rows, err := db.Db.Query("select owner_id,friend_id,comment,friend_flag,create_time from friend where friend_flag=1 and owner_id=?", uid)
	if err != nil {
		fmt.Println(err)
		if fmt.Sprint(err) == "sql: no rows in result set" {
			return nil, errorcode.FriendsNo
		}
		return nil, errorcode.DataErr
	}
	defer db.Db.Close()
	for rows.Next() {
		var friend Friend
		err = rows.Scan(&friend.OwnerId, &friend.FriendId, &friend.Comment, &friend.FriendFlag, &friend.CreateTime)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.DataErr
		}
		friends = append(friends, friend)
	}
	return friends, nil
}

func UpdateFriendComment(friend *Friend) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("update friend set comment=? where owner_id=? and friend_id=?",
		&friend.Comment,
		&friend.OwnerId,
		&friend.FriendId,
	)
	if err != nil {
		fmt.Println(err)
		return errorcode.DataErr
	}
	db.Db.Close()
	return errorcode.Ok
}

func DeleteFriend(friend *Friend) error {
	db.InitMysqlDB()
	_, err1 := db.Db.Exec("delete from friend where owner_id=? and friend_id=?",
		&friend.OwnerId,
		&friend.FriendId,
	)
	_, err2 := db.Db.Exec("delete from friend where owner_id=? and friend_id=?",
		&friend.FriendId,
		&friend.OwnerId,
	)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return errorcode.DataErr
	}

	var f1 FriendRequest
	err1 = db.Db.QueryRow("select req_id,user_id from friend_request where req_id=? and user_id=?",
		&friend.FriendId,
		&friend.OwnerId,
	).Scan(&f1.ReqId, &f1.UserId)
	if err1 != nil {
		if fmt.Sprint(err1) != "sql: no rows in result set" {
			return errorcode.DataErr
		}
	} else {
		_, err := db.Db.Exec("DELETE from friend_request where req_id=? and user_id=?", &f1.ReqId, &f1.UserId)
		if err != nil {
			fmt.Println(err)
			return errorcode.DataErr
		}
	}
	var f2 FriendRequest
	err2 = db.Db.QueryRow("select req_id,user_id from friend_request where req_id=? and user_id=?",
		&friend.OwnerId,
		&friend.FriendId,
	).Scan(&f2.ReqId, &f2.UserId)
	if err2 != nil {
		if fmt.Sprint(err2) != "sql: no rows in result set" {
			return errorcode.DataErr
		}
	} else {
		_, err := db.Db.Exec("DELETE from friend_request where req_id=? and user_id=?", &f2.ReqId, &f2.UserId)
		if err != nil {
			fmt.Println(err)
			return errorcode.DataErr
		}
	}
	db.Db.Close()
	return errorcode.Ok

}
