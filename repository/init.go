package repository

import (
	"strings"
	"sync"

	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

const (
	dbUserName = "root"
	dbPasswd   = "zz19980722"
	dbIPAddr   = "49.232.87.168"
	dbPort     = "3306"
	dbName     = "dousheng"
)

func init() {
	GetDB()
	initTable()
}

func GetDB() *gorm.DB {
	once.Do(func() {
		//change your dsn
		//mysqlDsn := "root:zz19980722@tcp(49.232.87.168:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
		mysqlDsn := strings.Join([]string{dbUserName, ":", dbPasswd, "@tcp(", dbIPAddr, ":", dbPort, ")/", dbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

		db, err := gorm.Open(mysql.Open(mysqlDsn))
		DB = db
		if err != nil {
			panic("mysql init err")
		}
	})
	return DB
}

func initTable() {
	err := DB.AutoMigrate(model.User{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(model.Video{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(model.Comment{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(model.Favorite{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(model.UserRelation{})
	if err != nil {
		panic(err)
	}
}
