// Package main
// @Time:2023/03/23 00:43
// @File:main.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

func getSubInfo(realUrl string) any {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", realUrl, nil)
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	r.Header.Add("upgrade-insecure-requests", "1")
	r.Header.Add("Accept", "*/*")
	r.Header.Add("Connection", "keep-alive")
	resp, err := client.Do(r)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/subinfo", func(c *gin.Context) {
		firstUrl := c.Query("url")
		realUrl, _ := url.QueryUnescape(firstUrl)
		if str, ok := getSubInfo(realUrl).(string); ok {
			c.String(http.StatusOK, str)
		} else {
			c.String(http.StatusOK, "订阅获取失败")
		}
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":35500")
}
