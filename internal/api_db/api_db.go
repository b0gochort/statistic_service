package api_db

import "github.com/restream/reindexer/v3"

type StatisticAPI interface {
}

type ApiDB struct {
	StatisticAPI
}

func NewAPIDB(db *reindexer.Reindexer) *ApiDB {
	return &ApiDB{
		StatisticAPI: NewChatApi(db),
	}
}
