package Dao

import (
	"log"
	"学生系统/Model"
	"学生系统/Tool"
)

type MemberDao struct {
	*Tool.Orm
}

func (md *MemberDao) InsertArticle(member Model.Article) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		log.Fatal(err.Error())
	}

	return result

}

func (md *MemberDao) InsertBuilding(member Model.Building) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		log.Fatal(err.Error())
	}

	return result

}

func (md *MemberDao) InsertStudent(member Model.StudentModel) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		log.Fatal(err.Error())
	}

	return result

}



func (md *MemberDao) QueryArticleByName(name string) *Model.Article{
	var article Model.Article
	_, err := md.Where("name = ? ", name).Get(&article)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &article

}

func (md *MemberDao) QueryBuildingByName(name string) *Model.Building {
	var building Model.Building
	_, err := md.Where(" building_name = ? ", name).Get(&building)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &building

}




func (md *MemberDao) QueryArticleById(id string) *Model.Article{
	var article Model.Article
	_, err := md.Where("id = ? ", id).Get(&article)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &article

}

func (md *MemberDao) QueryBuildingById(id string) *Model.Building {
	var building Model.Building
	_, err := md.Where("building_name = ? ", id).Get(&building)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &building

}

func (md *MemberDao) QueryStudentById(id string) *Model.StudentModel {
	var student Model.StudentModel
	_, err := md.Where("id = ? ", id).Get(&student)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &student

}
func (md *MemberDao) UpdataBuilding(member Model.Building) int64 {
	var Building Model.Building
	_, err := md.Where("id = ? ", member.Id).Get(&Building)

	result, err := md.Id(Building.Id).Update(member)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result

}


func (md *MemberDao) UpdataArticle(member Model.Article) int64 {
	var Article Model.Article
	_, err := md.Where("id = ? ", member.Id).Get(&Article)

	result, err := md.Id(Article.Id).Update(member)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result

}



func (md *MemberDao) UpdataStudent(member Model.StudentModel) int64 {
	var student Model.StudentModel
	_, err := md.Where("id = ? ", member.Id).Get(&student)

	result, err := md.Id(student.DatabaseId).Update(member)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result

}


func (md *MemberDao) DeleteBuilding(id string) int64 {
	var user Model.Building
	result, err := md.Where("id = ?", id).Delete(&user)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return result

}




func (md *MemberDao) DeleteArticle(id string) int64 {
	var user Model.Article
	result, err := md.Where("id = ?", id).Delete(&user)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return result

}


func (md *MemberDao) DeleteStudent(id string) int64 {
	var user Model.StudentModel
	result, err := md.Where("id = ?", id).Delete(&user)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return result

}

func (md *MemberDao) InitAdministrator() {
	var menber Model.StudentModel
	var truemember Model.StudentModel
	menber.Id = "12345678"
	menber.Password = "12345678"
	menber.Name = "admini"
	menber.Phone = "233333"
	md.Where("id=?", menber.Id).Get(&truemember)
	if truemember.Id == "" {
		_, err := md.InsertOne(&menber)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func (md *MemberDao)Tata() {

	ms := Model.StudentModel{}

	md.Where("password = ?", "12345678").And("id = ?", "123456789").Get(&ms)

	println(ms.Phone)

}

func (md *MemberDao)ShowAllArticle(BuildingName string) *[]Model.Article{
	var articles  []Model.Article
	md.Where("name = ?",BuildingName).Asc("id").Find(&articles)
	return &articles
	
}

