package request

import (
	"github.com/bernylinville/ops/server/model/common/request"
	"github.com/bernylinville/ops/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
