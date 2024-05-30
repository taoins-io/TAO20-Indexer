package table

type Tao20 struct {
	ID           int64  `gorm:"primarykey;column:id"`
	Updated      int64  `gorm:"column:updated;autoUpdateTime:milli"`
	Created      int64  `gorm:"column:created;autoCreateTime:milli"`
	Ticker       string `gorm:"column:ticker;"`
	TickerHex    string `gorm:"column:ticker_hex;"`
	Sender       string `gorm:"column:sender;"`
	Receiver     string `gorm:"column:receiver;"`
	ReceiverHex  string `gorm:"column:receiver_hex;"`
	Amount       int64  `gorm:"column:amount;"`
	MintMaxTimes int64  `gorm:"column:mint_max_times;"`
	Minted       int64  `gorm:"column:minted;"`
	EventIndex   string `gorm:"column:event_index;"`
	Block        int64  `gorm:"column:block_num;"`
	BlockTime    int64  `gorm:"column:block_time;"`
}

func (Tao20) TableName() string {
	return "tao20"
}
