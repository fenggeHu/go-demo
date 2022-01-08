package service

import (
	"errors"
	"log"
)

func Login(loginName string, pwd string) (bool, error) {
	log.Printf("username=%s, password=%s", loginName, pwd)
	if len(loginName)<5 || len(pwd) <8 {
		return false, errors.New("用户名或密码错误")
	}
	return true, nil
}
