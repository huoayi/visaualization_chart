package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"visualization_chart/internal/router/api"
)



const (
	V0 uint8 = 0
	V1 uint8 = 1
	V2 uint8 = 2
	V3 uint8 = 3
)

var (
	_hooks_V0, _hooks_V1, _hooks_V2, _hooks_V3 []Hook
)

type Hook func(router gin.IRoutes)

func Register(hook Hook, hookType uint8) {
	switch hookType {
	case V0:
		_hooks_V0 = append(_hooks_V0, hook)
		break
	case V1:
		_hooks_V1 = append(_hooks_V1, hook)
		break
	case V2:
		_hooks_V2 = append(_hooks_V2, hook)
		break
	case V3:
		_hooks_V3 = append(_hooks_V3, hook)
		break
	default:
		glog.Error("Register Error")
	}
}
//func Run() {
//	r := gin.New()
//	r.Use(gin.Logger())
//	r.Use(gin.Recovery())
//	v0 := r.Group("/api")
//	RegisterRouter(_hooks_V0,v0)
//	r.Run(fmt.Sprintf("%s:%d",global.Config.App.Host,global.Config.App.Port))
//	fmt.Println(r.Handlers)
//}
func RegisterRouter(hooks []Hook, v *gin.RouterGroup) {
	for _, hook := range hooks {
		hook(v)
	}
}

func Run(){
	r :=gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/upload",api.Upload)
	r.POST("/confirm",api.Confirm)
	r.GET("/get",api.GetList)
	r.GET("/getdata",api.GetData)
	r.Run(":8989")
}