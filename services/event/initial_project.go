package init_project

import (
	"log"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/pkg/database"
	"project-adhyaksa/pkg/logger"

	"go.uber.org/zap"
)

func NewProject() {
	path := `./services/event/log/error.log`
	l, err := logger.NewZapFileLogger(path)
	if err != nil {
		log.Fatal("logger error ", err)
	}
	zap.ReplaceGlobals(zap.Must(l, err))

	config, err := config.NewConfig("./services/event/")
	if err != nil {
		log.Fatal("config error ", err)
	}
	database.ConnectMYSQL("mysql", config)
}
