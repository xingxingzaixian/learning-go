package main

import (
	"fmt"
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
	distHandler := http.FileServer(http.Dir(filepath.Join(workDir, "static", "dist")))
	http.Handle("/", distHandler)

	// 文件服务器必须使用http.StripPrefix设置路由前缀，如果路由为/，则不用设置；
	// 目录访问路由必须以/结尾，否则无法访问目录
	// 路由地址与路由前缀设置必须一致，这样才能通过http://xx.xx.xx.xx:8280/img/去访问目录
	imgHandler := http.FileServer(http.Dir(filepath.Join(workDir, "static", "images")))
	http.Handle("/img/", http.StripPrefix("/img/", imgHandler))

	http.ListenAndServe(":8280", nil)
}
