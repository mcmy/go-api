package test

import (
	"api/config"
	"api/drivers/db"
	"api/drivers/db/model"
	"api/i18n"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)
import "gorm.io/gen"

func init() {
	log.Println("default lang set:", i18n.DefaultLang)
	if err := config.InitConfig("../conf/config.toml"); err != nil {
		log.Fatalln("config init error:", err)
	}
}

func TestGenModel(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../drivers/db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		//FieldWithTypeTag: true,
		//FieldWithIndexTag: true,
		//FieldNullable:     true,
		//FieldCoverable:    true,
		//FieldSignable:     true,
	})

	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"datetime": func(columnType gorm.ColumnType) (dataType string) { return "string" },
	}

	g.WithDataTypeMap(dataMap)

	gdb, _ := gorm.Open(mysql.Open(config.T.DB.DataSource))
	g.UseDB(gdb)
	g.ApplyBasic(
		append(g.GenerateAllTable())...,
	)
	g.Execute()
}

func TestGenMethod(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../drivers/db/query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldSignable: true,
		//FieldWithTypeTag:  true,
		//FieldWithIndexTag: true,
		//FieldNullable:     true,
		//FieldCoverable:    true,
	})

	gdb, _ := gorm.Open(mysql.Open(config.T.DB.DataSource))
	g.UseDB(gdb)

	//g.ApplyBasic(model.User{}, model.UserAuth{})
	g.ApplyInterface(func(db.UserMethod) {}, model.User{})
	g.ApplyInterface(func(db.UserAuthMethod) {}, model.UserAuth{})

	// Generate default DAO interface for those generated structs from database
	//companyGenerator := g.GenerateModelAs("company", "MyCompany")
	//g.ApplyBasic(
	//	g.GenerateModel("users"),
	//	companyGenerator,
	//	g.GenerateModelAs("people", "Person",
	//		gen.FieldIgnore("deleted_at"),
	//		gen.FieldNewTag("age", `json:"-"`),
	//	),
	//)

	g.Execute()
}
