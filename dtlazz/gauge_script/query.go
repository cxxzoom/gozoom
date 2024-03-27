package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// TODO 组装数据并发起请求

type login struct {
	account  string
	password string
}

type queryJson struct {
	Uri       string `json:"uri"`        // 请求的的地址
	LoginUri  string `json:"login_uri"`  // 登录地址
	CommitUri string `json:"commit_uri"` // 提交地址
}

var qJson queryJson

func getCookie(login2 *login, cnf *queryJson) (string, error) {
	data := []byte(fmt.Sprintf(`{"userName": %s, "password": %s,"type":1}`, (*login2).account,
		(*login2).password))
	request, err := http.NewRequest(http.MethodPost, cnf.Uri+cnf.LoginUri, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Host", cnf.Uri)
	request.Header.Add("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	cookie := resp.Cookies()
	//fmt.Println(cookie[0])
	defer resp.Body.Close()
	var cookieStr string
	for i, cookie := range cookie {
		if i > 0 {
			cookieStr += "; "
		}
		cookieStr += cookie.Name + "=" + cookie.Value
	}

	return cookieStr, nil
}

func commit(account string, payload2 *payload) (interface{}, error) {
	err := parseQueryJson()
	if err != nil {
		return nil, err
	}

	login2 := &login{
		account:  account,
		password: "123456",
	}

	cookie, err := getCookie(login2, &qJson)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(*payload2)

	request, err := http.NewRequest(http.MethodPost, qJson.Uri+qJson.CommitUri, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	//fmt.Println(qJson.Uri+qJson.CommitUri, string(data), cookie)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Host", qJson.Uri)
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Cookie", cookie)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// TODO 这里开始发起请求
	return nil, err
}

func parseQueryJson() error {
	filePath := "./query.json"
	jsonCtx, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonCtx, &qJson)
	if err != nil {
		return err
	}

	return nil
}
