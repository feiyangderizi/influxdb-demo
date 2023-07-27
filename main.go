package main

import (
	"github.com/feiyangderizi/ginServer/controller"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"

	"github.com/feiyangderizi/ginServer/initialize"
)

const configFile = "application.yml"

//@title	ginServer example
//@version 	1.0.0(ginServer)
//@description	ginServer演示范例

func main() {
	//初始化配置，自动连接数据库
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	initialize.Init(path + "/" + configFile)

	//GIN的模式，生产环境可以设置成release
	gin.SetMode("debug")

	var ctrl controller.ImportController
	ctrl.Import()
}
