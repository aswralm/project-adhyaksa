package repository

import (
	"context"
	"database/sql"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
	queryfilter "project-adhyaksa/services/event/domain/query_filter"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/internal/customerror"
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
		result, err := eventModel.MapEventEntityList(eventModels)
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

func (r *eventRepository) GetByID(id string, ctx context.Context) (*entity.Event, error) {
	var (
		eventModel model.Event
	)
	duration, err := time.ParseDuration(r.config.CustomTime)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	// depreciated because i have error and i not find the solving yet

	// stmt := fmt.Sprintf("SELECT * FROM %s LEFT JOIN %s ON %s.id = %s.branch_id WHERE %s.id = ?",
	// 	eventModel.GetTableName(),
	// 	branch.GetTableName(),
	// 	branch.GetTableName(),
	// 	eventModel.GetTableName(),
	// 	eventModel.GetTableName())
	// opts := &dbq.Options{SingleResult: true, ConcreteStruct: eventModel, DecoderConfig: dbq.StdTimeConversionConfig()}
	// data, err := dbq.Q(ctx, r.db, stmt, opts, id)
	// if err != nil {
	// 	zap.L().Error(err.Error())
	// 	return nil, err
	// }

	// event := data.(*model.Event)
	// if event.ID == "" {
	// 	return nil, &customerror.Err{
	// 		Code:   customerror.ERROR_NOT_FOUND,
	// 		Errors: nil,
	// 	}
	// }
	// result, err := event.EntitySingleMapping()
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil

	//change using gorm
	err = r.gormDB.
		WithContext(ctx).
		Table(eventModel.GetTableName()).
		Preload("Branch").
		Where("id = ?", id).
		First(&eventModel).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error(err.Error())
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, &customerror.Err{
			Code:   customerror.ERROR_NOT_FOUND,
			Errors: nil,
		}
	}

	result, err := model.MapEventEntity(&eventModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}
