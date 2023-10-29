package config

import (
	"api/utils"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
)

var T = DefaultConfig

func InitConfig(configFiles ...string) error {
	configFile := utils.GetExistFile("conf/config.toml", configFiles...)
	if !utils.FileExist(configFile) {
		_ = saveConfig(configFile)
	}
	tree, err := toml.LoadFile(configFile)
	if err != nil {
		return err
	}
	err = tree.Unmarshal(&T)
	if err != nil {
		return err
	}
	_ = saveConfig(configFile)
	return nil
}

func saveConfig(file string) error {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	marshal, err := toml.Marshal(T)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, marshal, os.ModePerm)
	return err
}
