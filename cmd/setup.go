package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/itrepablik/isked"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/onumahkalusamuel/bookieguard/internal/webserver"
)

func Setup() {

	ct.Foreground(ct.Green, false)
	myFigure := figure.NewFigure(config.AppShortDisplayName, "", true)
	myFigure.Print()
	fmt.Println()
	ct.ResetColor()

	fmt.Println("Installing and setting up...")

	// start the webserver
	go webserver.StartWebServer()

	// launch admin panel in browser
	go webserver.LaunchBrowser()

	// set update checks
	isked.TaskName("update_checks").
		Frequently().
		Minutes(config.ISKED_UPDATES).
		ExecFunc(internal.Update).
		AddTask()

	// post gathered hosts
	isked.TaskName("send_hosts").
		Frequently().
		Minutes(config.ISKED_SEND_HOSTS).
		ExecFunc(internal.SendHosts).AddTask()

	// run application
	isked.Run()
}
