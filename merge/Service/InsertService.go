package Service

import (
	"学生系统/Dao"
	"学生系统/Model"
	"学生系统/Param"
	"学生系统/Tool"
)

type MemberService struct {
}

var ModelStudent Model.StudentModel

func (ms *MemberService)InsertArticle(massage Param.Article){
	article := Model.Article{}
	article.Id = massage.Id
	article.Name = massage.Name
	article.UpId = massage.UpId
	article.UpTime = massage.UpTime
	article.TextSrc = massage.TextSrc
	article.PicSrc = massage.PicSrc
	article.CommentSrc = massage.CommentSrc
	article.CommentNum = massage.CommentNum
	article.LikeNum = massage.LikeNum
	article.ViewNum = massage.ViewNum
	article.Part1 = massage.Part1
	article.Rank = massage.Rank
	dao := Dao.MemberDao{Tool.DbEngine}
	dao.InsertArticle(article)



}

func (ms *MemberService) InsertBuilding(massage Param.Building) {
	building := Model.Building{}
	//必要数据
	building.Id = massage.Id
	building.BuildingName = massage.BuildingName
	building.BuildingStatus = massage.BuildingStatus
	building.BuildingClass = massage.BuildingClass
	building.BuildingStyle = massage.BuildingStyle
	//后台数据
	building.TextSrc = massage.TextSrc
	building.PicSrc = massage.PicSrc
	building.StartTime = massage.StartTime
	building.ContentNum = massage.ContentNum
	building.CommentNum = massage.CommentNum
	building.LikeNum = massage.LikeNum
	building.ViewNum = massage.ViewNum
	building.Part1 = massage.Part1
	building.Rank =massage.Rank
	dao := Dao.MemberDao{Tool.DbEngine}
	dao.InsertBuilding(building)

}

func (ms *MemberService) InsertStudent(massage Param.StudentParam) {

	student := Model.StudentModel{}

	student.Password = massage.Password
	student.Name = massage.Name
	student.Id = massage.Id
	student.Phone = massage.Phone
	student.OtherThing = massage.OtherThing

	dmo := Dao.MemberDao{Tool.DbEngine}

	dmo.InsertStudent(student)

}
func (ms *MemberService) QueryArticleById(id string) *Model.Article{

	dmo := Dao.MemberDao{Tool.DbEngine}

	article:= dmo.QueryArticleById(id)

	return article

}

func (ms *MemberService) QueryArticleByName(Name string) *Model.Article{

	dmo := Dao.MemberDao{Tool.DbEngine}

	article:= dmo.QueryArticleById(Name)

	return article

}

func (ms *MemberService) QueryBuildingById(id string) *Model.Building{

	dmo := Dao.MemberDao{Tool.DbEngine}

	building:= dmo.QueryBuildingById(id)

	return building

}
func (ms *MemberService) QueryBuildingByName(Name string) *Model.Building{

	dmo := Dao.MemberDao{Tool.DbEngine}

	building:= dmo.QueryBuildingById(Name)

	return building

}



func (ms *MemberService) QueryStudentById(id string) *Model.StudentModel {

	dmo := Dao.MemberDao{Tool.DbEngine}

	student := dmo.QueryStudentById(id)

	return student

}



func (ms *MemberService) UpdataArticle(massage Param.Article) int64 {
	article := Model.Article{}
	article.Id = massage.Id
	article.Name = massage.Name
	article.UpId = massage.UpId
	article.UpTime = massage.UpTime
	article.TextSrc = massage.TextSrc
	article.PicSrc = massage.PicSrc
	article.CommentSrc = massage.CommentSrc
	article.CommentNum = massage.CommentNum
	article.LikeNum = massage.LikeNum
	article.ViewNum = massage.ViewNum
	article.Part1 = massage.Part1
	article.Rank = massage.Rank
	dmo := Dao.MemberDao{Tool.DbEngine}
	result := dmo.UpdataArticle(article)
	return result
}



func (ms *MemberService) UpdataBuilding(massage Param.Building) int64 {
	building := Model.Building{}

	//必要数据
	building.Id = massage.Id
	building.BuildingName = massage.BuildingName
	building.BuildingStatus = massage.BuildingStatus
	building.BuildingClass = massage.BuildingClass
	building.BuildingStyle = massage.BuildingStyle
	//后台数据
	building.TextSrc = massage.TextSrc
	building.PicSrc = massage.PicSrc
	building.StartTime = massage.StartTime
	building.ContentNum = massage.ContentNum
	building.CommentNum = massage.CommentNum
	building.LikeNum = massage.LikeNum
	building.ViewNum = massage.ViewNum
	building.Part1 = massage.Part1
	building.Rank =massage.Rank
	dmo := Dao.MemberDao{Tool.DbEngine}
	result := dmo.UpdataBuilding(building)
	return result
}


func (ms *MemberService) UpdataStudent(massage Param.StudentParam) int64 {
	student := Model.StudentModel{}

	student.Password = massage.Password
	student.Name = massage.Name
	student.Id = massage.Id
	student.Phone = massage.Phone
	student.OtherThing = massage.OtherThing
	dmo := Dao.MemberDao{Tool.DbEngine}
	result := dmo.UpdataStudent(student)
	return result
}
func (ms *MemberService) DeleteBuilding(id string) int64 {
	dmo := Dao.MemberDao{Tool.DbEngine}

	result := dmo.DeleteBuilding(id)

	return result

}

func (ms *MemberService) DeleteArticle(id string) int64 {
	dmo := Dao.MemberDao{Tool.DbEngine}

	result := dmo.DeleteArticle(id)

	return result

}

func (ms *MemberService) DeleteStudent(id string) int64 {
	dmo := Dao.MemberDao{Tool.DbEngine}

	result := dmo.DeleteStudent(id)

	return result

}

func (ms *MemberService) InitAdministrator() {
	dmo := Dao.MemberDao{Tool.DbEngine}

	dmo.InitAdministrator()

}

func (ms *MemberService)ShowAllArticle(BuildingName string) *[]Model.Article {
	dao := Dao.MemberDao{Tool.DbEngine}
	articles := dao.ShowAllArticle(BuildingName)
	return articles

}


