package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

// DONE 创建数据连接

type config struct {
	User      string `json:"user"`      // JSON键 "user" 映射到此字段
	Password  string `json:"password"`  // JSON键 "password" 映射到此字段
	Host      string `json:"host"`      // JSON键 "host" 映射到此字段
	Port      string `json:"port"`      // JSON键 "port" 映射到此字段
	DbName    string `json:"dbname"`    // JSON键 "dbname" 映射到此字段
	CycleTime int    `json:"cycleTime"` // JSON键 "cycleTime" 映射到此字段
}

type gaugeRes struct {
	Id      int         // 量表id
	Options []gaugeOpts // 量表的题目以及选项
}

type gaugeOpts struct {
	questionId int    // 题目id
	answerStr  string // 选项列表
	answers    []string
}

type userOpts struct {
	Id      int
	Account string
}

// mysqlConn 连接数据库
func mysqlConn(config config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// 测试创建连接是否成功
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// getGaugeInfo 从数据库里获取量表信息
func getGaugeInfo(titles gaugeCnf) (*map[int]gaugeRes, error) {
	res := make(map[int]gaugeRes, len(titles.Name))
	cnf, err := parseMysqlJson()
	if err != nil {
		return nil, err
	}

	conn, err := mysqlConn(*cnf)
	if err != nil {
		return nil, err
	}

	for _, name := range titles.Name {
		sqlCtx := `
		SELECT m.id as measure_id,q.id as question_id, q.options
		FROM qy_measure m
		LEFT JOIN qy_measure_question q ON m.id = q.measure_id
		WHERE m.status = 1 AND m.measure_title = ?`

		rows, err := conn.Query(sqlCtx, name)

		if err != nil {
			return nil, err
		}

		tmp := gaugeRes{}
		opts := gaugeOpts{}

		for rows.Next() {
			err := rows.Scan(&tmp.Id, &opts.questionId, &opts.answerStr)
			if err != nil {
				return nil, err
			}
			tmp.Options = append(tmp.Options, opts)
		}

		res[tmp.Id] = tmp

	}

	return &res, nil
}

// getUserInfo 获取用户信息
func getUserInfo(accounts *userJson) (*[]userOpts, error) {
	cnf, err := parseMysqlJson()
	if err != nil {
		return nil, err
	}

	conn, err := mysqlConn(*cnf)
	if err != nil {
		return nil, err
	}

	placeholders := make([]string, len((*accounts).Account))
	args := make([]interface{}, len((*accounts).Account))

	for i, acc := range (*accounts).Account {
		placeholders[i] = "?"
		args[i] = acc
	}

	query := fmt.Sprintf("SELECT id, account FROM qy_user WHERE account IN (%s)", strings.Join(placeholders, ","))
	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var res []userOpts
	opts := userOpts{}
	for rows.Next() {
		err := rows.Scan(&opts.Id, &opts.Account)
		if err != nil {
			return nil, err
		}

		res = append(res, opts)
	}

	return &res, nil
}

// parseMysqlJson 解析获取数据库的配置
func parseMysqlJson() (*config, error) {
	filePath := "./config.json"
	jsonCtx, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	cnf := config{}
	err = json.Unmarshal(jsonCtx, &cnf)
	if err != nil {
		return nil, err
	}

	return &cnf, err
}
