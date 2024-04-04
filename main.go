package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ListenPort = 8080

func main() {
	var bindAddress = fmt.Sprintf("0.0.0.0:%d", ListenPort)

	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	tmpl, _ := template.New(PageName).Parse(Page)
	engine.SetHTMLTemplate(tmpl)

	engine.GET("/", mainPageHandler)

	_ = engine.Run(bindAddress)
}

func mainPageHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, PageName, nil)
}
