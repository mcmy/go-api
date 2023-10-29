package test

import (
	"api/i18n"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGenI18n(t *testing.T) {
	err := I18n2Struct("../conf/lang.toml", "../i18n/i18nStruct.go")
	log.Println(err)
}

func I18n2Struct(inputToml string, outputGoStruct string) error {
	err := i18n.InitI18n(inputToml)
	if err != nil {
		return err
	}
	out := ""

	//out += "\ntype I18n struct {\n"
	//for k := range i18n.L {
	//	KType := "LangStruct"
	//	k1 := k
	//	if k == "config" {
	//		KType = "LangConfig"
	//		k1 = "LangConfig"
	//	}
	//	out += fmt.Sprintf("%s %s  `toml:\"%s\"`\n", ToCamel(strings.ReplaceAll(k1, "-", "")), KType, k)
	//}
	//out += "\n}\n"

	//out += "\ntype LangConfig struct{\n"
	//c := i18n.L["config"]
	//for k := range c {
	//	out += fmt.Sprintf("%s string `toml:\"%s\"`\n", ToCamel(k), k)
	//}
	//out += "\n}\n"

	{
		out += "\ntype LangStruct struct {\n"
		m := i18n.L[i18n.DefaultLang]
		for k := range m {
			out += fmt.Sprintf("%s string `toml:\"%s\"`\n", ToCamel(k), k)
		}
		out += "\n}\n"
	}

	{
		out += "\nvar (\n"
		m := i18n.L[i18n.DefaultLang]
		for k := range m {
			out += fmt.Sprintf("%s =\"%s\"\n", ToCamel(k), k)
		}
		out += "\n)\n"
	}

	f, err := os.Create(outputGoStruct)
	if err != nil {
		fmt.Println("Can not write file")
		return err
	}
	defer f.Close()

	_, err = f.WriteString("package i18n\n\n" + out)
	if err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", outputGoStruct)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ToCamel(s string) string {
	parts := strings.Split(s, "-")
	for i, part := range parts {
		if part != "" {
			parts[i] = strings.ToUpper(string(part[0])) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func ToCamel_(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if part != "" {
			parts[i] = strings.ToUpper(string(part[0])) + part[1:]
		}
	}
	return strings.Join(parts, "")
}
