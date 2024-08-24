package common

import (
	"os"
	"pledge-backend-practise/log"
)

func init() {
	log.Logger.Info("init private key.")
	if key, ok := os.LookupEnv("plgr_admin_private_key"); ok {
		PlgrAdminPrivateKey = key
	} else {
		//panic("plgr_admin_private_key is required")
		log.Logger.Error("plgr_admin_private_key is required")
	}
}

var PlgrAdminPrivateKey string
