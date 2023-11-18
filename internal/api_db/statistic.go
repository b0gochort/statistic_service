package api_db

import (
	"fmt"
	"github.com/b0gochort/statistic_service/internal/model"
	"github.com/restream/reindexer/v3"
)

type StatisticAPIImpl struct {
	db *reindexer.Reindexer
}

func NewChatApi(db *reindexer.Reindexer) *StatisticAPIImpl {
	return &StatisticAPIImpl{
		db: db,
	}
}

func (a *StatisticAPIImpl) SetOnline(user model.OnlineItem) error {
	err := a.db.OpenNamespace("online", reindexer.DefaultNamespaceOptions(), model.OnlineItem{})
	if err != nil {
		return fmt.Errorf("chatApi.SetOnline.OpenNamespace: %v", err)
	}

	ok, err := a.db.Insert("online", &user)
	if err != nil {
		return fmt.Errorf("chatApi.db.Insert: %v", err)
	}

	if ok == 0 {
		return fmt.Errorf("nil insert")
	}

	return nil
}

func (a *StatisticAPIImpl) GetOnline() (model.OnlineLens, error) {

	err := a.db.OpenNamespace("online", reindexer.DefaultNamespaceOptions(), model.OnlineItem{})
	if err != nil {
		return model.OnlineLens{}, fmt.Errorf("chatApi.GetOnline.OpenNamespace: %v", err)
	}

	query1 := a.db.Query("online").Where("Auth", reindexer.EQ, true).ReqTotal()
	query2 := a.db.Query("online").Where("Auth", reindexer.EQ, false).ReqTotal()
	lenAuth := query1.Exec().Count()
	lenUnAuth := query2.Exec().Count()

	lens := model.OnlineLens{
		LenAuth:   lenAuth,
		LenUnAuth: lenUnAuth,
	}

	return lens, nil
}

func (a *StatisticAPIImpl) GetAllChats() ([]model.NewChatItem, error) {
	err := a.db.OpenNamespace("support_chat", reindexer.DefaultNamespaceOptions(), model.NewChatItem{})
	if err != nil {
		return nil, fmt.Errorf("StatisticAPIImpl.GetAllChats.OpenNamespace: %v", err)
	}
	query1 := a.db.Query("support_chat")
	var res []model.NewChatItem

	iter := query1.Exec()
	if iter.Error() != nil {
		return nil, fmt.Errorf("StatisticAPIImpl.GetAllChats.Exec: %v", err)
	}

	for iter.Next() {
		elem := iter.Object().(*model.NewChatItem)
		res = append(res, model.NewChatItem{
			ID:            elem.ID,
			Time:          elem.Time,
			Message:       elem.Message,
			LastHostStaff: elem.LastHostStaff,
			UID:           elem.UID,
			IP:            elem.IP,
		})
	}

	return res, nil
}
