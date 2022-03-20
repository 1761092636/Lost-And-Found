package db

// 定义数据库的访问模型

// News 定义新闻结构体模型
type News struct {
	// 结构体成员tag
	NewsId  int    `sql:"news_id" json:"news_id" form:"news_id"`     // 唯一新闻的标识
	Content string `sql:"content" json:"content"  form:"content"`    // 新闻的内容
	Title   string `sql:"title" json:"title" form:"title" `          // 新闻标题
	NewsImg string `sql:"news_img" json:"news_img" form:"news_img" ` // 新闻图片地址
	Author  string `sql:"author" json:"author" form:"author" `       //作者
	Timeset string `sql:"timeset" json:"timeset" form:"timeset" `    //时间
	Like    int    `sql:"like" json:"like" form:"like" `             //点赞数
}

// Med 定义药品结构体模型
type Med struct {
	// 结构体成员tag
	MedId    int    `sql:"med_id" json:"med_id" form:"med_id"`          // 药品的唯一标识
	MedName  string `sql:"med_name" json:"med_name" form:"med_name"`    // 药品名
	MedTxt   string `sql:"med_txt" json:"med_txt" form:"med_txt"`       // 药品介绍
	MedPrice int    `sql:"med_price" json:"med_price" form:"med_price"` // 药品的价格
	MedImg   string `sql:"med_img" json:"med_img" form:"med_img"`       //药品照片
}

// Cart 定义购物车结构体模型
type Cart struct {
	// 结构体成员tag
	CartId int `sql:"cart_id" json:"cart_id" form:"cart_id"` // 评论的唯一标识

}
