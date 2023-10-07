package repository

import (
	"context"
	"database/sql"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/internal/repository/model"
	"project-adhyaksa/services/event/internal/repository/queries"
	"time"

	"github.com/rocketlaunchr/dbq"
	"go.uber.org/zap"
)

type eventRepository struct {
	db     *sql.DB
	config *config.Config
}

func NewEventRepository(db *sql.DB, config *config.Config) repository.EventRepository {
	return &eventRepository{db: db, config: config}
}

func (r *eventRepository) Create(event entity.Event, ctx context.Context) error {

	var eventModel model.Event

	eventModel.New(event)
	eventModel.CreatedAt = time.Now()

	duration, err := time.ParseDuration(r.config.CustomTime)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := dbq.INSERT(eventModel.GetTableName(),
		queries.RegisterEventStatment,
		1,
	)

	_, err = dbq.E(ctx, r.db, stmt, nil, queries.RegisterEventArgument(&eventModel))
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	return nil
}
