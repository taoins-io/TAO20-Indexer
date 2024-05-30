package table

type Tao20Balance struct {
	ID            int64  `gorm:"primarykey;column:id"`
	Updated       int64  `gorm:"column:updated;autoUpdateTime:milli"`
	Created       int64  `gorm:"column:created;autoCreateTime:milli"`
	Ticker        string `gorm:"column:ticker;"`
	TickerHex     string `gorm:"column:ticker_hex;"`
	Owner         string `gorm:"column:owner;"`
	Amount        int64  `gorm:"column:amount;"`
	SendingAmount int64  `gorm:"column:sending_amount;"`
	LastBlock     int64  `gorm:"column:last_block;"`
	LastBlockTime int64  `gorm:"column:Last_block_time;"`
}

func (Tao20Balance) TableName() string {
	return "tao20_balance"
}
