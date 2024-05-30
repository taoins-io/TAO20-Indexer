package gdb

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"tao/database/gdb/table"
	"tao/logger"
)

var balanceLock sync.Mutex

func (object *ChainDB) Tao20Mint(tao20Operation table.Tao20Operation) {
	//check tao20 valid & mintBlock >= tao20 Block
	tao20 := object.GetByTickerHex(tao20Operation.TickerHex)
	if tao20.ID < 1 || tao20Operation.Block < tao20.Block {
		return
	}
	if tao20Operation.Block == tao20.Block && tao20Operation.EventIndex <= tao20.EventIndex {
		return
	}
	//Does not exceed the single mint limit and the mint limit is sufficient
	if tao20Operation.Amount > tao20.MintMaxTimes || tao20Operation.Amount+tao20.Minted > tao20.Amount {
		return
	}
	isSaved := object.SaveOperation(tao20Operation)
	if !isSaved {
		return
	}
	err := object.db.Model(&table.Tao20{}).Where("id=?", tao20.ID).Update("minted", tao20.Minted+tao20Operation.Amount).Error
	if err != nil {
		logger.GetLogger().Errorf("Tao20Mint err %v", err)
		return
	}
	err = object.UpdateBalance(tao20Operation.Sender, tao20Operation.Ticker, tao20Operation.TickerHex, tao20Operation.Amount, 0, tao20Operation.Block, tao20Operation.BlockTime)
	if err != nil {
		logger.GetLogger().Errorf("Tao20Mint err %v", err)
	}
}

func (object *ChainDB) UpdateBalance(owner, ticker, tickerHex string, amount, sendingAmount, block, blockTime int64) error {
	balanceLock.Lock()
	defer balanceLock.Unlock()
	tao20Balance := object.GetByOwnerTickerHex(table.Tao20Balance{
		Owner:     owner,
		Ticker:    ticker,
		TickerHex: tickerHex,
	})
	if tao20Balance.ID < 1 {
		tao20Balance.Ticker = ticker
		tao20Balance.TickerHex = tickerHex
		tao20Balance.Owner = owner
	}
	if tao20Balance.Amount+amount < 0 {
		return errors.New("Insufficient balance")
	}
	if tao20Balance.Amount+amount-tao20Balance.SendingAmount-sendingAmount < 0 {
		return errors.New("Insufficient balance")
	}
	if tao20Balance.SendingAmount+sendingAmount < 0 {
		return errors.New("Insufficient sending balance")
	}
	tao20Balance.LastBlock = block
	tao20Balance.LastBlockTime = blockTime
	tao20Balance.Amount = tao20Balance.Amount + amount
	tao20Balance.SendingAmount = tao20Balance.SendingAmount + sendingAmount
	err := object.db.Save(&tao20Balance).Error
	if err != nil {
		logger.GetLogger().Errorf("UpdateBalance err %v", err)
	}
	return err
}

func (object *ChainDB) GetByOwnerTickerHex(balance table.Tao20Balance) (tao20Balance table.Tao20Balance) {
	err := object.db.Model(&table.Tao20Balance{}).Where("owner = ? and ticker_hex = ?", balance.Owner, balance.TickerHex).First(&tao20Balance).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.GetLogger().Errorf("GetByOwnerTickerHex err %v", err)
	}
	return
}
