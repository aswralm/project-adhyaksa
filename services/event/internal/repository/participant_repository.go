package repository

import (
	"context"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/services/event/domain/entity"
	"project-adhyaksa/services/event/domain/repository"
	"project-adhyaksa/services/event/internal/repository/model"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type participantRepository struct {
	gormDB *gorm.DB
	config *config.Config
}

func NewParticipantRepository(config *config.Config) repository.ParticipantRepository {
	return &participantRepository{gormDB: config.GormDB, config: config}
}

func (r *participantRepository) Create(participant *entity.Participant, ctx context.Context) error {
	duration, err := time.ParseDuration(r.config.CustomTime)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	var participantModel model.Participant
	data := participantModel.New(participant)
	if err := r.gormDB.WithContext(ctx).Table(participantModel.GetTableName()).Create(data).Error; err != nil {
		zap.L().Error(err.Error())
		return err
	}

	return nil
}
