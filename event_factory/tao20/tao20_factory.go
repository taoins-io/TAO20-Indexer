package tao20

import (
	"sync"
	"tao/database/gdb/table"
)

var OperationTemplate = make(map[string]OperationHandler)

const (
	Deploy   = "deploy"
	Mint     = "mint"
	Transfer = "transfer"
	Cancel   = "cancel"
)

type BaseOperation struct {
}

var OperationOnce sync.Once

func Handle(tao20Operation table.Tao20Operation) {
	OperationOnce.Do(func() {
		OperationTemplate[Deploy] = new(DeployOperation)
		OperationTemplate[Mint] = new(MintOperation)
		OperationTemplate[Transfer] = new(TransferOperation)
		OperationTemplate[Cancel] = new(CancelOperation)
	})
	if _, ok := OperationTemplate[tao20Operation.Operation]; !ok {
		return
	}
	OperationTemplate[tao20Operation.Operation].Handle(tao20Operation)
}
