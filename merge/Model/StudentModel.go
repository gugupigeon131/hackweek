package Model

type StudentModel struct {
	DatabaseId int64  `xorm:"pk autoincr" json:"database_id"`
	Phone      string `xorm:"varchar(50)" json:"phone"`
	Name       string `xorm:"varchar(50)" json:"name"`
	Password   string `xorm:"varchar(50)" json:"password"`
	Id         string `xorm:"varchar(50)" json:"id"`
	OtherThing string `xorm:"varchar(100)" json:"other_thing"`
}

type DateReq struct {
	ID    string
	Token string
}

type Building struct {
	Id             string `xorm:"pk unique varchar(50) " json:"id"`
	BuildingName   string `xorm:"varchar(20)" json:"building_name"`
	BuildingStatus int64  `xorm:"BigInt" json:"building_status"`
	BuildingClass  string  `xorm:"varchar(20)" json:"building_class"`
	BuildingStyle  int64  `xorm:"bigint" json:"building_style"`
	TextSrc        string `xorm:"varchar(100)" json:"text_src"`
	PicSrc         string `xorm:"varchar(100)" json:"pic_src"`
	StartTime      int64  `xorm:"bigint" json:"start_time"`

	ContentNum int     `xorm:"int" json:"content_num"`
	CommentNum int     `xorm:"int" json:"comment_num"`
	LikeNum    int     `xorm:"int" json:"like_num"`
	ViewNum    int     `xorm:"int" json:"view_num"`
	Part1      float64 `xorm:"double" json:"part_1"`
	Rank       int     `xorm:"int" json:"rank"`
}

type Article struct {
	Id         string  `xorm:"pk varchar(50) " json:"id"`
	Name       string  `xorm:"varchar(50)" json:"name"`   // 所属建筑名字
	UpId       string  `xorm:"varchar(16)" json:"up_id"`
	UpTime     int64   `xorm:"bigint " json:"up_time"`
	TextSrc    string  `xorm:"varchar(100)" json:"text_src"`
	PicSrc     string  `xorm:"varchar(100)" json:"pic_src"`
	CommentSrc string  `xorm:"varchar(100)" json:"comment_src"`
	CommentNum int     `xorm:"int" json:"comment_num"`
	LikeNum    int     `xorm:"int" json:"like_num"`
	ViewNum    int     `xorm:"int" json:"view_num"`
	Part1      float64 `xorm:"double" json:"part_1"`
	Rank       int     `xorm:"int" json:"rank"`
}



////////////////////////////////////////////////////////////////////add
type Like struct {
	UserId     	string `xorm:"pk unique varchar(100)" json:"user_id"`
	AtId		string `xorm:"varchar(100)" json:"at_id"`
}

type Save struct {
	UserId     	string `xorm:"pk unique varchar(100)" json:"user_id"`
	AtId		string `xorm:"varchar(100)" json:"at_id"`
}