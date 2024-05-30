package gdb

import (
	"sync"
	"tao/database/gdb/table"
	"tao/logger"
)

var tao20Lock sync.Mutex

func (object *ChainDB) SaveTao20(tao20Operation table.Tao20Operation) error {
	tao20Lock.Lock()
	defer tao20Lock.Unlock()
	object.SaveOperation(tao20Operation)
	tao20 := table.Tao20{
		Sender:       tao20Operation.Sender,
		Receiver:     tao20Operation.To,
		ReceiverHex:  tao20Operation.ToHex,
		Ticker:       tao20Operation.Ticker,
		TickerHex:    tao20Operation.TickerHex,
		Amount:       tao20Operation.Amount,
		MintMaxTimes: tao20Operation.Additional,
		Minted:       0,
		EventIndex:   tao20Operation.EventIndex,
		Block:        tao20Operation.Block,
		BlockTime:    tao20Operation.BlockTime,
	}
	err := object.db.Model(&table.Tao20{}).Where("ticker_hex = ?", tao20.TickerHex).FirstOrCreate(&tao20).Error
	if err != nil {
		logger.GetLogger().Errorf("SaveTao20 err %v", err)
	}
	return err
}

func (object *ChainDB) GetByTickerHex(tickerHex string) (tao20 table.Tao20) {
	err := object.db.Model(&table.Tao20{}).Where("ticker_hex = ?", tickerHex).First(&tao20).Error
	if err != nil {
		logger.GetLogger().Errorf("GetByTickerHex err %v", err)
	}
	return
}
