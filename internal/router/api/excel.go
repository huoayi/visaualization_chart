package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/xuri/excelize/v2"
	"net/http"
	"os"
	"path"
	"strings"
	"visualization_chart/global"
	e "visualization_chart/global/err_handler"
	"visualization_chart/internal/base"
	"visualization_chart/pkg"
)

//func init() {
//	router.Register(func(router gin.IRoutes) {
//		fmt.Println("1231231231312")
//		router.POST("/upload",upload)
//	},router.V0)
//}

func Upload(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400,err)
		fmt.Println(err)
		return
	}
	dst := path.Join("./upload", header.Filename)
	err = ctx.SaveUploadedFile(header, dst)
	defer os.Remove(dst)
	if err != nil{
		glog.Error(err)
		ctx.JSON(410,"错误")
	}
	xlsx ,err := excelize.OpenFile(dst)

	if err != nil{
		glog.Error(err)
		ctx.JSON(411,"")
	}
	if xlsx == nil{
	    glog.Info("文件读取失败")
	    ctx.JSON(412,"文件读取失败")
	}
	if ok,err := checkFile(xlsx,header.Filename); !ok{
		glog.Error(err)
		ctx.JSON(413,"未通过校验")
		return
	}

	if ans, err := getData(xlsx);err != nil{
		glog.Error(err)
		ctx.JSON(414,"获取失败")
	}else {
		ctx.JSON(200,ans)
	}

	ctx.JSON(200,"添加成功")
}
//

func getData(xlsx *excelize.File) ([][]string , error) {
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil{
		glog.Error(err)
		return nil,err
	}
	return rows, nil
}

//检查是否符合文件格式要求
func checkFile(xlsx *excelize.File,filename string)(bool,error) {
	ind := global.DB
	var temp base.FileName
	if err := ind.DataBase.Table(ind.Tables).Where("file_name = ?",filename).First(&temp).Error;err == nil {
		return false, e.GinError{
			Status:  http.StatusOK,
			Code:    1001,
			Message: "该文件[" + filename + "]已存在, 无需重复录入",
		}
	}
	return true,nil
}
func GetList(ctx *gin.Context){
	ind := global.DB
	var dal []base.FileName
	ind.DataBase.Table(ind.Tables).Find(&dal)
	var u []base.GetFileName
	for _, row := range dal{
		var v = base.GetFileName{
			File_name: row.File_name,
			Id : row.ID,
		}
		u = append(u, v)
	}
	 ctx.JSON(200,u)
}


func GetData(ctx *gin.Context)  {
	name := ctx.Query("filename")
	ind := global.DB
	mps := []map[string]interface{}{}
	ind.DataBase.Table(name).Find(&mps)
	ctx.JSON(200,mps)
}

//确认后将数据保存进数据库
func Confirm(ctx *gin.Context) {
	body := base.Body{}
	err := ctx.BindJSON(&body)
	if err != nil {
		glog.Error(err)
		ctx.JSON(411,"错误信息")
		return
	}
	if body.Data == nil{
		ctx.JSON(413,"传入数据为空")
	}
	setInDb(body.Data,body.Name)
}

func setInDb(rows [][]string,fileName string)  {

	ind := global.DB
	file := base.FileName{}
	file.File_name = fileName
	le := len(rows)
	ind1 := make([]string,le)

	for k, v := range rows[0]{
		ind1[k] = v
	}
	file.Column1 = ind1[0]
	file.Column2 = ind1[1]
	file.Column3 = ind1[2]
	file.Column4 = ind1[3]
	file.Column5 = ind1[4]
	file.Column6 = ind1[5]
	file.Column7 = ind1[6]
	file.Column8 = ind1[7]
	file.Column9 = ind1[8]
	file.Column10 = ind1[9]
	file.Column11 = ind1[10]
	file.Column12 = ind1[11]
	file.Column13 = ind1[12]
	file.Column14 = ind1[13]
	file.Column15 = ind1[14]
	file.Column16 = ind1[15]
	file.Column17 = ind1[16]
	file.Column18 = ind1[17]
	file.Column19 = ind1[18]
	file.Column20 = ind1[19]
	ind.DataBase.Table(ind.Tables).Create(&file)
	filename := strings.Split(file.File_name,".")[0]
	pkg.CreatTable(rows,filename)
}

