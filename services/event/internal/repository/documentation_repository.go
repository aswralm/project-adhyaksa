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
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// Commit transaction
	return tx.Commit()
}

func (r *documentationRepository) Create(photo entity.Photo, ctx context.Context) error {
	var (
		documentationModel model.Documentation
		photoModel         model.Photo
	)

	documentationModel.New(*photo.GetDocumentation())
	documentationModel.CreatedAt = time.Now()

	photoModel.New(photo)
	photoModel.CreatedAt = time.Now()

	return r.transaction(func(tx *sql.Tx) error {
		//create documentation first
		stmtDocumentation := dbq.INSERT(documentationModel.GetTableName(),
			[]string{"id", "admin_id", "branch_id", "name", "date", "participant", "location", "description", "created_at"},
			1,
		)

		argsDocumentation := []interface{}{
			documentationModel.ID,
			documentationModel.AdminID,
			documentationModel.BranchID,
			documentationModel.Name,
			documentationModel.Date,
			documentationModel.Participant,
			documentationModel.Location,
			documentationModel.Description,
			documentationModel.CreatedAt,
		}

		_, err := dbq.E(ctx, r.db, stmtDocumentation, nil, argsDocumentation)
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
			photoModel.ID,
			photoModel.DocumentationID,
			photoModel.PublicID,
			photoModel.URL,
			photoModel.Name,
			photoModel.CreatedAt,
		}

		_, err = dbq.E(ctx, r.db, stmtPhoto, nil, argsPhoto)
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		return nil
	})
}
