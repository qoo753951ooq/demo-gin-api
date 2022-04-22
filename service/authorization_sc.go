package service

import (
	"demo-gin-api/model/vo"
	"demo-gin-api/service/authorization"
)

func PostUserLogin(body vo.UserPostVO) (string, error) {
	token, err := authorization.UserLogin(body)
	return token, err
}
