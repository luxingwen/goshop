package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/crmeb_hh?charset=utf8")

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

	routerStr := ""
	for _, itemTable := range tables {

		genStruct(itemTable)
		genController(itemTable)
		if filterTable(itemTable.Name) {
			continue
		}
		routerStr += genRouter(itemTable)
		routerStr += "\n"
	}

	genMainRouter(routerStr)

	// fmt.Println(itemTable)
}

func filterTable(name string) bool {
	s := []string{
		"eb_store_combination_attr",
		"eb_store_combination_attr_result",
		"eb_store_combination_attr_value",
		"eb_store_coupon_issue_user",
		"eb_store_order_cart_info",
		"eb_store_order_status",
		"eb_store_product_attr",
		"eb_store_product_attr_result",
		"eb_store_product_attr_value",
		"eb_store_product_relation",
		"eb_store_seckill_attr",
		"eb_store_seckill_attr_result",
		"eb_store_seckill_attr_value",
		"eb_system_attachment",
		"eb_wechat_user",
		"eb_article_content",
		"eb_cache",
	}

	for _, item := range s {
		if name == item {
			return true
		}
	}
	return false
}

func genStruct(table *schemas.Table) {

	name := strings.TrimPrefix(table.Name, "eb_")
	structName := toTitle(name)
	modelName := strings.ToLower(string(structName[0])) + structName[1:]

	structContent := genFlieds(table.Columns())
	tpls := getDefaultModelTpl()

	tpls = strings.Replace(tpls, "{{structDesc}}", table.Comment, 1)
	tpls = strings.Replace(tpls, "{{ModelName}}", structName, -1)
	tpls = strings.Replace(tpls, "{{modelName}}", modelName, -1)
	tpls = strings.Replace(tpls, "{{tableName}}", table.Name, 1)
	tpls = strings.Replace(tpls, "{{modelStruct}}", structContent, -1)
	fmt.Println(tpls)

	filename := "gen_" + name + ".go"

	f, err := os.Create("gen/" + filename)
	if err != nil {
		panic(err)
	}
	f.Write([]byte(tpls))
	f.Close()
	cmd := exec.Command("gofmt", "-w", "gen/"+filename)
	if err := cmd.Run(); err != nil {
		log.Printf("Error while running gofmt: %s", err)
	}

}

// 生成 controller
func genController(table *schemas.Table) {
	name := strings.TrimPrefix(table.Name, "eb_")
	structName := toTitle(name)
	modelName := strings.ToLower(string(structName[0])) + structName[1:]

	tpls := getDefaultControllerTpl()

	tpls = strings.Replace(tpls, "{{modelName}}", modelName, -1)
	tpls = strings.Replace(tpls, "{{ModelName}}", structName, -1)

	filename := "gen_" + name + ".go"
	f, err := os.Create("gen_controllers/" + filename)
	if err != nil {
		panic(err)
	}
	f.WriteString(tpls)
	f.Close()
	cmd := exec.Command("gofmt", "-w", "gen_controllers/"+filename)
	if err := cmd.Run(); err != nil {
		log.Printf("Error while running gofmt: %s", err)
	}

}

func genRouter(table *schemas.Table) string {
	name := strings.TrimPrefix(table.Name, "eb_")
	structName := toTitle(name)
	modelName := strings.ToLower(string(structName[0])) + structName[1:]

	tpls := getDefaultRouterTpl()

	tpls = strings.Replace(tpls, "{{modelName}}", modelName, -1)
	tpls = strings.Replace(tpls, "{{ModelName}}", structName, -1)
	return tpls

}

func genMainRouter(str string) {
	tpls := getDefaultRoutersTpls()
	tpls = strings.Replace(tpls, "{{RoutersTpls}}", str, -1)

	filename := "gen_router.go"
	f, err := os.Create("gen_router/" + filename)
	if err != nil {
		panic(err)
	}
	f.WriteString(tpls)
	f.Close()
	cmd := exec.Command("gofmt", "-w", "gen_router/"+filename)
	if err := cmd.Run(); err != nil {
		log.Printf("Error while running gofmt: %s", err)
	}

}

func toTitle(s string) (r string) {
	for _, item := range strings.Split(s, "_") {
		r += strings.Title(item)
	}
	return
}

func genFlieds(cs []*schemas.Column) (r string) {
	for _, item := range cs {
		typ := getTyp(item.SQLType)
		r += fmt.Sprintf("\t%s\t%s\t`gorm:\"column:%s\"`  //%s \n", toTitle(item.Name), typ, item.Name, item.Comment)
	}
	return
}

func getTyp(sqlTyp schemas.SQLType) (r string) {
	switch true {
	case strings.ToUpper(sqlTyp.Name) == "DECIMAL":
		return "float64"
	case sqlTyp.IsNumeric():
		return "int"
	case strings.ToUpper(sqlTyp.Name) == "BOOL" || strings.ToLower(sqlTyp.Name) == "BOOLEAN":
		return "bool"
	case sqlTyp.IsTime():
		return "time.Time"
	}
	return "string"
}

