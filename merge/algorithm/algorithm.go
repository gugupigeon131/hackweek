package algorithm

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
	"学生系统/Model"
	"学生系统/db"
	"学生系统/jsonFile"
)

//const
var percent float64 = 0.75
var part1Rate float64 = 0.95
var timeLength float64 = 24.0
var demoNum int = 12
var recordDataPath string = "./algorithm/recordData.json"
var timeFormatDay = "2006-01-02"
var timeFormatSec string = "2006-01-02 15:04:05"
var table = "building"

//
var building []Model.Building
var obj []KeyID
var preFloat []float64
var buildingNum int
var recordData RecordData
var now time.Time


type KeyID struct {
	Key float64 `json:"key"`
	ID string	`json:"id"`
}

type RecordData struct {
	TargetVal float64		`json:"targetVal"`
	PreSortTimeStr string	`json:"preSortTimeStr"`		//time.Parse("2006-01-02 15:04:05", PreSortTime)
	PreDayStr string		`json:"preDayStr"`			//time.Parse("2006-01-02", preDay)
}


func sort(obj []KeyID,x int,y int) {
	if buildingNum==0 {
		return
	}
	mid:=obj[(x+y)/2].Key
	var i,j=x,y
	for i<=j {
		for obj[i].Key>mid {
			i++
		}
		for obj[j].Key<mid {
			j--
		}
		if i<=j {
			obj[i],obj[j]=obj[j],obj[i]
			i++
			j--
		}
	}
	if x<j  {
		sort(obj,x,j)
	}
	if i<y  {
		sort(obj,i,y)
	}
	return
}

func readBuilding() {
	//读数据库表
	var err error
	building,err=db.SelectAll(table)
	if err!=nil {
		panic(err)
	}
	buildingNum=len(building)
	fmt.Println(building,buildingNum)
	return
}


func keyIDSort() (err error) {//用于读入、更新数据，并排序
	readBuilding()

	//计算排序关键字，存入obj并排序
	obj=[]KeyID{}
	var rankVal float64
	nowUnix:=time.Now().Unix()
	for i:=0;i<buildingNum;i++ {
		rankVal=building[i].Part1+recordData.TargetVal*timeLength/((float64(nowUnix-building[i].StartTime))/3600+timeLength)
		obj=append(obj,KeyID{
			Key:rankVal,
			ID:building[i].Id,
		})
	}
	sort(obj,0,buildingNum-1)

	//更新数据库表排名
	for i:=0;i<buildingNum;i++ {
		err=db.Update(table,db.KeyVal{Key:"id",Val: obj[i].ID},db.KeyVal{Key: "rank",Val: strconv.Itoa(i + 1)})
		if err!=nil {
			panic(err)
		}
	}

	//更新preFloat
	preFloat=[]float64{}
	preFloat=append(preFloat,math.Pow(0.5,6/float64(buildingNum)))
	for i:=1;i<buildingNum;i++ {
		preFloat=append(preFloat,preFloat[i-1]+math.Pow(0.5,6.0*float64(i)/float64(buildingNum)))
	}

	//更新recordData文件
	if buildingNum!=0 {
		recordData.TargetVal=obj[int(math.Floor(float64(buildingNum)*percent))].Key
	}
	recordData.PreSortTimeStr=time.Now().Format(timeFormatSec)
	err=jsonFile.Write(recordDataPath,recordData)
	if err!=nil {
		panic(err)
	}
	return nil
}

//快速幂
func pow(base float64,x int64) float64 {
	var ans float64 = 1
	for x>0 {
		if (x&1) ==1 {
			ans*=base
		}
		base=base*base
		x>>=1
	}
	return ans
}

//判断是否对part1处理
func checkAndDoPart1() error {
	PreDay,err:=time.Parse(timeFormatDay, recordData.PreDayStr)
	if err!=nil {
		panic(err)
	}
	nowDay,err:=time.Parse(timeFormatDay, now.Format(timeFormatDay))
	if err!=nil {
		panic(err)
	}
	interval:=(nowDay.Unix()-PreDay.Unix())/86400

	if interval>0 {
		factor:=pow(part1Rate,interval)

		/*
		factor:=0.1
		for interval>0 {//可用快速幂优化，虽然实际情况用了跟没用没什么区别
			interval--
			factor=factor*part1Rate
			interval>>=1
		}
		*/

		//保存修改到数据库
		for i:=0;i<buildingNum;i++ {
			err=db.Update(table,db.KeyVal{Key: "id",Val:obj[i].ID},db.KeyVal{Key: "part1",Val: strconv.FormatFloat(building[i].Part1*factor,'E',-1,64)})
			if err!=nil {
				panic(err)
			}
		}

		//更新recordData文件
		recordData.PreDayStr=time.Now().Format(timeFormatDay)
		err:=jsonFile.Write(recordDataPath,recordData)
		if err!=nil {
			panic(err)
		}

		//更新obj
		err=keyIDSort()
		if err!=nil {
			panic(err)
		}
	}

	return nil
}

