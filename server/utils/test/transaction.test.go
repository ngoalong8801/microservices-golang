package utils

import (
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/dal/model"
	"gorm.io/gorm"
)

func TestTransaction() {

	err := dal.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.User{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&model.Product{ProductID: 123}).Error; err != nil {
			return err
		}
		tx.Rollback()
		return nil
	})

	if err != nil {
		return
	}

}
