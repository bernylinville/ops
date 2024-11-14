package request

import (
	"github.com/bernylinville/ops/server/model/common/request"
	"github.com/bernylinville/ops/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
