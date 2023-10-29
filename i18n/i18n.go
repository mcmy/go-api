package i18n

import (
	"api/utils"
	"github.com/pelletier/go-toml"
	"log"
)

var (
	L           = map[string]map[string]string{}
	DefaultLang string
)

func InitI18n(i18nFiles ...string) error {
	i18nFile := utils.GetExistFile("conf/lang.toml", i18nFiles...)
	tree, err := toml.LoadFile(i18nFile)
	if err != nil {
		return err
	}
	err = tree.Unmarshal(&L)
	if err != nil {
		return err
	}
	DefaultLang = L["config"]["default"]
	if DefaultLang == "" {
		log.Fatalf("default lang is empty")
	}
	return nil
}

func Use(lang string) MappingInterface {
	d := &DefaultImpl{}
	d.Use(lang)
	return d
}

func Get(mapping string) string {
	return Use(DefaultLang).Get(mapping)
}
