package service

import (
	"sync"

	"github.com/b3log/solo.go/model"
)

var Preference = &preferenceService{
	mutex: &sync.Mutex{},
}

type preferenceService struct {
	mutex *sync.Mutex
}

func (srv *preferenceService) GetPreference(preferenceName string, blogID uint) *model.Setting {
	ret := &model.Setting{}
	if nil != db.Where("name = ? AND blog_id = ?", preferenceName, blogID).Find(ret).Error {
		return nil
	}

	return ret
}

func (srv *preferenceService) GetPreferences(blogID uint, preferenceNames ...string) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	settings := []*model.Setting{}
	if nil != db.Where("name IN (?) AND blog_id = ?", preferenceNames, blogID).Find(&settings).Error {
		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}
