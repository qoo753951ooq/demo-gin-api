package authorization

import (
	"demo-gin-api/middleware"
	"demo-gin-api/model"
	"demo-gin-api/model/vo"
	"fmt"
)

func UserLogin(data vo.UserPostVO) (string, error) {

	if data.Account != "alan" || data.Password != "testPwd" {
		return model.Empty_string, fmt.Errorf("%s\n", model.Id_not_found_string)
	}

	token, err := middleware.GenerateToken(data.Account)

	if err != nil {
		return model.Empty_string, err
	}

	return token, nil
}
