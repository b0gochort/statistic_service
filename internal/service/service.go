package service

import (
	"github.com/b0gochort/statistic_service/internal/api_db"
	"github.com/b0gochort/statistic_service/internal/model"
)

type StatisticService interface {
	GetOnline() (model.OnlineLens, error)
	SetOnline(userId int64, auth bool) error
	GetStatisticByHour() ([]int, error)
	GetCategoryRatio() (map[string]int, error)
}

type Service struct {
	StatisticService
}

func NewService(ApiDB *api_db.ApiDB) *Service {
	return &Service{
		StatisticService: NewStatisticService(ApiDB.StatisticAPI),
	}
}
