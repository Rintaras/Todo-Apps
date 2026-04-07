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
	
func BuildDBConfig() *DBConfig {
	portstr := strings.TrimSpace(os.Getenv("MYSQL_PORT"))
	port, _ := strconv.Atoi(portstr)

 dbConfig := DBConfig{
  Host: strings.TrimSpace(os.Getenv("MYSQL_HOST")),
  Port: port,
  User: strings.TrimSpace(os.Getenv("MYSQL_USER")),
  Password: os.Getenv("MYSQL_PASSWORD"),
  DBName: strings.TrimSpace(os.Getenv("MYSQL_DATABASE")),
 }
 return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
 return fmt.Sprintf(
  "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
  dbConfig.User,
  dbConfig.Password,
  dbConfig.Host,
  dbConfig.Port,
  dbConfig.DBName,
 )
}