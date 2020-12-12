package Param

type StudentParam struct {
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Id         string `json:"username"`
	Token      string `json:"token"`
	OtherThing string `json:"other_thing"`
}
type Building struct {
	Id             string  `json:"id"`
	BuildingName   string  `json:"building_name"`
	BuildingStatus int64   `json:"building_status"`
	BuildingClass  string   `json:"building_class"`
	BuildingStyle  int64   `json:"building_style"`
	TextSrc        string  `json:"text_src"`
	PicSrc         string  `json:"pic_src"`
	StartTime      int64   `json:"start_time"`
	ContentNum     int     `json:"content_num"`
	CommentNum     int     `json:"comment_num"`
	LikeNum        int     `json:"like_num"`
	ViewNum        int     `json:"view_num"`
	Part1          float64 `json:"part_1"`
	Rank           int     `json:"rank"`
}
type Others struct {
	Buildings []string `json:"buildings"`
	Articles  []string `json:"articles"`
}
type Article struct {
	Id         string  ` json:"id"`
	Name       string  `json:"name"`
	UpId       string  ` json:"up_id"`
	UpTime     int64   ` json:"up_time"`
	TextSrc    string  `json:"text_src"`
	PicSrc     string  ` json:"pic_src"`
	CommentSrc string  ` json:"comment_src"`
	CommentNum int     ` json:"comment_num"`
	LikeNum    int     ` json:"like_num"`
	ViewNum    int     ` json:"view_num"`
	Part1      float64 ` json:"part_1"`
	Rank       int     ` json:"rank"`
}
