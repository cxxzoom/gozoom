package main

import (
	"encoding/json"
	"os"
)

// TODO 从数据库里获取量表数据并解析

type gaugeCnf struct {
	Name []string `json:"name"`
}

// parseGaugeJson 解析json文件并获取待处理的量表
func parseGaugeJson() (*gaugeCnf, error) {
	path := "./gauge.json"
	jsonCtx, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ctx := gaugeCnf{}

	err = json.Unmarshal(jsonCtx, &ctx)
	if err != nil {
		return nil, err
	}

	return &ctx, err
}
