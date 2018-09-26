package main

import (
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"

	_ "mindoc/routers"
	"mindoc/commands/daemon"
	"github.com/kardianos/service"
	"fmt"
	"os"
	"mindoc/commands"
)

func main() {

	commands.RegisterCommand()
	d := daemon.NewDaemon()
	s, err := service.New(d, d.Config())
	if err != nil {
		fmt.Println("Create service error => ", err)
		os.Exit(1)
	}

	s.Run()
}
