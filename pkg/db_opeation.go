package pkg

import (
	"fmt"
	"github.com/golang/glog"
	"gorm.io/gorm/clause"
	"visualization_chart/global"
)
//新建数据库表
func CreatTable(rows [][]string,filename string) {
	type Ind struct {
		Id int
	}
	ind := global.DB
	ind.DataBase.Migrator().CreateTable(&Ind{})
	for _, v := range rows[0]{
		table := clause.Table{Name:v}

		ind.DataBase.Exec("alter table inds add column ? varchar(256);",table)
		err := ind.DataBase.Migrator().AddColumn(&Ind{},v)
		if err != nil{
			glog.Error(err)
		}
		glog.Info(ind.DataBase.Migrator().HasColumn(&Ind{},v))
	}
	mps := make([]map[string]interface{},0)
	for i,row := range rows{
		mp := make(map[string]interface{})
		if i == 0 {
			continue
		}
		for j, v := range row{
			mp[rows[0][j]] = v
		}
		fmt.Println(mp)
		mps = append(mps, mp)
	}
	ind.DataBase.Model(&Ind{}).Create(mps)
	defer ind.DataBase.Migrator().RenameTable("inds",filename)
}
