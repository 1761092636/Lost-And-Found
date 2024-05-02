package eth

import (
	"Lost_and_Found/db"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// LostItem 定义失物信息结构体模型
type LostItem struct {
	LostItemID             *big.Int       `json:"lostItemId" form:"lostItemId"`                         // 失物ID
	User                   common.Address `json:"user,omitempty" form:"user,omitempty"`                 // 用户地址
	FoundOne               common.Address `json:"foundOne,omitempty" form:"foundOne,omitempty"`         // 找到者地址
	LostItemName           string         `json:"lostItemName" form:"lostItemName"`                     // 失物名称
	Description            string         `json:"description,omitempty" form:"description,omitempty"`   // 描述
	LostDate               *big.Int       `json:"lostDate,omitempty" form:"lostDate,omitempty"`         // 丢失日期
	Location               string         `json:"location,omitempty" form:"location,omitempty"`         // 丢失地点
	ItemCategory           string         `json:"itemCategory,omitempty" form:"itemCategory,omitempty"` // 物品种类
	IsFound                bool           `json:"isFound" form:"isFound"`                               // 是否找到
	FoundDate              *big.Int       `json:"foundDate,omitempty" form:"foundDate,omitempty"`       // 找到日期
	Reward                 *big.Int       `json:"reward,omitempty" form:"reward,omitempty"`             // 赏金
	FoundApplicationsCount *big.Int       `json:"applicationCount" form:"applicationCount"`             // 发现申请数量
	//FoundApplications     map[*big.Int]Application `json:"foundApplications"`                       // 发现申请映射
}

type Application struct {
	ApplicationId *big.Int       `sql:"application_id" json:"application_id" form:"application_id"`
	RelateItemID  *big.Int       `sql:"lost_item_id" json:"lost_item_id" form:"lost_item_id"`                // 与失物关联的ID
	Sender        common.Address `sql:"submitter_address" json:"submitter_address" form:"submitter_address"` // 提交者的以太坊账户地址
	Receiver      common.Address `json:"receiver" form:"receiver"`
	SubmitDate    *big.Int       `sql:"found_date" json:"found_date" form:"found_date"`    // 拾得日期和时间
	IsConfirmed   bool           `sql:"confirmed" json:"confirmed" form:"confirmed"`       // 是否已被确认
	RewardPaid    bool           `sql:"reward_paid" json:"reward_paid" form:"reward_paid"` // 是否已支付赏金
	Result        bool           `sql:"result" json:"result" form:"result" `
}
type History struct {
	HistoryID         *big.Int
	Operator          common.Address `json:"operator" form:"operator"`
	Operation         string
	OperationTime     *big.Int
	RelatedLostItemID *big.Int
}

// RankInfo 存储排名信息
type RankInfo struct {
	User   db.User  // 用户信息
	Points *big.Int // 用户积分
	Rank   int      // 用户排名
}
type Tran struct {
	Amount *big.Int `json:"amount" form:"amount"`
}

// TraceInfo 是溯源信息的结构体
type TraceInfo struct {
	LostItemID *big.Int         `json:"lostItemID"`
	FoundOnes  []common.Address `json:"foundOnes"`
	Total      *big.Int         `json:"total"`
	TraceHash  [32]byte         `json:"traceHash" form:"traceHash"`
}
type TraceHash struct {
	TraceHash string `json:"traceHash" form:"traceHash"`
}
