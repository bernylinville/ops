package initialize

import (
	"github.com/bernylinville/ops/server/global"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}