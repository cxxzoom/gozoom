package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	User     string `json:"user"`     // JSON键 "user" 映射到此字段
	Password string `json:"password"` // JSON键 "password" 映射到此字段
	Host     string `json:"host"`     // JSON键 "host" 映射到此字段
	Port     string `json:"port"`     // JSON键 "port" 映射到此字段
	DbName   string `json:"dbname"`   // JSON键 "dbname" 映射到此字段
}

type mysql struct {
	*sql.DB
}

// 连接数据库
func mysqlConn(config config) (*mysql, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &mysql{db}, nil
}

func (m *mysql) updateUsers() {
	sql := "UPDATE judge_user set passwd = 'Adc@123' WHERE passwd = 123456"
	_, err := m.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

const configPath = "config.json"

func main() {
	count := 0

	cnf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysqlConn(*cnf)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println(" 1 min modify wake password...")
	// 定时执行修改密码的操作
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			// 更新密码
			log.Println(" working...")
			db.updateUsers()
			count++
			clearScreen(count)
		}
	}

}

// getConfig 获取数据库配置文件
func getConfig() (*config, error) {
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cnf := config{}

	// 解析 JSON 数据到 user 变量
	err = json.Unmarshal(jsonData, &cnf)
	if err != nil {
		return nil, err
	}

	return &cnf, nil
}

// 清除屏幕
func clearScreen(count int) {
	if count%20 != 0 {
		return
	}

	//switch runtime.GOOS {
	//case "windows":
	//	cmd := exec.Command("cmd", "/c", "cls")
	//	cmd.Stdout = os.Stdout
	//	cmd.Run()
	//case "linux", "darwin":
	//	cmd := exec.Command("clear")
	//	cmd.Stdout = os.Stdout
	//	cmd.Run()
	//default:
	//	fmt.Println("Unsupported operating system")
	//}
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	log.Println("1 min modify wake password...")
}
