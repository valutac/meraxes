package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/ngurajeka/meraxes/meraxes"
	"github.com/ngurajeka/meraxes/notification"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// CheckCmd command to run the cli version of meraxes
type CheckCmd struct {
	ui     cli.Ui
	logger *zap.Logger
}

// NewCheckCmd create new CheckCmd command
func NewCheckCmd(ui cli.Ui, logger *zap.Logger) *CheckCmd {
	return &CheckCmd{ui, logger}
}

// Help return help text
func (c *CheckCmd) Help() string {
	helpText := `
			Usage: meraxes check [options] ...
			Check URIs.
	`
	return strings.TrimSpace(helpText)
}

// Synopsis return synopsis text
func (c *CheckCmd) Synopsis() string {
	return "Running the checking process"
}

// Run running the checking process
func (c *CheckCmd) Run(args []string) int {
	hosts := viper.GetStringSlice("hosts")

	svc := meraxes.NewService(c.logger)
	notificationSvc := notification.NewService(c.logger)
	notificationSvc.SetTelegramToken(viper.GetString("notification.telegram.token"))

	result := svc.CheckAll(hosts)
	svc.AsTable(result)

	if statuses := result.Errors(); len(statuses) > 0 {
		notificationSvc.Notify(statuses, viper.GetString("notification.email.target"), notification.TypeEmail)
		notificationSvc.Notify(statuses, viper.GetString("notification.telegram.target"), notification.TypeTelegram)
	}

	return 0
}
