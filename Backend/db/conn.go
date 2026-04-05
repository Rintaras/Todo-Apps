package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

// Conn は各 repository から利用する *sql.DB（init で一度だけオープンする）。
var Conn *sql.DB

func init() {
	user := strings.TrimSpace(os.Getenv("MYSQL_USER"))
	password := os.Getenv("MYSQL_PASSWORD")
	host := strings.TrimSpace(os.Getenv("MYSQL_HOST"))
	port := strings.TrimSpace(os.Getenv("MYSQL_PORT"))
	database := strings.TrimSpace(os.Getenv("MYSQL_DATABASE"))

	missing := []string{}
	if user == "" {
		missing = append(missing, "MYSQL_USER")
	}
	if host == "" {
		missing = append(missing, "MYSQL_HOST")
	}
	if port == "" {
		missing = append(missing, "MYSQL_PORT")
	}
	if database == "" {
		missing = append(missing, "MYSQL_DATABASE")
	}
	if len(missing) > 0 {
		log.Fatalf("db: 次の環境変数が未設定です: %v（`env-sample` を参照）", missing)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	var err error
	Conn, err = sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}

	Conn.SetMaxOpenConns(25)
	Conn.SetMaxIdleConns(5)

	if err := Conn.Ping(); err != nil {
		log.Fatalf("db: MySQL に接続できません (user=%s host=%s port=%s database=%s, passwordはログに出しません): %v",
			user, host, port, database, err)
	}
}
