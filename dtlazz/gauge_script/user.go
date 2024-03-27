package main

import (
	"encoding/json"
	"os"
)

// TODO 获取用户数据

type userJson struct {
	Account []string `json:"account"`
}

// 解析并获取所属有用户
func parseUserJson() (*userJson, error) {
	filePath := "./user.json"
	jsonCtx, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cnf := userJson{}
	err = json.Unmarshal(jsonCtx, &cnf)
	if err != nil {
		return nil, err
	}

	return &cnf, nil
}
