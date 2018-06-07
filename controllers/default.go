package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) IndexGet() {

	fmt.Println("1111111111111111111111111")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["json"] = "xxxxxxxxxxxx"
	c.TplName = "esayuiDemo/demo/datagrid/basic.html"
}

func (c *MainController) GetResult() {

	if c.IsAjax() {

		println("xxxxxxxxxxxxxxxxxx11111111111111111")
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pwd)
		path, err := filepath.Abs("./controllers/datagrid_data1.json")
		fmt.Println(path, err)
		data, err := ioutil.ReadFile(path)
		fmt.Println("xxxxxxxxxxxxxxxxxxxx", err)
		jsonData := gjson.Parse(string(data))
		total := jsonData.Get("total").Int()

		rows := jsonData.Get("rows").String()
		fmt.Println("xxxxxxxxxxxxxxxxxxx", total, rows)

		//{field:'Id',title:'ID',hidden:true},
		//{field:'Title',title:'显示名',width:150},
		//{field:'Name',title:'应用名',width:150}

		ss := make(map[string]interface{})
		ss["Id"] = 1
		ss["Title"] = "go"
		ss["Name"] = "工程学设计"

		nodes := []orm.Params{ss}

		c.Data["json"] = &map[string]interface{}{"total": total, "rows": &nodes}
		//c.Data["json"] = string(data)

		c.ServeJSON()
	} else {

		println("xxxxxxxxxxxxxxxxxx222222222222")

		c.Data["grouplist"] = "111111"
		c.Data["roleid"] = "2222222222"
		c.TplName = "esayuiDemo/demo/datagrid/fluid.html"

	}

	//c.TplName = "esayuiDemo/demo/datagrid/fluid.html"
}
