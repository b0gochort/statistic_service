package api_db

import (
	"github.com/b0gochort/statistic_service/internal/model"
	"github.com/restream/reindexer/v3"
)

type StatisticAPI interface {
	GetOnline() (model.OnlineLens, error)
	SetOnline(user model.OnlineItem) error
	GetAllChats() ([]model.NewChatItem, error)
}

type ApiDB struct {
	StatisticAPI
}

func NewAPIDB(db *reindexer.Reindexer) *ApiDB {
	return &ApiDB{
		StatisticAPI: NewChatApi(db),
	}
}
