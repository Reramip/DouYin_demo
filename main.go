/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 16:22:29
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 22:21:12
 */
package main

import (
	"github.com/Anzz-bot/DouYin_demo/bootstrap"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//load config
	bootstrap.InitializeConfig()

	//load log
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	//load DB
	global.App.DB = bootstrap.InitializeDB()
	global.App.Log.Info("DB init success!")
	//close DB
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	r := gin.Default()

	//test gin router
	//success return OK
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	//start gin server
	r.Run(":" + global.App.Config.App.Port)
}
