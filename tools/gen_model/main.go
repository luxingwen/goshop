package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:mysql57@tcp(localhost:8300)/gostore?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	tables, err := engine.DBMetas()
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(tables, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_ = b
	//fmt.Println(string(b))

	for _, itemTable := range tables {
		genStruct(itemTable)
	}

	// fmt.Println(itemTable)
}

func genStruct(table *schemas.Table) {

	if table.Name != "eb_store_order" {
		return
	}
	fmt.Println("//", table.Name)
	fmt.Println("//", table.Comment)
	name := strings.TrimPrefix(table.Name, "eb_")
	structName := toTitle(name)
	fmt.Printf("type %s struct {\n", structName)

	cols := table.Columns()

	b, err := json.MarshalIndent(cols, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	genFlieds(table.Columns())
	fmt.Println("}\n")
}

func toTitle(s string) (r string) {
	for _, item := range strings.Split(s, "_") {
		r += strings.Title(item)
	}
	return

}

func genFlieds(cs []*schemas.Column) {

	for _, item := range cs {
		typ := getTyp(item.SQLType)
		fmt.Printf("\t%s\t%s\t`gorm:\"column:%s\"`  //%s \n", toTitle(item.Name), typ, item.Name, item.Comment)

	}
}

func getTyp(sqlTyp schemas.SQLType) (r string) {
	switch true {
	case strings.ToUpper(sqlTyp.Name) == "DECIMAL":
		return "float64"
	case sqlTyp.IsNumeric():
		return "int"
	case strings.ToUpper(sqlTyp.Name) == "BOOL" || strings.ToLower(sqlTyp.Name) == "BOOLEAN":
		return "bool"
	}
	return "string"
}
