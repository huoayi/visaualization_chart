package logic

import (
	"flag"
	"visualization_chart/cmd/web"
)

func Run() {
	flag.Parse()
	Init("./config.yaml")
	web.Start()
}
