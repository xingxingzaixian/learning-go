package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 获取项目目录
	workPath := filepath.Dir(filepath.Dir(dir))
	fmt.Println("dir: ", workPath)
	return workPath
}

func main() {
	workDir := getWorkingDirPath()

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/dist"
		router.HandleContext(ctx)
	})

	router.Static("/dist", filepath.Join(workDir, "static", "dist"))
	router.StaticFS("/img", http.Dir(filepath.Join(workDir, "static", "images")))
	router.Run(":8290")
}
