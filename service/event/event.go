package event

import (
	"events_go/models"
	"time"
)

type Event struct {
	Id           uint
	Name         string
	Type_id      uint
	Location_id  uint
	Capacity     uint
	Description  string
	Performer_id uint
	Duration     time.Duration
	StartTime    time.Time
	EndTime      time.Time
}

// Constructor
func NewEvent(id uint, name string, type_id uint, location_id uint, capacity uint, description string, performer_id uint, duration time.Duration, startTime time.Time, endTime time.Time) *Event {
	return &Event{Id: id, Name: name, Type_id: type_id, Location_id: location_id, Capacity: capacity, Description: description, Performer_id: performer_id, Duration: duration, StartTime: startTime, EndTime: endTime}
}

func (e *Event) Add() error {
	event := map[string]interface{}{
		"id":           e.Id,
		"name":         e.Name,
		"type_id":      e.Type_id,
		"location_id":  e.Location_id,
		"description":  e.Capacity,
		"performer_id": e.Performer_id,
		"duration":     e.Duration,
		"startTime":    e.StartTime,
		"endTime":      e.EndTime,
	}

	if err := models.AddEvent(event); err != nil {
		return err
	}

	return nil
}

func (e *Event) Edit() error {
	return models.EditEvent(e.Id, map[string]interface{}{
		"id":           e.Id,
		"name":         e.Name,
		"type_id":      e.Type_id,
		"location_id":  e.Location_id,
		"description":  e.Capacity,
		"performer_id": e.Performer_id,
		"duration":     e.Duration,
		"startTime":    e.StartTime,
		"endTime":      e.EndTime,
	})
}

func (e *Event) Get() (*models.Event, error) {
	//	var cacheArticle *models.Event

	//cache := cache_service.Article{ID: a.ID}
	//key := cache.GetArticleKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info(err)
	//	} else {
	//		json.Unmarshal(data, &cacheArticle)
	//		return cacheArticle, nil
	//	}
	//}

	event, err := models.GetEvent(e.Id)
	if err != nil {
		return nil, err
	}
	//
	//gredis.Set(key, article, 3600)
	//return article, nil
}

func (e *Event) GetAll() ([]*models.Event, error) {
	var (
	//events, cacheEvents []*models.Event
	)
	//
	//cache := cache_service.Article{
	//	TagID: a.TagID,
	//	State: a.State,
	//
	//	PageNum:  a.PageNum,
	//	PageSize: a.PageSize,
	//}
	//key := cache.GetArticlesKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info(err)
	//	} else {
	//		json.Unmarshal(data, &cacheArticles)
	//		return cacheArticles, nil
	//	}
	//}
	//
	//events, err := models.GetEvents(e.PageNum, e.PageSize, e.getMaps())
	//if err != nil {
	//	return nil, err
	//}
	//return
}

//
//	gredis.Set(key, articles, 3600)
//	return articles, nil
//}

func (e *Event) Delete() error {
	return models.DeleteEvent(e.Id)
}

func (e *Event) ExistByID() (bool, error) {
	return models.ExistEventByID(e.Id)
}

func (e *Event) Count() (int, error) {
	return models.GetEventsTotal(e.getMaps())
}

func (e *Event) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	//maps["deleted_on"] = 0
	//if a.State != -1 {
	//	maps["state"] = a.State
	//}
	//if a.TagID != -1 {
	//	maps["tag_id"] = a.TagID
	//}

	return maps
}
