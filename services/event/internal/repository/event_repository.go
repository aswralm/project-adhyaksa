package repository

import (
	"context"
	"database/sql"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/internal/customerror"
	"project-adhyaksa/services/event/internal/repository/model"
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
		[]string{"id", "admin_id", "branch_id", "name", "start_time", "end_time", "location", "description", "created_at", "updated_at", "deleted_at"},
		1,
	)

	args := []interface{}{
		eventModel.ID,
		eventModel.AdminID,
		eventModel.BranchID,
		eventModel.Name,
		eventModel.StartTime,
		eventModel.EndTime,
		eventModel.Location,
		eventModel.Description,
		eventModel.CreatedAt,
		eventModel.UpdatedAt,
		eventModel.DeletedAt,
	}

	result, err := dbq.E(ctx, r.db, stmt, nil, args)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	if rowAffected == 0 {
		return &customerror.Err{
			Code:   customerror.ERROR_NOT_FOUND,
			Errors: nil,
		}
	}

	return nil
}
