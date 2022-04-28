package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/model"
	"oceanim/go_server/src/utils/auth"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	if (user.Phone == "") || (user.PassWord == "") {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.UserInfoEmpty.Code, "message": errorcode.UserInfoEmpty.Message})
		return
	}

	tmpUser, tmpErr := model.UserInfoByPhone(user.Phone)
	if tmpErr != nil {
		fmt.Println(tmpErr)
		tmpCode, tmpMsg := errorcode.ParserErr(tmpErr)
		if tmpCode == errorcode.UserNotExist.Code {
			var uid string
			uid, err = model.SaveUser(&user)
			code, msg := errorcode.ParserErr(err)
			if code == errorcode.Ok.Code {
				var myClaims auth.Claims
				myClaims.Uid = uid
				token, tokenErr := auth.CreateToken(&myClaims)
				if tokenErr != nil {
					c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
					return
				}
				c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "token": token})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"code": tmpCode, "message": tmpMsg})
			return
		}
	}
	if tmpUser != nil {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.UserIsExist.Code, "message": errorcode.UserIsExist.Message})
		return
	}
}

func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	if (user.Phone == "") || (user.PassWord == "") {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.UserInfoEmpty.Code, "message": errorcode.UserInfoEmpty.Message})
		return
	}
	tmpUser, err := model.UserInfoByPhone(user.Phone)
	if err != nil {
		fmt.Println(err)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		return
	}
	if tmpUser == nil {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.UserNotExist.Code, "message": errorcode.UserNotExist.Message})
		return
	}

	if tmpUser.PassWord == user.PassWord {
		var myClaims auth.Claims
		myClaims.Uid = tmpUser.Uid
		myClaims.Phone = tmpUser.Phone
		myClaims.Email = tmpUser.Email
		myClaims.NickName = tmpUser.NickName
		token, tokenErr := auth.CreateToken(&myClaims)
		if tokenErr != nil {
			c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "token": token})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.UserPassErr.Code, "message": errorcode.UserPassErr.Message})
		return
	}

}

func ChangePass(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {
	var userDetail model.UserDetail
	var token auth.Token
	err := c.ShouldBindJSON(&token)
	if err != nil {
		fmt.Println(err)
		return
	}
	uid, tokenErr := auth.ValidateToken(token.Token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	user, err := model.UserInfoByUid(uid)
	if err != nil {
		fmt.Println(err)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		return
	}
	userDetail.Uid = user.Uid
	userDetail.NickName = user.NickName
	userDetail.Phone = user.Phone
	userDetail.Email = user.Email
	userDetail.Birth = user.Birth
	userDetail.HeadImg = user.HeadImg
	userDetail.Gender = user.Gender
	userDetail.Ex = user.Ex
	c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "userInfo": userDetail})
	return

}

func GetUserInfoByUid(c *gin.Context) {
	var userDetail model.UserDetail
	var uidInfo map[string]string
	err :=c.ShouldBindJSON(&uidInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(c.GetHeader("Token"))
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}

	user, err := model.UserInfoByUid(uidInfo["uid"])
	if err != nil {
		fmt.Println(err)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		return
	}
	userDetail.Uid = user.Uid
	userDetail.NickName = user.NickName
	userDetail.Phone = user.Phone
	userDetail.Email = user.Email
	userDetail.Birth = user.Birth
	userDetail.HeadImg = user.HeadImg
	userDetail.Gender = user.Gender
	userDetail.Ex = user.Ex
	c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "userInfo": userDetail})
	return
}

func GetUserInfoByPhone(c *gin.Context) {
	var userDetail model.UserDetail
	var phoneInfo map[string]string
	err :=c.ShouldBindJSON(&phoneInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	uid, tokenErr := auth.ValidateToken(c.GetHeader("Token"))
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}

	user, err := model.UserInfoByPhone(phoneInfo["phone"])
	if err != nil {
		fmt.Println(err)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		return
	}
	userDetail.Uid = user.Uid
	userDetail.NickName = user.NickName
	userDetail.Phone = user.Phone
	userDetail.Email = user.Email
	userDetail.Birth = user.Birth
	userDetail.HeadImg = user.HeadImg
	userDetail.Gender = user.Gender
	userDetail.Ex = user.Ex
	c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "userInfo": userDetail})
	return
}

func UpdateUserInfo(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	model.UpdateUserInfo(&user)
	return
}
