package main

import (
	"github.com/gin-gonic/gin"
	"学生系统/Controller"
	"学生系统/Service"
	"学生系统/Tool"
)

func main() {
	cfg := Tool.ParseConfig("./Config/config.json")


	eng := gin.Default()

	//gorm := Tool.Gorm{}
	//gorm.GormEng(cfg)

	_,err :=Tool.OrmEngine(cfg)
	if err != nil {
		panic(err)
	}

	/////////////////////////////////////////////////////////add
	//之前已经测试过可用，包括与其相关的路由:/square
	//但是，后续的修改已经导致该模块不可用，具体错误未知，反正用不上
	/*
		err=algorithm.Init()
		if err != nil {
			panic(err)
		}
	*/
	/////////////////////////////////////////////////////////


	eng.Use(Controller.Cors())
	serv := Service.MemberService{}

	serv.InitAdministrator()

	RegisterRouter(eng)







	eng.Run(":"+cfg.EnginePort)

}

func RegisterRouter(eng *gin.Engine)  {
	new(Controller.MembersController).Router(eng)
}