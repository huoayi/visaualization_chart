package logic

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
	"visualization_chart/config"
	"visualization_chart/global"
)


func Init(file string) {
	readConfig(file)
	InitMysql()
}
func readConfig(file string) {
	//导入配置文件
	global.Config = &config.Server{}
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		glog.Errorf("导入配置文件失败，err msg:[%v]",err)
	}
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile,global.Config)
	if err != nil {
		glog.Errorf("读取配置文件入结构体失败，err msg:[%v]",err)
	}
	fmt.Println(global.Config.Db.Mysql.Database)
}


func InitMysql() {
	var err error
	dbInfo := global.Config.Db
	var(

		userName = global.Config.Db.Mysql.UserName
		address = global.Config.Db.Mysql.Address
		password = global.Config.Db.Mysql.Password
		port = global.Config.Db.Mysql.Port
		database = global.Config.Db.Mysql.Database
	)
	path := strings.Join([]string{userName,":",password,"@tcp(",address,":",port,")/",database,"?charset=utf8mb4&parseTime=True&loc=Local"},"")
	//DB,err = gorm.Open(mysql.Open(path),&gorm.Config{})
	db,err := gorm.Open(mysql.Open(path),&gorm.Config{})
	fmt.Println(path)
	if err != nil{
		glog.Errorf("数据库连接失败")
	}
	glog.Info("数据库连接成功")
	fmt.Println(dbInfo.Mysql.DbTable)
	global.DB.Tables = dbInfo.Mysql.DbTable
	global.DB.DataBase = db
}
