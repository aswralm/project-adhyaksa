package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/pkg/pagination"
	"project-adhyaksa/services/event/domain/entity"
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

type documentationRepository struct {
	db     *sql.DB
	gormDB *gorm.DB
	config *config.Config
}

func NewDocumentationRepository(config *config.Config) repository.DocumentationRepository {
	return &documentationRepository{db: config.Db, gormDB: config.GormDB, config: config}
}

func (r *documentationRepository) transaction(fn func(tx *sql.Tx) error) error {

	// create new db transaction
	tx, err := r.db.Begin()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	// If there are any panic it will rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Run the transaction
	err = fn(tx)
	if err != nil {
		// Rollback if we have error
		if rbErr := tx.Rollback(); rbErr != nil {
			zap.L().Error(err.Error())
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// Commit transaction
	return tx.Commit()
}

func (r *documentationRepository) Create(documentation entity.Documentation, photo entity.Photo, ctx context.Context) error {
	var (
		documentationModel model.Documentation
		photoModel         model.Photo
	)

	duration, err := time.ParseDuration(r.config.CustomTime)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	documentationmodel := documentationModel.New(documentation)
	documentationmodel.CreatedAt = time.Now()

	photomodel := photoModel.New(photo)
	photomodel.CreatedAt = time.Now()

	return r.transaction(func(tx *sql.Tx) error {
		//create documentation first
		stmtDocumentation := dbq.INSERT(documentationModel.GetTableName(), queries.RegisterDocumentationStatment, 1)

		_, err := dbq.E(ctx, tx, stmtDocumentation, nil, queries.RegisterDocumentationArgument(documentationmodel))
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}

		//create photo
		stmtPhoto := dbq.INSERT(photoModel.GetTableName(), queries.RegisterPhotoStatment, 1)

		_, err = dbq.E(ctx, tx, stmtPhoto, nil, queries.RegisterPhotoArgument(photomodel))
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		return nil
	})
}

func (r *documentationRepository) GetListPaginated(pagin *pagination.Paginator, ctx context.Context) ([]*entity.Documentation, error) {
	var (
		documentationModel model.Documentation
		documentations     []*entity.Documentation
		concurrentCount    = 2
		errChan            = make(chan error, concurrentCount)
		wg                 sync.WaitGroup
	)

	wg.Add(concurrentCount)

	scopeFilter := queries.GetListDocumentationFilterGORM(pagin)

	go func() {
		defer wg.Done()
		var documentationModels []model.Documentation

		err := r.gormDB.
			WithContext(ctx).
			Scopes(scopeFilter).
			Find(&documentationModels).
			Error
		if err != nil {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		result, err := documentationModel.MapDocumentationEntityList(documentationModels)
		if err != nil {
			errChan <- err
			return
		}

		documentations = result
	}()

	go func() {
		defer wg.Done()

		var count int64
		if err := queries.
			GetListDocumentationCountGORM(r.gormDB).
			WithContext(ctx).
			Count(&count).
			Error; err != nil {
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
			return nil, err
		}
	}

	return documentations, nil
}

func (r *documentationRepository) GetByID(id string, ctx context.Context) (*entity.Documentation, error) {
	var (
		documentationModel model.Documentation
	)
	duration, err := time.ParseDuration(r.config.CustomTime)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	err = r.gormDB.
		WithContext(ctx).
		Table(documentationModel.GetTableName()).
		Preload("Branch").
		Preload("Photo").
		Where("id = ?", id).
		First(&documentationModel).
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

	result, err := model.MapDocumentationEntity(&documentationModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}