func getDefaultModelTpl() string {
	return `//generate by gen
package models

import (
	"goshop/restful/common"
)

//{{structDesc}}
type {{ModelName}} struct {
{{modelStruct}}
}


//修改默认表名
func ({{ModelName}}) TableName() string {
	return "{{tableName}}"
}

func ({{modelName}} *{{ModelName}}) Insert() error {
	err :=  common.GetDB().Create({{modelName}}).Error
	return err
}

func ({{modelName}} *{{ModelName}}) Patch() error {
	err :=  common.GetDB().Model({{modelName}}).Updates({{modelName}}).Error
	return err
}

func ({{modelName}} *{{ModelName}}) Update() error {
	err :=  common.GetDB().Save({{modelName}}).Error
	return err
}

func ({{modelName}} *{{ModelName}}) Delete() error {
	return  common.GetDB().Delete({{modelName}}).Error
}

func ({{modelName}} *{{ModelName}}) List(rawQuery string, rawOrder string, offset int, limit int) (*[]{{ModelName}}, int, error) {
	{{modelName}}s := []{{ModelName}}{}
	total := 0
	db :=  common.GetDB().Model({{modelName}})
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &{{modelName}}s, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &{{modelName}}s, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&{{modelName}}s).
		Count(&total)
	err = db.Error
	return &{{modelName}}s, total, err
}

func ({{modelName}} *{{ModelName}}) Get() (*{{ModelName}}, error) {
	err :=  common.GetDB().Find(&{{modelName}}).Error
	return {{modelName}}, err
}
`
}

func getDefaultControllerTpl() string {
	return `//generate by gen
package controllers
import (
	"goshop/restful/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type {{ModelName}}Controller struct {
}

// @Summary Create
// @Tags    {{ModelName}}
// @Param body body models.{{ModelName}} true "{{ModelName}}"
// @Success 200 {string} string ""
// @Router /{{modelName}}s [post]
func (ctl *{{ModelName}}Controller) Create(c *gin.Context) {
	{{modelName}} := models.{{ModelName}}{
	}
	if err := ParseRequest(c, &{{modelName}}); err != nil {
		return
	}
	if err := {{modelName}}.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, {{modelName}})
}
// @Summary  Delete
// @Tags     {{ModelName}}
// @Param  {{modelName}}Id  path string true "{{modelName}}Id"
// @Success 200 {string} string ""
// @Router /{{modelName}}s/{{{modelName}}Id} [delete]
func (ctl *{{ModelName}}Controller) Delete(c *gin.Context) {
	{{modelName}} := models.{{ModelName}}{}
	id := c.Param("{{modelName}}Id")
	var err error
	{{modelName}}.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = {{modelName}}.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
// @Summary Put
// @Tags    {{ModelName}}
// @Param body body models.{{ModelName}} true "{{modelName}}"
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {string} string ""
// @Router /{{modelName}}s/{{{modelName}}Id} [put]
func (ctl *{{ModelName}}Controller) Put(c *gin.Context) {
	{{modelName}} := models.{{ModelName}}{}
	id := c.Param("{{modelName}}Id")
	var err error
	{{modelName}}.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &{{modelName}}); err != nil {
		return
	}
	err = {{modelName}}.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
// @Summary Patch
// @Tags    {{ModelName}}
// @Param body body models.{{ModelName}} true "{{modelName}}"
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {string} string ""
// @Router /{{modelName}}s/{{{modelName}}Id} [patch]
func (ctl *{{ModelName}}Controller) Patch(c *gin.Context) {
	{{modelName}} := models.{{ModelName}}{}
	id := c.Param("{{modelName}}Id")
	var err error
	{{modelName}}.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &{{modelName}}); err != nil {
		return
	}
	err = {{modelName}}.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
// @Summary List
// @Tags    {{ModelName}}
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.{{ModelName}} "{{modelName}} array"
// @Router /{{modelName}}s [get]
func (ctl *{{ModelName}}Controller) List(c *gin.Context) {
	{{modelName}} := &models.{{ModelName}}{}
	var err error
	pageParam := c.DefaultQuery("page", "-1")
	pageSizeParam := c.DefaultQuery("pageSize", "-1")
	rawQuery := c.DefaultQuery("query", "")
	rawOrder := c.DefaultQuery("order", "")
	pageInt, err := strconv.Atoi(pageParam)
	pageSizeInt, err := strconv.Atoi(pageSizeParam)
	offset := pageInt*pageSizeInt - pageSizeInt
	limit := pageSizeInt
	if pageInt < 0 || pageSizeInt < 0 {
		limit = -1
	}
	{{modelName}}s, total, err := {{modelName}}.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data": {{modelName}}s,
	})
}
// @Summary Get
// @Tags    {{ModelName}}
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {object} models.{{ModelName}} "{{modelName}} object"
// @Router /{{modelName}}s/{{{modelName}}Id} [get]
func (ctl *{{ModelName}}Controller) Get(c *gin.Context) {
	{{modelName}} := &models.{{ModelName}}{}
	id := c.Param("{{modelName}}Id")

	var err error
	{{modelName}}.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}


	{{modelName}}, err = {{modelName}}.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, {{modelName}})
}
`
}

func getDefaultRouterTpl() string {
	return `
	{{modelName}}Controller := controllers.{{ModelName}}Controller{}
	{{modelName}}Group := r.Group("/{{modelName}}s")
	{
		{{modelName}}Group.GET("", {{modelName}}Controller.List)
		{{modelName}}Group.POST("", {{modelName}}Controller.Create)
		{{modelName}}Group.DELETE("/:{{modelName}}Id", {{modelName}}Controller.Delete)
		{{modelName}}Group.PUT("/:{{modelName}}Id", {{modelName}}Controller.Put)
		{{modelName}}Group.GET("/:{{modelName}}Id", {{modelName}}Controller.Get)
		{{modelName}}Group.PATCH("/:{{modelName}}Id", {{modelName}}Controller.Patch)
	}
	//!!do not delete gen will generate router code at here
`
}

func getDefaultRoutersTpls() string {
	return `	
//generate by gen
package routers

import (
	"github.com/gin-gonic/gin"

	"goshop/restful/controllers"
)


func GenRouters(r *gin.Engine){
	{{RoutersTpls}}
}`
}
