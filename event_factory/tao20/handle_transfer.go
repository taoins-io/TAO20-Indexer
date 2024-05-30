package tao20

import (
	"tao/database/gdb"
	"tao/database/gdb/table"
)

type TransferOperation struct {
	BaseOperation
}

func (transfer TransferOperation) Handle(tao20Operation table.Tao20Operation) {
	gdb.Inst().SaveOperation(tao20Operation)
}
