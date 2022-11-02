package repository

import (
	"app-indihomesmart/infra/database"
	"app-indihomesmart/infra/logger"
)

func Save(model interface{}) interface{} {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, not save data %v", err)
	}
	return err
}

func Get(model interface{}, limit, offset int) interface{} {
	err := database.DB.Limit(limit).Offset(offset).Find(model).Error
	return err
}

func GetOne(model interface{}) error {
	err := database.DB.Last(model).Error
	return err
}

func Update(model, data interface{}) interface{} {
	err := database.DB.Find(model).Updates(data).Error
	return err
}

func Delete(model interface{}) interface{} {
	err := database.DB.Delete(model).Error
	return err
}
