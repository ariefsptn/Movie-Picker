package main

import (
	_ "Movie-Picker/routers"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"
)

func main() {
	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()
	
	beego.Run()
}

