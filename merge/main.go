package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
		log.Fatal(err.Error())
	}

	/////////////////////////////////////////////////////////add
	/*
		err=algorithm.Init()
		if err != nil {
			log.Fatal(err.Error())
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