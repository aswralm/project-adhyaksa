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
	"sync"
	"time"

	"github.com/rocketlaunchr/dbq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type eventRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
	config *config.Config
}

func NewEventRepository(config *config.Config) repository.EventRepository {
	return &eventRepository{db: config.Db, gormDB: config.GormDB, config: config}
}

func (r *eventRepository) Create(event entity.Event, ctx context.Context) error {
	duration, err := time.ParseDuration(r.config.CustomTime)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	var eventModel model.Event

	eventModel.New(event)
	eventModel.CreatedAt = time.Now()

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

func (r *eventRepository) GetListPaginated(
	pagin *pagination.Paginator,
	filter *queryfilter.GetEventQueryFilter,
) ([]*entity.Event, error) {
	var (
		eventModel      model.Event
		events          []*entity.Event
		concurrentCount = 2
		errChan         = make(chan error, concurrentCount)
		wg              sync.WaitGroup
	)

	wg.Add(concurrentCount)

	scopeFilter := queries.GetListEventFilterGORM(pagin, filter)

	go func() {
		defer wg.Done()
		var eventModels []model.Event

		err := r.gormDB.
			Scopes(scopeFilter).
			Find(&eventModels).
			Error
		if err != nil {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		result, err := eventModel.EntityMapping(eventModels)
		if err != nil {
			errChan <- err
			return
		}

		events = result
	}()

	go func() {
		defer wg.Done()

		var count int64
		if err := queries.GetListEventCountGORM(pagin, filter, r.gormDB).Count(&count).Error; err != nil {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		pagin.SetTotal(count)
	}()

	// Close the channel after both goroutines are done
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check if any error exists
	for i := 0; i < concurrentCount; i++ {
		if err := <-errChan; err != nil {
			return events, err
		}
	}

	return events, nil
}
