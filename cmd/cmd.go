package cmd

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/githinkcn/whale/config"
)

func init() {
	config.InitDB()
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}

func Execute() {
	if flag, _ := beego.AppConfig.Bool("cors"); flag == true { //开关
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true,
		}))
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	usae()
	beego.Run()
}
func usae() {
	usageStr := `
 _  _  _  _             _           __      ______  _                    _ 
| || || || |           | |         / /     / _____)| |                  | |
| || || || | _    ____ | |  ____  / /____ | /      | |  ___   _   _   _ | |
| ||_|| || || \  / _  || | / _  )|___   _)| |      | | / _ \ | | | | / || |
| |___| || | | |( ( | || |( (/ /     | |  | \_____ | || |_| || |_| |( (_| |
 \______||_| |_| \_||_||_| \____)    |_|   \______)|_| \___/  \____| \____|

`
	fmt.Printf("%s\n", usageStr)
}
