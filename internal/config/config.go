package config

import "log"

var (
	LogCfg LogConfig
	SysCfg SystemConfig
)

func Init() {
	err := LogCfg.readConf("config/log.yml")
	if err != nil {
		log.Fatalf("[system] logCfg readConf err: %v\n", err)
	}

	err = SysCfg.readConf("config/system.yml")
	if err != nil {
		log.Fatalf("[system] sysCfg readConf err: %v\n", err)
	}
}
