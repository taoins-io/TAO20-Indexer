package tao20

import (
	"tao/database/gdb"
	"tao/database/gdb/table"
)

type DeployOperation struct {
	BaseOperation
}

func (deploy DeployOperation) Handle(tao20Operation table.Tao20Operation) {
	gdb.Inst().SaveTao20(tao20Operation)
}
