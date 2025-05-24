package config

import "time"

type SettingConfig struct {
	SettingLinkVisitSnapshotInterval time.Duration `env:"SETTING_LINK_VISIT_SNAPSHOT_INTERVAL" default:"1h"`
}

var settingConfig SettingConfig

func LoadSettingConfig() SettingConfig {
	if settingConfig != (SettingConfig{}) {
		return settingConfig
	}

	loadConfig(&settingConfig)

	return settingConfig
}
