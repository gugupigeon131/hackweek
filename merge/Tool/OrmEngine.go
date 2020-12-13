package Tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"学生系统/Model"
)



type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm
func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.DatabaseConfig
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Dbname+"?charset=utf8"
	eng, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		panic(err)
		return nil, err
	}

	/////////////////////////////////////////////////////////////////////modify
	err = eng.Sync2(new(Model.StudentModel),new(Model.Building),new(Model.Article),new(Model.Like),new(Model.Save))
	/////////////////////////////////////////////////////////////////////modify

	if err != nil {
		panic(err)
	}
	orm := new(Orm)
	orm.Engine = eng

	DbEngine = orm

	return orm, nil 

}


