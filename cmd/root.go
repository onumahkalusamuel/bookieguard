package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/itrepablik/isked"
	"github.com/kardianos/service"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal"
	"github.com/onumahkalusamuel/bookieguard/internal/webserver"
	"github.com/spf13/cobra"
)

var logger service.Logger

var rootCmd = &cobra.Command{
	Use:     config.AppName,
	Short:   config.AppDesc,
	Long:    config.AppLongDesc,
	Version: config.AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing and starting...")
		MakeService()
	},
}

// Execute...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	ct.Foreground(ct.Green, false)
	myFigure := figure.NewFigure(config.AppShortDisplayName, "", true)
	myFigure.Print()
	fmt.Println()
	ct.ResetColor()
}

func MakeService() {
	svcConfig := &service.Config{
		Name:        config.AppName,
		DisplayName: config.AppDisplayName,
		Description: config.AppDesc,
	}
	runAsService(svcConfig, func() {

		// start the webserver
		go webserver.StartWebServer()

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
	})
}

func runAsService(svcConfig *service.Config, run func()) error {
	s, err := service.New(&program{exec: run}, svcConfig)
	if err != nil {
		return err
	}
	logger, err = s.Logger(nil)
	if err != nil {
		return err
	}
	return s.Run()
}

type program struct {
	exec func()
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.exec()
	return nil
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
