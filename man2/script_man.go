package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/exec"
	"time"
)

type config struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
}

type mysql struct {
	*sql.DB
}

// 连接数据库
func mysqlConn(config config) (*mysql, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.user, config.password, config.host, config.port, config.dbname)
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

func main() {
	count := 0
	cnf := config{
		user:     "root",
		password: "Cqzlyy@cqqy#8888",
		host:     "localhost",
		port:     "3306",
		dbname:   "NuoHeTest_db",
	}

	db, err := mysqlConn(cnf)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println(" 1 min modify wake password...")
	// 定时执行修改密码的操作
	ticker := time.NewTicker(9 * time.Second)
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
	cmd.Run()
	log.Println("1 min modify wake password...")
}
