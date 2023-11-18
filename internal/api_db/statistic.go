package api_db

import "github.com/restream/reindexer/v3"

type StatisticAPIImpl struct {
	db *reindexer.Reindexer
}

func NewChatApi(db *reindexer.Reindexer) *StatisticAPIImpl {
	return &StatisticAPIImpl{
		db: db,
	}
}
