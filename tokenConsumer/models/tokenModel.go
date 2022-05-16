package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TokenData struct {
	ID        string    `db:"id" json:"id" gorm:"id"`
	Token     string    `db:"token" json:"token" gorm:"id"`
	Count     int64     `db:"count" json:"count" gorm:"count"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type TokenDataOperations struct {
	DB *gorm.DB
}

func (tdo *TokenDataOperations) InsertTokenData(tokenData *TokenData) error {
	err := tdo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Take(&TokenData{Token: tokenData.Token}).Scan(&tokenData); err.Error != nil {
			if err.Error.Error() == "No Record Found" {
				tokenData.ID = uuid.New().String()
				tokenData.CreatedAt = time.Now()
				tokenData.UpdatedAt = time.Now()
				tx.Create(&tokenData)
				return nil
			} else {
				return err.Error
			}
		}
		if err := tx.Where("id = ?", tokenData.ID).Updates(&TokenData{Count: tokenData.Count + 1, UpdatedAt: time.Now()}); err.Error != nil {
			return err.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (tdo *TokenDataOperations) ListTokenData() ([]TokenData, error) {
	var tokenDataRecords []TokenData
	if err := tdo.DB.Find(&tokenDataRecords); err.Error != nil {
		return nil, err.Error
	}
	return tokenDataRecords, nil
}
