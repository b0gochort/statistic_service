package service

import "github.com/b0gochort/statistic_service/internal/api_db"

type StatisticService interface {
}

type Service struct {
	StatisticService
}

func NewService(ApiDB *api_db.ApiDB) *Service {
	return &Service{
		StatisticService: NewChatService(ApiDB.StatisticAPI),
	}
}
