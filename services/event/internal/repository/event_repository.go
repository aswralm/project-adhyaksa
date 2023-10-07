package repository

import (
	"context"
	"database/sql"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
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

func (r *eventRepository) GetListPaginated(ctx context.Context,
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]entity.Event, error) {
	var (
		event           model.Event
		branch          model.Branch
		events          []entity.Event
		concurrentCount = 2
		errChan         = make(chan error, concurrentCount)
	)
	//create scope filter
	query, argument := queries.GetListEventFilter(&event, &branch, pagin, filter)
	queryCount := queries.GetListEventCount(event)

	go func() {
		defer close(errChan)

		opts := &dbq.Options{SingleResult: false, ConcreteStruct: event, DecoderConfig: dbq.StdTimeConversionConfig()}
		data := dbq.MustQ(ctx, r.db, query, opts, argument)

		if err, ok := data.(error); ok {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		result, err := event.EntityMapping(data.([]*model.Event))
		zap.L().Error(err.Error())
		if err != nil {
			errChan <- err
			return
		}
		events = result
	}()

	go func() {
		defer close(errChan)

		var count int64
		opts := &dbq.Options{SingleResult: true, ConcreteStruct: count, DecoderConfig: dbq.StdTimeConversionConfig()}
		data := dbq.MustQ(ctx, r.db, queryCount, opts)

		if err, ok := data.(error); ok {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		pagin.SetTotal(data.(int64))
	}()

	// Check if any error exists
	for i := 0; i < concurrentCount; i++ {
		if err := <-errChan; err != nil {
			return events, err
		}
	}

	return events, nil
}
