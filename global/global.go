package global

import (
	"github.com/spf13/viper"
	"gin-create/config"
)

var (
	GVA_VP		*viper.Viper
	GVA_CONFIG	config.Server
)
