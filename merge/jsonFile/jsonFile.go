package jsonFile

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Write(path string,v interface{}) error {//将v序列化后写入path文件,调用:jsonFile.Write(path,v)
	outData,err:=json.Marshal(v)
	if err!=nil {
		return err
	}
	err = ioutil.WriteFile(path, outData, 0777)
	if err!=nil {
		return err
	}
	return nil
}

func Read(path string,v interface{}) error {//将path文件反序列化后写入v，调用:jsonFile.Read(path,&v)
	buf, err := ioutil.ReadFile(path)
	if err!=nil {
		return err
	}
	err=json.Unmarshal(buf,&v)//读入基础信息
	if err!=nil {
		return err
	}
	return nil
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}