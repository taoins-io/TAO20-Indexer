package gdb

import (
	"gorm.io/gorm"
	"strconv"
	"tao/database/gdb/table"
	"tao/logger"
	"tao/util"
	"tao/vo"
)

func (object *ChainDB) SaveOperation(operation table.Tao20Operation) (isSave bool) {
	status := 0
	if operation.Operation == "transfer" {
		eventIndex, _ := strconv.ParseInt(operation.EventIndex, 10, 64)
		operationDb := object.GetOperationByBlock(operation.Block, eventIndex)
		if operationDb.ID > 0 {
			return false
		}
		//update sendingAmount
		err := object.UpdateBalance(operation.Sender, operation.Ticker, operation.TickerHex, 0, operation.Amount, operation.Block, operation.BlockTime)
		if err != nil {
			//update operation status=-2
			status = -2
			logger.GetLogger().Errorf("transfer UpdateBalance err %v", err)
		}
	} else if operation.Operation == "cancel" {
		operationDb := object.GetNoSendOperation(operation.Sender, operation.Additional, operation.Block)
		if operationDb.ID < 1 {
			return false
		}
		//update sendingAmount
		err := object.UpdateBalance(operation.Sender, operation.Ticker, operation.TickerHex, 0, -operationDb.Amount, operation.Block, operation.BlockTime)
		if err != nil {
			logger.GetLogger().Errorf("cancel UpdateBalance err %v", err)
		}
		//update status
		operationDb.Status = -1
		operationDb.CancelBlock = operation.Block
		operationDb.CancelBlockTime = operation.BlockTime
		operationDb.CancelEventIndex = operation.EventIndex
		err = object.db.Save(operationDb).Error
		if err != nil {
			logger.GetLogger().Errorf("SaveOperation cancel err %v", err)
		}
		return false
	}
	operation.Status = status
	result := object.db.Model(&table.Tao20Operation{}).Where("block_num = ? and event_index = ?", operation.Block, operation.EventIndex).FirstOrCreate(&operation)
	if result.Error != nil {
		logger.GetLogger().Errorf("SaveOperation err %v", result.Error)
		return false
	}
	return result.RowsAffected == 1 && status == 0
}

func (object *ChainDB) SendOperation(eventNode vo.EventNode) error {
	operationDb := object.GetNoSendOperation(eventNode.From, eventNode.Amount, eventNode.BlockNumber)
	if operationDb.ID < 1 {
		return nil
	}
	operationDb.Receiver = eventNode.To
	operationDb.SendBlock = eventNode.BlockNumber
	operationDb.SendBlockTime = util.TimeByHeight(eventNode.BlockNumber).UnixMilli()
	operationDb.SendEventIndex = strconv.FormatInt(eventNode.Id, 10)
	operationDb.Status = 1
	err := object.db.Save(&operationDb).Error
	if err != nil {
		logger.GetLogger().Errorf("GetNoSendOperation err %v", err)
		return err
	}
	//update balance
	err = object.UpdateBalance(eventNode.From, operationDb.Ticker, operationDb.TickerHex, -operationDb.Amount, -operationDb.Amount, eventNode.BlockNumber, operationDb.SendBlockTime)
	if err != nil {
		logger.GetLogger().Errorf("GetNoSendOperation err %v", err)
		return err
	}
	err = object.UpdateBalance(eventNode.To, operationDb.Ticker, operationDb.TickerHex, operationDb.Amount, 0, eventNode.BlockNumber, operationDb.SendBlockTime)
	if err != nil {
		logger.GetLogger().Errorf("GetNoSendOperation err %v", err)
	}
	return err
}

func (object *ChainDB) GetNoSendOperation(from string, additional, blockNum int64) (operation table.Tao20Operation) {
	err := object.db.Model(&table.Tao20Operation{}).Where("sender = ? and additional = ? and block_num <= ? and status=0", from, additional, blockNum).First(&operation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.GetLogger().Errorf("GetNoSendOperation err %v", err)
	}
	return
}

func (object *ChainDB) GetOperationByBlock(block, eventIndex int64) (operation table.Tao20Operation) {
	err := object.db.Model(&table.Tao20Operation{}).Where("block_num = ? and event_index = ?", block, eventIndex).First(&operation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.GetLogger().Errorf("GetNoSendOperation err %v", err)
	}
	return
}
