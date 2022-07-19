package cmd

import (
	"fmt"

	"github.com/itrepablik/isked"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal/tasks"
	"github.com/onumahkalusamuel/bookieguard/internal/webserver"
)

func Setup() {
	fmt.Println("Installing and setting up...")

	// start the webserver
	go webserver.StartWebServer()

	// launch admin panel in browser
	go webserver.LaunchBrowser()

	// set update checks
	isked.TaskName("update_checks").
		Frequently().
		Minutes(config.ISKED_UPDATES).
		ExecFunc(tasks.Update).
		AddTask()

	// report system status
	isked.TaskName("system_status").
		Frequently().
		Minutes(config.ISKED_SYSTEM_STATUS).
		ExecFunc(tasks.SystemStatus).AddTask()

	// post gathered hosts
	isked.TaskName("send_hosts").
		Frequently().
		Minutes(config.ISKED_SEND_HOSTS).
		ExecFunc(tasks.SendHosts).AddTask()

	// run application
	isked.Run()
}
