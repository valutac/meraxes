package command

import (
	"github.com/Valutac/meraxes/meraxes"
	"github.com/Valutac/meraxes/notification"
	"github.com/mitchellh/cli"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
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

	if viper.GetBool("notification.active") {
		if statuses := result.Errors(); len(statuses) > 0 {
			_ = notificationSvc.Notify(statuses, viper.GetString("notification.email.target"), notification.TypeEmail)
			_ = notificationSvc.Notify(statuses, viper.GetString("notification.telegram.target"), notification.TypeTelegram)
		}
	}

	return 0
}
