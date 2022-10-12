package main

import "github.com/gin-gonic/gin"

func RunServer() {
	r := gin.Default()

	// 搜索包含key的文件
	r.GET("/search/:key", func(c *gin.Context) {
		c.JSON(200, codeIndex.search(c.Param("key")))
	})

	// 返回文档总数和token
	r.GET("/", func(c *gin.Context) {
		var maps = map[string]interface{}{}
		maps["len"] = len(files)
		maps["token"] = func() []string {
			var key = []string{}
			for k, _ := range codeIndex {
				key = append(key, k)
			}
			return key
		}()
		c.JSON(200, maps)
	})

	// 返回文档总数和token
	r.GET("/files/all", func(c *gin.Context) {
		var files_ = []File{}
		for _, v := range files {
			files_ = append(files_, File{v.ID, v.Path, ""})
		}
		c.JSON(200, files_)
	})
	r.Run()
}
