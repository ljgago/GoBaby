package repository_domain

import (
	repository_adapters "GoBaby/cmd/web/domain/repository/adapters"
	db_config "GoBaby/cmd/web/domain/repository/config"
	"GoBaby/internal/models"
	"log/slog"
)

func InitializeBD() {
	_, err := db_config.InitializeDb()
	if err != nil {
		slog.Error(err.Message)

		// we want it to panic if the database is not initialized as it is crucial for the application
		panic(err.Err)
	}
}

func GetUserByUUID(uuid int) (models.User, *models.AppError) {
	return repository_adapters.GetUserByUUID(uuid)
}

func SetUser(user *models.User) *models.AppError {
	return repository_adapters.SetUser(user)
}

func AddLogByUUID(uuid int, log models.Log) *models.AppError {
	return repository_adapters.AddLogByUUID(uuid, log)
}

func AddMonitorLog(monitorLog models.MonitorLog) {
	err := repository_adapters.AddMonitorLog(monitorLog)
	if err != nil {
		slog.Error(err.Message)
	}
}
