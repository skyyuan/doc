package commands

import (
	"fmt"
	"os"
	"github.com/astaxie/beego/orm"
	"flag"
	"github.com/astaxie/beego"
	"mindoc/conf"
	"mindoc/models"
	"mindoc/utils"
)

//系统安装.
func Install() {

	fmt.Println("Initializing...")

	err := orm.RunSyncdb("default", false, true)
	if err == nil {
		initialization()
	} else {
		panic(err.Error())
		os.Exit(1)
	}
	fmt.Println("Install Successfully!")
	os.Exit(0)

}

func Version() {
	if len(os.Args) >= 2 && os.Args[1] == "version" {
		fmt.Println(conf.VERSION)
		os.Exit(0)
	}
}

//修改用户密码
func ModifyPassword() {
	var account,password string

	//账号和密码需要解析参数后才能获取
	if len(os.Args) >= 2 && os.Args[1] == "password" {
		flagSet := flag.NewFlagSet("MinDoc command: ", flag.ExitOnError)

		flagSet.StringVar(&account, "account", "", "用户账号.")
		flagSet.StringVar(&password, "password", "", "用户密码.")

		if err := flagSet.Parse(os.Args[2:]); err != nil {
			beego.Error("解析参数失败 -> ",err)
			os.Exit(1)
		}

		if len(os.Args) < 2 {
			fmt.Println("Parameter error.")
			os.Exit(1)
		}

		if account == "" {
			fmt.Println("Account cannot be empty.")
			os.Exit(1)
		}
		if password == "" {
			fmt.Println("Password cannot be empty.")
			os.Exit(1)
		}
		member,err := models.NewMember().FindByAccount(account)

		if err != nil {
			fmt.Println("Failed to change password:",err)
			os.Exit(1)
		}
		pwd,err := utils.PasswordHash(password)

		if err != nil {
			fmt.Println("Failed to change password:",err)
			os.Exit(1)
		}
		member.Password = pwd

		err = member.Update("password")
		if err != nil {
			fmt.Println("Failed to change password:",err)
			os.Exit(1)
		}
		fmt.Println("Successfully modified.")
		os.Exit(0)
	}


}

//初始化数据
func initialization() {

	err := models.NewOption().Init()

	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}

	member, err := models.NewMember().FindByFieldFirst("account", "admin")
	if err == orm.ErrNoRows {

		member.Account = "admin"
		member.Avatar = conf.URLForWithCdnImage("/static/images/headimgurl.jpg")
		member.Password = "itering123"
		member.AuthMethod = "local"
		member.Role = 0
		member.Email = "skyyuan@itering.com"

		if err := member.Add(); err != nil {
			panic("Member.Add => " + err.Error())
			os.Exit(0)
		}
	}
}
