package main

import (
	"github.com/kardianos/service"
	"github.com/onumahkalusamuel/bookieguard/config"
	"github.com/onumahkalusamuel/bookieguard/internal/webserver"
)

var logger service.Logger

func main() {

	svcConfig := &service.Config{
		Name:        config.AppServiceName,
		DisplayName: config.AppDisplayName,
		Description: config.AppDesc,
	}
	runAsService(svcConfig, func() {
		go webserver.StartWebServer()
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
