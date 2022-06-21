package main

import (
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"
	"github.com/yohamta/dagu/internal/admin"
	"github.com/yohamta/dagu/internal/runner"
	"github.com/yohamta/dagu/internal/utils"
)

func newSchedulerCommand() *cli.Command {
	l := &admin.Loader{}
	return &cli.Command{
		Name:  "scheduler",
		Usage: "dagu scheduler",
		Action: func(c *cli.Context) error {
			cfg, err := l.LoadAdminConfig(
				path.Join(utils.MustGetUserHomeDir(), ".dagu/admin.yaml"))
			if err == admin.ErrConfigNotFound {
				cfg = admin.DefaultConfig()
			} else if err != nil {
				return err
			}
			return startScheduler(cfg)
		},
	}
}

func startScheduler(cfg *admin.Config) error {
	agent := runner.NewAgent(cfg)

	listenSignals(func(sig os.Signal) {
		agent.Stop()
	})

	log.Printf("starting dagu scheduler")
	return agent.Start()
}
