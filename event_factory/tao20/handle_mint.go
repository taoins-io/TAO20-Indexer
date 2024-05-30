package tao20

import (
	"tao/database/gdb"
	"tao/database/gdb/table"
)

type MintOperation struct {
	BaseOperation
}

func (mint MintOperation) Handle(tao20Operation table.Tao20Operation) {
	gdb.Inst().Tao20Mint(tao20Operation)
}
