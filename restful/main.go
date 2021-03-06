/*
 * @Author: kslamp
 * @Date: 2019-12-11 20:11:22
 * @LastEditTime: 2019-12-12 13:38:26
 * @FilePath: /goshop/restful/main.go
 * @Description:
 */
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"goshop/restful/common"

	"goshop/restful/config"
	//"goshop/restful/models"
	"goshop/restful/routers"

	"goshop/libs/cache"
)

func main() {
	common.Init()
	cache.Init(config.MySqlConf.UserName, config.MySqlConf.PassWd, config.MySqlConf.Host, config.MySqlConf.Port, config.MySqlConf.DbName)
	//db := common.Init()

	//db.AutoMigrate(&models.Address{}, &models.Item{}, &models.ItemDesc{}, &models.Order{}, &models.Role{}, &models.User{})

	gin.SetMode(gin.DebugMode)
	s := &http.Server{
		Addr:           ":" + config.ServerConf.Port,
		Handler:        routers.Routers(),
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("start server  on ", config.ServerConf.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
