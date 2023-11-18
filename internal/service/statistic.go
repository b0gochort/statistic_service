package service

import "github.com/b0gochort/statistic_service/internal/api_db"

type ChatServiceImpl struct {
	statisticAPI api_db.StatisticAPI
}

func NewChatService(statisticAPI api_db.StatisticAPI) *ChatServiceImpl {
	return &ChatServiceImpl{
		statisticAPI: statisticAPI,
	}
}
