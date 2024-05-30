package table

type Tao20Operation struct {
	ID               int64  `gorm:"primarykey;column:id"`
	Updated          int64  `gorm:"column:updated;autoUpdateTime:milli"`
	Created          int64  `gorm:"column:created;autoCreateTime:milli"`
	Ticker           string `gorm:"column:ticker;"`
	TickerHex        string `gorm:"column:ticker_hex;"`
	Sender           string `gorm:"column:sender;"`
	To               string `gorm:"column:to;"`
	ToHex            string `gorm:"column:to_hex;"`
	Receiver         string `gorm:"column:receiver;"`
	Amount           int64  `gorm:"column:amount;"`
	Operation        string `gorm:"column:operation;"`
	Additional       int64  `gorm:"column:additional;"`
	EventIndex       string `gorm:"column:event_index;"`
	Block            int64  `gorm:"column:block_num;"`
	BlockTime        int64  `gorm:"column:block_time;"`
	SendEventIndex   string `gorm:"column:send_event_index;"`
	SendBlock        int64  `gorm:"column:send_block_num;"`
	SendBlockTime    int64  `gorm:"column:send_block_time;"`
	CancelEventIndex string `gorm:"column:cancel_event_index;"`
	CancelBlock      int64  `gorm:"column:cancel_block_num;"`
	CancelBlockTime  int64  `gorm:"column:cancel_block_time;"`
	Status           int    `gorm:"column:status;"`
}

func (Tao20Operation) TableName() string {
	return "tao20_operation"
}
