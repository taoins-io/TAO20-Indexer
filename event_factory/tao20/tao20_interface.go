package tao20

import (
	"tao/database/gdb/table"
)

type OperationHandler interface {
	Handle(tao20Operation table.Tao20Operation)
}
