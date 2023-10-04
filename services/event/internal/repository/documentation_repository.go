package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/internal/repository/model"
	"time"

	"github.com/rocketlaunchr/dbq"
	"go.uber.org/zap"
)

type documentationRepository struct {
	db     *sql.DB
	config *config.Config
}

func NewDocumentationRepository(db *sql.DB, config *config.Config) repository.DocumentationRepository {
	return &documentationRepository{db: db, config: config}
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

	documentationmodel := documentationModel.New(documentation)
	documentationmodel.CreatedAt = time.Now()

	photomodel := photoModel.New(photo)
	photomodel.CreatedAt = time.Now()

	return r.transaction(func(tx *sql.Tx) error {
		//create documentation first
		stmtDocumentation := dbq.INSERT(documentationModel.GetTableName(),
			[]string{"id", "admin_id", "branch_id", "name", "date", "participant", "location", "description", "created_at"},
			1,
		)

		argsDocumentation := []interface{}{
			documentationmodel.ID,
			documentationmodel.AdminID,
			documentationmodel.BranchID,
			documentationmodel.Name,
			documentationmodel.Date,
			documentationmodel.Participant,
			documentationmodel.Location,
			documentationmodel.Description,
			documentationmodel.CreatedAt,
		}

		_, err := dbq.E(ctx, tx, stmtDocumentation, nil, argsDocumentation)
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		//create photo
		stmtPhoto := dbq.INSERT(photoModel.GetTableName(),
			[]string{"id", "documentation_id", "public_id", "url", "name", "created_at"},
			1,
		)
		argsPhoto := []interface{}{
			photomodel.ID,
			photomodel.DocumentationID,
			photomodel.PublicID,
			photomodel.URL,
			photomodel.Name,
			photomodel.CreatedAt,
		}

		_, err = dbq.E(ctx, tx, stmtPhoto, nil, argsPhoto)
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		return nil
	})
}
