package cmd

import (
	"github.com/urfave/cli/v2"
	"go-netdisk/internal/db"
	"go-netdisk/internal/db/initial"
	"go-netdisk/internal/settings"
	"log"
)

var Migrate = &cli.Command{
	Name:        "init",
	Usage:       "Initialize database",
	Description: `Create table and insert initial data.`,
	Action:      runInit,
	Flags: []cli.Flag{
		stringFlag("config", "", "Custom configuration file path", []string{"c"}),
		boolFlag("verbose, v", "Show process details", []string{"v"}),
	},
}

func runInit(c *cli.Context) error {
	cfg := settings.GetCfg()
	cfg.LoadSettings(c)
	_, _ = db.InitDB()
	initial.InitData()
	log.Println("Init database success.")
	return nil
}
