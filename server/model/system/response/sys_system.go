package response

import "github.com/bernylinville/ops/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
