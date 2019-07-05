package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

type Path struct {
	ID      uint     `gorm:"primary_key" json:"id"`
	Path    string   `gorm:"size:255;unique_index:unique_path_index" json:"path"`
	Desc    string   `gorm:"size:255" json:"desc"`
	Methods []Method `json:"methods"`
}

type Method struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	PathId   uint   `json:"pathId"`
	Method   string `gorm:"size:30" json:"method"`
	Response string `sql:"type:text;" json:"response"`
}

const dbName = "mockHttp.db"

var DB *gorm.DB

func InitDb() error {
	dbPath := viper.GetString("sqliteDbPath")
	if dbPath == "" {
		homeDir, _ := os.UserHomeDir()
		dbPath = path.Join(homeDir, dbName)
	}

	dbFile := path.Join(dbPath, dbName)
	logrus.Infof("db file path is %s", dbFile)

	d, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		logrus.Errorf("err to create or connect to %s from %s", dbName, dbFile)
		panic("load db file error")
		return err
	}
	DB = d

	autoMerge()

	return nil
}

func autoMerge() {
	DB.LogMode(true)
	DB.AutoMigrate(&Path{}, &Method{})
}
