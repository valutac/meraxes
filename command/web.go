package command

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/ngurajeka/meraxes/meraxes"
	"github.com/ngurajeka/meraxes/notification"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/mitchellh/cli"
	"go.uber.org/zap"
)

// WebCmd command to run the web version of meraxes
type WebCmd struct {
	ui     cli.Ui
	logger *zap.Logger
}

// NewWebCmd create new WebCmd command
func NewWebCmd(ui cli.Ui, logger *zap.Logger) *WebCmd {
	return &WebCmd{ui, logger}
}

// Help return help text
func (c *WebCmd) Help() string {
	helpText := `
			Usage: meraxes web [options] ...
			Running Meraxes Web Version.
	`
	return strings.TrimSpace(helpText)
}

// Synopsis return synopsis text
func (c *WebCmd) Synopsis() string {
	return "Running the web app"
}

// Run running the web server and scheduler
func (c *WebCmd) Run(args []string) int {
	serverPort := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	format := "2006-01-02 15:04:05"

	var (
		result  = meraxes.NewResult()
		hosts   = viper.GetStringSlice("hosts")
		lastRun *time.Time
		nextRun = time.Now().Add(time.Minute * time.Duration(viper.GetInt("sheduler.time")))

		meraxesSvc      = meraxes.NewService(c.logger)
		notificationSvc = notification.NewService(c.logger)
	)

	notificationSvc.SetTelegramToken(viper.GetString("notification.telegram.token"))
	notificationSvc.SetSMTP(
		viper.GetString("notification.email.sender"),
		viper.GetString("notification.email.host"),
		viper.GetString("notification.email.username"),
		viper.GetString("notification.email.password"),
		viper.GetInt("notification.email.port"))

	job := func() {
		result = meraxesSvc.CheckAll(hosts)
		now := time.Now()
		lastRun = &now
		nextRun = now.Add(time.Minute * time.Duration(viper.GetInt("sheduler.time")))
		if statuses := result.Errors(); len(statuses) > 0 {
			go notificationSvc.NotifyAll(statuses, viper.GetStringSlice("notification.email.target"), notification.TypeEmail)
			go notificationSvc.NotifyAll(statuses, viper.GetStringSlice("notification.telegram.target"), notification.TypeTelegram)
		}
	}

	scheduler.Every(viper.GetInt("scheduler.time")).Minutes().NotImmediately().Run(job)

	context := func() gin.H {
		ctx := gin.H{"statuses": result.Statuses, "nextRun": nextRun.Format(format)}
		if lastRun != nil {
			ctx["lastRun"] = lastRun.Format(format)
		} else {
			ctx["lastRun"] = "-"
		}
		return ctx
	}

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	if viper.GetString("auth.username") != "" {
		r.Use(gin.BasicAuth(gin.Accounts{
			viper.GetString("auth.username"): viper.GetString("auth.password"),
		}))
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", context())
	})

	r.GET("/result", func(c *gin.Context) {
		c.JSON(http.StatusOK, context())
	})

	r.POST("/trigger", func(c *gin.Context) {
		result = meraxesSvc.CheckAll(hosts)
		now := time.Now()
		lastRun = &now
		if statuses := result.Errors(); len(statuses) > 0 {
			go notificationSvc.NotifyAll(statuses, viper.GetStringSlice("notification.email.target"), notification.TypeEmail)
			go notificationSvc.NotifyAll(statuses, viper.GetStringSlice("notification.telegram.target"), notification.TypeTelegram)
		}
		c.JSON(http.StatusOK, context())
	})

	if err := r.Run(serverPort); err != nil {
		return 1
	}

	return 0
}
