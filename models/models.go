package models

import (
	"blog/pkg/set"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	Uuid      uuid.UUID `gorm:"primary_key" json:"id"`
	CreatedAt int32     `json:"createdAt"`
	UpdatedOn int32     `json:"updatedAt"`
}

func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	id, _ := uuid.NewV4()
	err := scope.SetColumn("Uuid", id)
	if err != nil {
		fmt.Println("uuid err: ", err)
		return err
	}
	err = scope.SetColumn("CreatedAt", time.Now().Unix())
	return err
}

func (m *Model) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("UpdatedOn", time.Now().Unix())
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := set.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//禁用表名为复数的形式
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&User{}, &Role{}, &Module{}, &Api{}, &Article{}, &Category{}, &Tag{})
}

func CloseDB() {
	defer db.Close()
}
