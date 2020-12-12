package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"学生系统/Dao"
	"学生系统/Model"
	"学生系统/Param"
	"学生系统/Service"
	"学生系统/Token"
	"学生系统/Tool"
)

type MembersController struct {
}

func (mc *MembersController) Router(GinEngine *gin.Engine) {
	GinEngine.POST("/register",mc.Register)
	GinEngine.POST("/login",mc.Login)
	GinEngine.POST("/newbuilding",mc.CheckToken,mc.NewBuilding)
	GinEngine.POST("/newarticle",mc.CheckToken,mc.NewArticle)
	GinEngine.POST("/updateuser",mc.CheckToken,mc.UpdataStudent)
	GinEngine.POST("/updatebuilding",mc.CheckToken,mc.UpdataBuilding)
	GinEngine.POST("/updatearticle",mc.CheckToken,mc.UpdataArticle)
	GinEngine.POST("/deleteuser",mc.CheckToken,mc.DeleteStudent)
	GinEngine.POST("/deletebuilding",mc.CheckToken,mc.DeleteBuilding)
	GinEngine.POST("/deletearticle",mc.CheckToken,mc.DeleteArticle)
	GinEngine.POST("/queryself",mc.CheckToken,mc.QueryStudent)
	GinEngine.POST("/querybuildingbyname",mc.CheckToken,mc.QueryBuildingByName)
	GinEngine.POST("/queryarticlebyname",mc.CheckToken,mc.QueryArticleByName)
	GinEngine.POST("/showallarticle",mc.CheckToken,mc.ShowAllArticle)


	/////////////////////////////////////////////////////////add
	//GinEngine.GET("/square",mc.CheckToken,mc.RecommendBuilding)//无用的代码罢了
	GinEngine.POST("/like",mc.CheckToken,mc.Like)
	GinEngine.POST("/save",mc.CheckToken,mc.Save)
	GinEngine.POST("/querysave",mc.CheckToken,mc.QuerySave)//save砍了，只返回发布过的,
	GinEngine.POST("/querybuildingbyclass",mc.CheckToken,mc.QueryBuildingByClass)
	GinEngine.POST("/querysavebuilding",mc.CheckToken,mc.QuerySaveBuilding)
	///////////////////////////////////////////////////////

}

func (mc *MembersController)NewArticle(context *gin.Context)  {
	var article Param.Article
	err := Tool.Decode(context.Request.Body,&article)
	if err != nil {
		log.Fatal(err.Error())
	}

	/////////////////////////////////////////add
	now:=time.Now()
	article.UpTime=now.Unix()
	article.Id= strconv.FormatInt(now.UnixNano(), 10)
	/////////////////////////////////////////add

	serv := Service.MemberService{}
	TheArticle := serv.QueryArticleById(article.Id)
	if TheArticle.Id == "" {
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "成功",
			"data":  article,
		})
		serv.InsertArticle(article)
	}else {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": TheArticle.Id+"已经存在" ,
			"data":    nil,
		})

	}


}





func (mc *MembersController) NewBuilding(context *gin.Context) {
	var building Param.Building
	err := Tool.Decode(context.Request.Body, &building)
	if err != nil {
		log.Fatal(err.Error())
	}

	/////////////////////////////////////////add
	now:=time.Now()
	building.StartTime=now.Unix()
	building.Id= strconv.FormatInt(now.UnixNano(), 10)
	/////////////////////////////////////////add

	serv := Service.MemberService{}
	TheBuilding := serv.QueryBuildingById(building.Id)
	if TheBuilding.Id == ""{
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "成功",
			"data":  building,
		})
		serv.InsertBuilding(building)
	}else {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": TheBuilding.BuildingName+"已经存在" ,
			"data":    nil,
		})

	}
}




func (mc *MembersController) Register (context *gin.Context) {
	var StudentMassage Param.StudentParam

	err := Tool.Decode(context.Request.Body, &StudentMassage)

	if err != nil {
		log.Fatal(err.Error())
	}
	serv := Service.MemberService{}
	TheStudent := serv.QueryStudentById(StudentMassage.Id)
	//fmt.Println(StudentMassage)
	if StudentMassage.Id =="" {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": "用户名不能为空，注册失败" ,
			"data":   StudentMassage,
		})
		context.Abort()
	}
	if TheStudent.Id == "" {
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "成功",
			"data":    StudentMassage,
		})
		serv.InsertStudent(StudentMassage)
		Token.GenerateToken(context, TheStudent)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": TheStudent.Name+"已经注册" ,
			"data":   StudentMassage,
		})

	}

}

