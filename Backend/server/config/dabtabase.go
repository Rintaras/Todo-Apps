package Config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() (*DBConfig, error) {
	portstr := strings.TrimSpace(os.Getenv("MYSQL_PORT"))
	port := 3306
	if portstr != "" {
		p, err := strconv.Atoi(portstr)
		if err != nil || p <= 0 {
			return nil, fmt.Errorf("MYSQL_PORT が無効です: %q", portstr)
		}
		port = p
	}

	cfg := DBConfig{
		Host:     strings.TrimSpace(os.Getenv("MYSQL_HOST")),
		Port:     port,
		User:     strings.TrimSpace(os.Getenv("MYSQL_USER")),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   strings.TrimSpace(os.Getenv("MYSQL_DATABASE")),
	}

	var miss []string
	if cfg.User == "" {
		miss = append(miss, "MYSQL_USER")
	}
	if cfg.Host == "" {
		miss = append(miss, "MYSQL_HOST")
	}
	if cfg.DBName == "" {
		miss = append(miss, "MYSQL_DATABASE")
	}
	if len(miss) > 0 {
		return nil, fmt.Errorf(
			"MySQL 用の環境変数が不足しています: %v（シェルで export するか、リポジトリルートまたは Backend に .env を置いてください）",
			miss,
		)
	}

	return &cfg, nil
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}