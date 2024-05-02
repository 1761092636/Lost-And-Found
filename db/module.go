package db

import (
	"math/big"
)

// 定义数据库的访问模型

// LostItem 定义失物信息结构体模型
type LostItem struct {
	LostItemID             int    `json:"lostItemId" form:"lostItemId" sql:"lostItemID"`                           // 失物ID
	User                   string `json:"myuser,omitempty" form:"myuser,omitempty" sql:"User"`                     // 用户名
	FoundOne               string `json:"foundOne,omitempty" form:"foundOne,omitempty" sql:"FoundOne"`             // 找到者
	LostItemName           string `json:"lostItemName" form:"lostItemName" sql:"LostItemName"`                     // 失物名称
	Description            string `json:"description,omitempty" form:"description,omitempty" sql:"Description"`    // 描述
	LostDate               int    `json:"lostDate,omitempty" form:"lostDate,omitempty"  sql:"LostDate"`            // 丢失日期
	Location               string `json:"location,omitempty" form:"location,omitempty" sql:"Location"`             // 丢失地点
	ItemCategory           string `json:"itemCategory,omitempty" form:"itemCategory,omitempty" sql:"itemCategory"` // 物品种类
	IsFound                bool   `json:"isFound" form:"isFound" sql:"IsFound"`                                    // 是否找到
	FoundDate              int    `json:"foundDate,omitempty" form:"foundDate,omitempty"  sql:"FoundDate"`         // 找到日期
	Reward                 int    `json:"reward,omitempty" form:"reward,omitempty" sql:"Reward"`                   // 赏金
	FoundApplicationsCount int    `json:"applicationCount" form:"applicationCount" sql:"FoundApplicationsCount"`   // 发现申请数量
	LostImg                string `json:"lostimg" form:"lostimg" sql:"LostImg"`                                    // 图片路径
}

// History 表示历史记录表
type History struct {
	HistoryID         int    `json:"history_id"`
	Operator          string `json:"user_id,omitempty"`
	Operation         string `json:"operation"`
	OperationTime     string `json:"operation_time,omitempty"`
	RelatedLostItemID int    `json:"related_lost_item_id,omitempty"`
}

// User 用户
type User struct {
	UserId       int      `sql:"user_id" json:"user_id" form:"user_id"`                   //用户id唯一标识
	UserName     string   `sql:"user_name" json:"user_name" form:"user_name"`             // 用户用户名
	UserPassword string   `sql:"user_password" json:"user_password" form:"user_password"` // 用户密码
	Address      string   `sql:"address" json:"address" form:"address"`                   // 以太坊账户地址
	UtcFile      string   `sql:"utcfile" json:"utcfile" form:"utcfile"`
	Name         string   `sql:"name" json:"name" form:"name"`
	Phone        string   `sql:"phone" json:"phone" form:"phone"`
	Balance      *big.Int `json:"balance" form:"balance"`
}
type Application struct {
	ApplicationId int    `sql:"id" json:"application_id" form:"application_id"`
	RelateItemID  int    `sql:"lost_item_id" json:"lost_item_id" form:"lost_item_id"`                // 与失物关联的ID
	Sender        string `sql:"submitter_address" json:"submitter_address" form:"submitter_address"` // 提交者的以太坊账户地址
	Receiver      string `sql:"receiver" json:"receiver"  form:"receiver"`
	SubmitDate    int    `sql:"found_date" json:"found_date" form:"found_date"`    // 拾得日期和时间
	IsConfirmed   bool   `sql:"confirmed" json:"confirmed" form:"confirmed"`       // 是否已被确认
	RewardPaid    bool   `sql:"reward_paid" json:"reward_paid" form:"reward_paid"` // 是否已支付赏金
	Result        bool   `sql:"result" json:"result" form:"result"`
}
type Message struct {
	From string `sql:"from" json:"from" form:"from"`
	To   string `sql:"to" json:"to" form:"to"`
	Info string `sql:"info" json:"info" form:"info"`
}
