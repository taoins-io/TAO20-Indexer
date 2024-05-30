package tao20

import (
	"tao/database/gdb"
	"tao/database/gdb/table"
)

type CancelOperation struct {
	BaseOperation
}

func (cancel CancelOperation) Handle(tao20Operation table.Tao20Operation) {
	gdb.Inst().SaveOperation(tao20Operation)
}
