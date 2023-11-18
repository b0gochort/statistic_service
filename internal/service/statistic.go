package service

import (
	"fmt"
	"github.com/b0gochort/statistic_service/internal/api_db"
	"github.com/b0gochort/statistic_service/internal/model"
	"time"
)

type StatisticServiceImpl struct {
	statisticAPI api_db.StatisticAPI
}

func NewStatisticService(statisticAPI api_db.StatisticAPI) *StatisticServiceImpl {
	return &StatisticServiceImpl{
		statisticAPI: statisticAPI,
	}
}

func (s *StatisticServiceImpl) SetOnline(userId int64, auth bool) error {
	req := model.OnlineItem{
		UserId: userId,
		Auth:   auth,
		Time:   time.Now().Unix(),
	}

	if err := s.statisticAPI.SetOnline(req); err != nil {
		return fmt.Errorf("service.SetOnline.%v", err)
	}

	return nil
}

func (s *StatisticServiceImpl) GetOnline() (model.OnlineLens, error) {
	res, err := s.statisticAPI.GetOnline()
	if err != nil {
		return model.OnlineLens{}, fmt.Errorf("service.GetOnline.%v")
	}

	return res, nil
}

func (s *StatisticServiceImpl) GetStatisticByHour() ([]int, error) {
	res, err := s.statisticAPI.GetAllChats()
	if err != nil {
		return nil, fmt.Errorf("service.GetCategoryRatio.GetAllChats.%v", err)
	}

	hourlyCounts := make([]int, 24)

	for _, chat := range res {
		hour := time.Unix(chat.Time, 0).Hour()
		hourlyCounts[hour]++
	}

	return hourlyCounts, nil
}

func (s *StatisticServiceImpl) GetCategoryRatio() (map[string]int, error) {
	res, err := s.statisticAPI.GetAllChats()
	if err != nil {
		return nil, fmt.Errorf("service.GetCategoryRatio.GetAllChats.%v", err)
	}

	categoryCounts := make(map[string]int)

	for _, chat := range res {
		categoryCounts[chat.Category]++
	}

	return categoryCounts, nil

}
