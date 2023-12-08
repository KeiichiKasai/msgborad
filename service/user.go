package service

import (
	"errors"
	"fmt"
	"main.go/dao"
	"main.go/model"
	"main.go/utils"
	"strings"
)

func CheckUserIsExist(username string) bool {
	_, err := dao.GetPasswordByUsername(username)
	if err != nil {
		return false
	}
	return true
}
func CheckPassword(username, password string) bool {
	a, err := dao.GetPasswordByUsername(username)
	if err != nil {
		fmt.Println("找不到" + username + "的密码")
		return false
	}
	b, err := utils.Decrypt(a, utils.Key, utils.Iv)
	if err != nil {
		fmt.Println("对用户" + username + "的解密出错")
		return false
	}
	if password == b {
		return true
	} else {
		return false
	}
}
func CreateUser(username, password, phone string) error {
	a, err := utils.Encrypt(password, utils.Key, utils.Iv)
	if err != nil {
		fmt.Println(err)
		fmt.Println("对用户" + username + "的密码加密出错")
		return err
	}
	b, e := utils.Encrypt(phone, utils.Key, utils.Iv)
	if e != nil {
		fmt.Println(e)
		fmt.Println("对用户" + username + "的电话加密出错")
		return e
	}
	u := model.User{
		Username: username,
		Password: a,
		Phone:    b,
	}
	dao.InsertUser(u)
	return nil
}
func ChangePassword(username, password, phone string) error {
	data, err := dao.SearchUser(username)
	if err != nil {
		fmt.Println("调取用户信息失败")
		return err
	}
	temp := strings.Split(data, " ")
	p, err := utils.Decrypt(temp[2], utils.Key, utils.Iv)
	if err != nil {
		fmt.Println("对用户" + username + "手机号的解密出错")
		return err
	}
	if p == phone {
		a, er := utils.Encrypt(password, utils.Key, utils.Iv)
		if er != nil {
			fmt.Println(er)
			fmt.Println("对用户" + username + "的密码加密出错")
			return er
		}
		b, e := utils.Encrypt(phone, utils.Key, utils.Iv)
		if e != nil {
			fmt.Println(e)
			fmt.Println("对用户" + username + "的电话加密出错")
			return e
		}
		u := model.User{
			Username: username,
			Password: a,
			Phone:    b,
		}
		dao.UpdateUser(u)
		return nil
	}
	return errors.New("wrong phone")
}
