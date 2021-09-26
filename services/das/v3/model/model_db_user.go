package model

import (
	"encoding/json"

	"strings"
)

// 注册的数据库用户信息
type DbUser struct {
	// 数据库用户ID

	DbUserId string `json:"db_user_id"`
	// 数据库用户名称

	DbUsername string `json:"db_username"`
}

func (o DbUser) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DbUser struct{}"
	}

	return strings.Join([]string{"DbUser", string(data)}, " ")
}