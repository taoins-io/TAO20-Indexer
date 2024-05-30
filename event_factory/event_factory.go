package event_factory

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"tao/database/gdb"
	"tao/database/gdb/table"
	"tao/event_factory/tao20"
	"tao/util"
	"tao/vo"
)

var contractTemplate = make(map[string]bool)

const (
	EventType = "ffffffff"
)

func EventAllFactory(eventNode vo.EventNode) {
	contractOnce.Do(func() {
		contractTemplate[EventType] = true
	})
	if _, ok := contractTemplate[eventNode.ToHex[:8]]; !ok {
		//If it does not start with EventType, handle whether it is the send of tao20 transfer.
		checkSend(eventNode)
		return
	}
	// check inscription OR Tao-20
	//ffffffff13000000000000692e74616f7562692e636f6d2f616972642e706e67
	assetType := eventNode.ToHex[8:9]
	contentType := eventNode.ToHex[9:10]
	if assetType+contentType == "20" {
		tickerHex := eventNode.ToHex[10:18]
		ticker := string(common.TrimLeftZeroes(common.Hex2Bytes(tickerHex)))
		operationHex := eventNode.ToHex[18:34]
		operation := string(common.TrimLeftZeroes(common.Hex2Bytes(operationHex)))
		amountHex := eventNode.ToHex[34:46]
		amountVal, _ := strconv.ParseInt(amountHex, 16, 48)
		additionalHex := eventNode.ToHex[48:64]
		additionalVal, _ := strconv.ParseInt(additionalHex, 16, 64)
		tao20Operation := table.Tao20Operation{
			Ticker:     ticker,
			TickerHex:  tickerHex,
			Sender:     eventNode.From,
			To:         eventNode.To,
			ToHex:      eventNode.ToHex,
			Amount:     amountVal,
			Block:      eventNode.BlockNumber,
			BlockTime:  util.TimeByHeight(eventNode.BlockNumber).UnixMilli(),
			EventIndex: strconv.FormatInt(eventNode.Id, 10),
			Operation:  operation,
			Additional: additionalVal,
		}
		tao20.Handle(tao20Operation)
	}
}

func checkSend(eventNode vo.EventNode) {
	/**
	1. Determine whether there is an unconfirmed transfer record of the same from and amount equal to additional operation that has already been recorded.
	2. If there is an unconfirmed record, update the status to confirmed, and update the tao20Balance of the send from and to;
	3. Otherwise, ignore the transaction
	*/
	gdb.Inst().SendOperation(eventNode)
}