func (mc *MembersController)QueryStudent(context *gin.Context) {
	var StudentMassage Param.StudentParam

	err := Tool.Decode(context.Request.Body, &StudentMassage)

	if err != nil {
		log.Fatal(err.Error())
	}

	serv := Service.MemberService{}

	student := serv.QueryStudentById(StudentMassage.Id)

	context.JSON(http.StatusOK, gin.H{
		"massage": "成功",
		"code":    200,
		"data":    student,
	})

}



func (mc *MembersController) QueryArticleByName(context *gin.Context) {
	var ArticleMassage Param.Article

	err := Tool.Decode(context.Request.Body, &ArticleMassage)

	if err != nil {
		log.Fatal(err.Error())
	}

	serv := Service.MemberService{}

	Article := serv.QueryArticleByName(ArticleMassage.Name)
	if Article.Name != " " {
		context.JSON(http.StatusOK, gin.H{
			"massage": "成功",
			"code":    200,
			"data":    Article,
		})
	}else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此建筑",
		})

	}


}



func (mc *MembersController) QueryBuildingByName(context *gin.Context) {
	var BuildingMassage Param.Building

	err := Tool.Decode(context.Request.Body, &BuildingMassage)

	if err != nil {
		log.Fatal(err.Error())
	}

	serv := Service.MemberService{}

	Building := serv.QueryBuildingByName(BuildingMassage.BuildingName)
	if Building.BuildingName != " " {
	context.JSON(http.StatusOK, gin.H{
		"massage": "成功",
		"code":    200,
		"data":    Building,
	})
	}else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此建筑",
		})

	}


}

func (mc *MembersController) UpdataStudent(context *gin.Context) {
	var student Param.StudentParam

	err := Tool.Decode(context.Request.Body, &student)

	if err != nil {
		log.Fatal(err.Error())
	}

	serv := Service.MemberService{}
	result := serv.UpdataStudent(student)
	if result != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"massage": "成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此人",
		})
	}

}

func (mc *MembersController) UpdataBuilding(context *gin.Context) {
	var Building Param.Building

	err := Tool.Decode(context.Request.Body, &Building)

	if err != nil {
		log.Fatal(err.Error())
	}

	serv := Service.MemberService{}
	result := serv.UpdataBuilding(Building)
	TheBuilding :=serv.QueryBuildingById(Building.Id)
	if result != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"massage": "成功",
			"data" : TheBuilding,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此建筑",
		})
	}

}

func (mc *MembersController) UpdataArticle(context *gin.Context) {
	var Article Param.Article

	err := Tool.Decode(context.Request.Body, &Article)

	if err != nil {
		log.Fatal(err.Error())
	}

	serv := Service.MemberService{}
	result := serv.UpdataArticle(Article)
	TheArticle :=serv.QueryArticleById(Article.Id)
	if result != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"massage": "成功",
			"data" : TheArticle,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此文章",
		})
	}

}


func (mc *MembersController) DeleteArticle(context *gin.Context) {
	var Article Param.Article
	err := Tool.Decode(context.Request.Body, &Article)
	if err != nil {
		log.Fatal(err.Error())
	}
	serv := Service.MemberService{}

	result := serv.DeleteArticle(Article.Id)

	if result != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"massage": "成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此人",
		})
	}

}


func (mc *MembersController) DeleteBuilding(context *gin.Context) {
	var Building Param.Building
	err := Tool.Decode(context.Request.Body, &Building)
	if err != nil {
		log.Fatal(err.Error())
	}
	serv := Service.MemberService{}

	result := serv.DeleteBuilding(Building.Id)

	if result != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"massage": "成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此人",
		})
	}

}



func (mc *MembersController) DeleteStudent(context *gin.Context) {
	var student Param.StudentParam
	err := Tool.Decode(context.Request.Body, &student)
	if err != nil {
		log.Fatal(err.Error())
	}
	serv := Service.MemberService{}

	result := serv.DeleteStudent(student.Id)

	if result != 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"massage": "成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":   400,
			"massage": "不存在此人",
		})
	}

}

func (mc *MembersController) SelfQuery(context *gin.Context) {
	id := context.Param("id")
	serv := Service.MemberService{}

	student := serv.QueryStudentById(id)

	context.JSON(http.StatusOK, gin.H{
		"massage": "成功",
		"code":    200,
		"data": student,
	})

}

func (mc *MembersController) Login(context *gin.Context) {
	var student Param.StudentParam
	err := Tool.Decode(context.Request.Body, &student)
	if err != nil {
		log.Fatal(err.Error())
	}
	serv := Service.MemberService{}
	TheStudent := serv.QueryStudentById(student.Id)

	if student.Password == TheStudent.Password {
		Token.GenerateToken(context, TheStudent)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"massage": "验证失败",
			"data":    nil,
		})
	}

}