//判断是否排序
func checkAndDoSort() error {
	PreSortTime,err:=time.Parse(timeFormatSec, recordData.PreSortTimeStr)
	if err!=nil {
		panic(err)
	}
	total:=float64(buildingNum)+2
	if now.Unix()-PreSortTime.Unix()>=int64(total*math.Log2(total)/100+1) {
		err:=keyIDSort()
		if err!=nil {
			panic(err)
		}
	}
	return nil
}

//对外开放的函数，判断各项操作是否进行
func CheckAndDo() (err error) {
	now=time.Now()
	err=checkAndDoPart1()
	if err!=nil {
		panic(err)
	}
	err=checkAndDoSort()
	if err!=nil {
		panic(err)
	}
	return nil
}

//推荐算法
func Recommend() ([]Model.Building,error) {
	selectedKeyIDs := make([]KeyID,0)
	selectedBuildings := make([]Model.Building,0)
	demoPart1 := demoNum/3
	demoPart2 := demoNum/3*2

	//fmt.Println(building,demoNum)
	//fmt.Println(demoPart1,demoPart2)
	if buildingNum<=demoNum {
		for i:=0;i<buildingNum;i++ {
			selectedBuilding,err:=db.Select(table,db.KeyVal{Key: "id",Val:obj[i].ID})
			//fmt.Println(selectedBuilding)
			if err!=nil {
				panic(err)
			}
			selectedBuildings=append(selectedBuildings,selectedBuilding)
		}
	} else {
		exist := make(map[int]bool)
		for i:=0;i<demoPart1;i++ {
			selectedBuilding,err:=db.Select(table,db.KeyVal{Key: "id",Val:obj[i].ID})
			if err!=nil {
				panic(err)
			}
			selectedBuildings=append(selectedBuildings,selectedBuilding)
			exist[i]=true
		}

		var randNum float64
		var l,r,m,tar int
		preFloatLen := len(preFloat)
		//fmt.Println(preFloatLen)
		for i:=0;i<demoPart2;i++ {
			randNum=rand.Float64()*preFloat[preFloatLen-1]
			l=0
			r=buildingNum-1
			for l<r {
				m=(l+r)>>1
				if randNum<preFloat[m] {
					r=m
				} else {
					l=m+1
				}
			}
			tar=l
			if _,ok:=exist[tar];ok {
				for j:=tar-1;j>=0;j-- {
					if _,ok:=exist[j];!ok {
						tar=j
						break
					}
				}
				if tar==l {
					for j:=tar+1;j<buildingNum;j++ {
						if _,ok:=exist[j];!ok {
							tar=j
							break
						}
					}
				}
			}
			selectedKeyIDs=append(selectedKeyIDs,obj[tar])
			exist[tar]=true
		}

		sort(selectedKeyIDs,0,len(selectedKeyIDs)-1)
		for i:=0;i<demoPart2;i++ {
			selectedBuilding,err:=db.Select(table,db.KeyVal{Key: "id",Val:selectedKeyIDs[i].ID})
			if err!=nil {
				panic(err)
			}
			selectedBuildings=append(selectedBuildings,selectedBuilding)
			exist[i]=true
		}
	}
	return selectedBuildings,nil
}

//初始化
func Init() error {
	now=time.Now()

	rand.Seed(now.UnixNano())

	err:=db.Init()
	if err!=nil {
		panic(err)
	}
	//读入recordData
	if !jsonFile.FileExist(recordDataPath) {
		err=jsonFile.Write(recordDataPath,RecordData{
			TargetVal: 0,
			PreSortTimeStr:now.Format(timeFormatSec),
			PreDayStr:now.Format(timeFormatDay),
		})
		if err!=nil {
			panic(err)
		}
	}
	err=jsonFile.Read(recordDataPath,&recordData)
	if err!=nil {
		panic(err)
	}

	if obj==nil {
		err=keyIDSort()
		if err!=nil {
			panic(err)
		}
	}
	//判断条件
	err=CheckAndDo()
	if err!=nil {
		panic(err)
	}

	return nil
}
func Test() {
	err:=keyIDSort()
	if err!=nil {
		panic(err)
	}
	fmt.Print(obj)
	fmt.Println(preFloat)
}