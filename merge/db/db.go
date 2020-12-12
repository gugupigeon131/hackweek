package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"学生系统/Model"
	"学生系统/jsonFile"
)

type UserData struct {
	Username	string		`db:"username"`
	Name		string		`db:"name"`
	PhoneNumber string		`db:"phoneNumber"`
}

type KeyVal struct {
	Key		string
	Val		string
}

type Config struct {
	EngineHost  string `json:"engine_host"`
	EnginePort  string  `json:"engine_port"`
	DatabaseConfig *DatabaseConfig `json:"database_config"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"db_name"`
	Charset  string `json:"charset"`
}

var config Config

func Init() (err error) {
	err=jsonFile.Read("./Config/config.json",&config)
	if  err!=nil {
		panic(err)
	}
	return
}

func LikeInsert(like Model.Like) error {
	db, err := sqlx.Open(config.DatabaseConfig.Driver, config.DatabaseConfig.User+":"+config.DatabaseConfig.Password+"@tcp("+config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port+")/"+config.DatabaseConfig.Dbname+"?charset="+config.DatabaseConfig.Charset)
	if err!=nil {
		return err
	}
	_, err = db.Exec("insert into like (user_id, at_id) values (?, ?)", like.UserId, like.AtId)
	if err!=nil {
		return err
	}
	return nil
}


func Select(table string,option KeyVal) (Model.Building,error) {
	v:=make([]Model.Building,0)
	db, err := sqlx.Open(config.DatabaseConfig.Driver, config.DatabaseConfig.User+":"+config.DatabaseConfig.Password+"@tcp("+config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port+")/"+config.DatabaseConfig.Dbname+"?charset="+config.DatabaseConfig.Charset)
	if err!=nil {
		panic(err)
	}
	err = db.Select(&v, "select * from "+table+" where "+option.Key+"=?", option.Val)
	if err!=nil {
		panic(err)
	}
	return v[0],nil
}

func LikeSelect(table string,query string) ([]Model.Like,error) {
	v:=make([]Model.Like,0)
	db, err := sqlx.Open(config.DatabaseConfig.Driver, config.DatabaseConfig.User+":"+config.DatabaseConfig.Password+"@tcp("+config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port+")/"+config.DatabaseConfig.Dbname+"?charset="+config.DatabaseConfig.Charset)
	if err!=nil {
		panic(err)
	}
	err = db.Select(&v, "select * from "+table+" where "+query)
	if err!=nil {
		panic(err)
	}
	return v,nil
}

func SelectAll(table string) ([]Model.Building,error) {
	v:=make([]Model.Building,0)
	db, err := sqlx.Open(config.DatabaseConfig.Driver, config.DatabaseConfig.User+":"+config.DatabaseConfig.Password+"@tcp("+config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port+")/"+config.DatabaseConfig.Dbname+"?charset="+config.DatabaseConfig.Charset)
	if err!=nil {
		panic(err)
	}
	err = db.Select(&v, "select * from "+table)
	if err!=nil {
		panic(err)
	}
	return v,nil
}

func Update(table string,option KeyVal,do KeyVal) error {
	db, err := sqlx.Open(config.DatabaseConfig.Driver, config.DatabaseConfig.User+":"+config.DatabaseConfig.Password+"@tcp("+config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port+")/"+config.DatabaseConfig.Dbname+"?charset="+config.DatabaseConfig.Charset)
	if err!=nil {
		panic(err)
	}
	query := "update `"+table+"` set `"+do.Key+"`=? where "+option.Key+"=?"
	_, err = db.Exec(query, do.Val, option.Val)
	if err!=nil {
		panic(err)
	}
	return nil
}

func Delete(table string,option KeyVal) (int64,error) {
	db, err := sqlx.Open(config.DatabaseConfig.Driver, config.DatabaseConfig.User+":"+config.DatabaseConfig.Password+"@tcp("+config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port+")/"+config.DatabaseConfig.Dbname+"?charset="+config.DatabaseConfig.Charset)
	if err!=nil {
		return 0,err
	}
	result, err := db.Exec("delete from "+table+" where "+option.Key+"=?", option.Val)
	if err!=nil {
		return 0,err
	}
	rows,err :=result.RowsAffected()
	if err!=nil {
		return 0,err
	}
	return rows,nil
}