func (mc *MembersController) CheckToken(context *gin.Context) {
	token := context.Request.Header.Get("token")
	if token == "" {
		context.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"massage": "token nil",
		})
		context.Abort()
		return

	}
	j := Token.NewJwt()
	claims, err := j.ParseToken(token)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    -1,
		})
		context.Abort()
	} else {
		context.Set("claims", claims)
		context.Next()
	}

}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range context.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}

func (mc *MembersController) ShowAllArticle(context *gin.Context) {
	var building Param.Building
	Tool.Decode(context.Request.Body, &building)
	serv := Service.MemberService{}
	articles := serv.ShowAllArticle(building.BuildingName)
	if articles !=nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"massage": "成功",
			"data": articles,
		})

	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"massage": "空空无物",
		})
	}
}



///////////////////////////////////////////////////////////////////////////add
/*
func (mc *MembersController) RecommendBuilding(context *gin.Context) {
	buildings,err:=algorithm.Recommend()
	if err!=nil {
		log.Fatal()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"buildings": buildings,
		"massage": "请求成功",
	})
}
*/

//记得建表
func (mc *MembersController) Like (context *gin.Context) {
	var like Model.Like
	err := Tool.Decode(context.Request.Body,&like)
	if err != nil {
		log.Fatal(err.Error())
	}
	md := Dao.MemberDao{Tool.DbEngine}
	choose:=Model.Like{}
	_,err=md.Where("user_id=?",like.UserId).And("at_id=?",like.AtId).Get(&choose)
	if err != nil {
		log.Fatal(err.Error())
	}
	if choose.AtId=="" {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"massage": "点赞成功",
		})
		_, err = md.InsertOne(&like)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"massage": "已经点过赞",
		})
	}
}

//记得建表
func (mc *MembersController) Save (context *gin.Context) {
	var save Model.Save
	err := Tool.Decode(context.Request.Body,&save)
	if err != nil {
		log.Fatal(err.Error())
	}
	md := Dao.MemberDao{Tool.DbEngine}
	choose:=Model.Save{}
	_,err=md.Where("user_id=?",save.UserId).And("at_id=?",save.AtId).Get(&choose)
	if err != nil {
		log.Fatal(err.Error())
	}
	if choose.AtId=="" {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"massage": "收藏成功",
		})
		_, err = md.InsertOne(&save)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"massage": "已经收藏过",
		})
	}
}
//记得建表

func (mc *MembersController) QuerySave (context *gin.Context) {
	var StudentMassage Param.StudentParam  // username

	err := Tool.Decode(context.Request.Body, &StudentMassage)

	if err != nil {
		log.Fatal(err.Error())
	}

	md := Dao.MemberDao{Tool.DbEngine}
	choose:=make([]Model.Article,0)
	err=md.Where("up_id=?",StudentMassage.Id).Find(&choose)
	if err != nil {
		log.Fatal(err.Error())
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"massage": "请求成功",
		"data":choose,
	})
}

func (mc *MembersController) QueryBuildingByClass (context *gin.Context) {
	var BuildingMassage Param.Building

	err := Tool.Decode(context.Request.Body, &BuildingMassage)

	if err != nil {
		log.Fatal(err.Error())
	}


	md := Dao.MemberDao{Tool.DbEngine}
	choose:=make([]Model.Building,0)
	err=md.Where("building_class=?",BuildingMassage.BuildingClass).Find(&choose)

	context.JSON(http.StatusOK, gin.H{
		"massage": "请求成功",
		"code":    200,
		"data":    choose,
	})
}

func (mc *MembersController) QuerySaveBuilding (context *gin.Context) {
	var StudentMassage Param.StudentParam  // username

	err := Tool.Decode(context.Request.Body, &StudentMassage)

	if err != nil {
		log.Fatal(err.Error())
	}

	md := Dao.MemberDao{Tool.DbEngine}
	choose:=make([]Model.Article,0)
	err=md.Where("up_id=?",StudentMassage.Id).Find(&choose)
	if err != nil {
		log.Fatal(err.Error())
	}

	chooseBuilding:=make([]Model.Building,0)
	chooseLen:=len(choose)
	serv := Service.MemberService{}

	for i:=0;i<chooseLen;i++ {///////////////////未去重
		building := *serv.QueryBuildingByName(choose[i].Name)
		chooseBuildingLen:=len(chooseBuilding)
		flag:=1
		for j:=0;j<chooseBuildingLen;j++ {
			if building.Id==chooseBuilding[j].Id {
				flag=0
				break
			}
		}
		if flag==1 {
			chooseBuilding=append(chooseBuilding,building)
		}
		//fmt.Println(choose[i].Name,building)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"massage": "请求成功",
		"data": chooseBuilding,
	})
}

