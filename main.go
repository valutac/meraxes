package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/ngurajeka/meraxes/command"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("server.port", 5000)
	viper.SetDefault("sheduler.time", 60)
	_ = viper.ReadInConfig()
}

var ui cli.Ui

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}
	logger, _ := zapConfig.Build()
	defer logger.Sync()

	ui = &cli.BasicUi{Writer: os.Stdout}

	commands := &cli.CLI{
		Args: os.Args[1:],
		Commands: map[string]cli.CommandFactory{
			"check": func() (cli.Command, error) {
				return command.NewCheckCmd(ui, logger), nil
			},
			"web": func() (cli.Command, error) {
				return command.NewWebCmd(ui, logger), nil
			},
		},
		HelpFunc: cli.BasicHelpFunc("meraxes"),
		Version:  "1.0.0",
	}

	exitCode, err := commands.Run()
	if err != nil {
		logger.Error("error executing meraxes",
			zap.Strings("args", os.Args),
			zap.Error(err),
		)
		os.Exit(1)
	}

	os.Exit(exitCode)
}
