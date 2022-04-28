package model

import (
	"fmt"
	"math/rand"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/common/db"
	"time"
)

type User struct {
	Uid        string    `db:"uid" json:"uid"`
	NickName   string    `db:"nick_name" json:"nickname"`
	Phone      string    `db:"phone" json:"phone"`
	PassWord   string    `db:"pass_word" json:"password"`
	Email      string    `db:"email" json:"email"`
	Birth      string    `db:"birth" json:"birth"`
	HeadImg    string    `db:"head_img" json:"head_img"`
	Gender     int     `db:"gender" json:"gender"`
	Ex         string    `db:"ex" json:"ex"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

type UserDetail struct {
	Uid      string `db:"uid" json:"uid"`
	NickName string `db:"nick_name" json:"nickname"`
	Phone    string `db:"phone" json:"phone"`
	Email    string `db:"email" json:"email"`
	Birth    string `db:"birth" json:"birth"`
	HeadImg  string `db:"head_img" json:"head_img"`
	Gender   int  `db:"gender" json:"gender"`
	Ex       string `db:"ex" json:"ex"`
}

func SaveUser(user *User) (string, error) {
	db.InitMysqlDB()
	var oldUid string
	var e error
	var uid string
	for fmt.Sprint(e) != "sql: no rows in result set" {
		uid = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		e = db.Db.QueryRow("select uid from user where uid =?", uid).Scan(&oldUid)
		fmt.Println(e)
	}
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	_, e = db.Db.Exec("INSERT INTO user(uid,nick_name,phone,pass_word,email,birth,head_img,gender,ex,create_time,update_time) values(?,?,?,?,?,?,?,?,?,?,?)",
		uid,
		&user.Phone,
		&user.Phone,
		&user.PassWord,
		&user.Email,
		&user.Birth,
		&user.HeadImg,
		&user.Gender,
		&user.Ex,
		&user.CreateTime,
		&user.UpdateTime,
	)
	if e != nil {
		fmt.Println("exec failed, ", e)
		return "", errorcode.DataErr
	}
	defer db.Db.Close()
	return uid, errorcode.Ok

}

func UserInfoByPhone(phone string) (*User, error) {
	db.InitMysqlDB()
	var user User
	fmt.Println(phone)
	err := db.Db.QueryRow("select uid,nick_name,phone,email,pass_word,birth,head_img,gender,ex from user where phone=? ", phone).Scan(
		&user.Uid,
		&user.NickName,
		&user.Phone,
		&user.Email,
		&user.PassWord,
		&user.Birth,
		&user.HeadImg,
		&user.Gender,
		&user.Ex,
	)
	if err != nil {
		fmt.Println("exec failed, ", err)
		if fmt.Sprint(err) == "sql: no rows in result set" {
			defer db.Db.Close()
			return nil, errorcode.UserNotExist
		} else {
			defer db.Db.Close()
			return nil, errorcode.DataErr
		}
	}
	fmt.Println(user)
	defer db.Db.Close()
	return &user, nil
}

func UserInfoByUid(uid string) (*User, error) {
	db.InitMysqlDB()
	var user User
	fmt.Println(uid)
	err := db.Db.QueryRow("select uid,nick_name,phone,email,pass_word,birth,head_img,gender,ex from user where uid=? ", uid).Scan(
		&user.Uid,
		&user.NickName,
		&user.Phone,
		&user.Email,
		&user.PassWord,
		&user.Birth,
		&user.HeadImg,
		&user.Gender,
		&user.Ex,
	)
	if err != nil {
		fmt.Println("exec failed, ", err)
		if fmt.Sprint(err) == "sql: no rows in result set" {
			defer db.Db.Close()
			return nil, errorcode.UserNotExist
		} else {
			defer db.Db.Close()
			return nil, errorcode.DataErr
		}
	}
	fmt.Println(user)
	defer db.Db.Close()
	return &user, nil
}

func UpdateUserInfo(user *User) {
	db.InitMysqlDB()
	_, err := db.Db.Exec("update user set nick_name=?,email=? where phone=?", &user.NickName, &user.Email, &user.Phone)
	if err != nil {
		fmt.Println("sql exec error", err)
		return
	}
	defer db.Db.Close()
	return
}
