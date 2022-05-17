package models

import (
	"fmt"
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

func (tdo *TokenDataOperations) BulkInsertTokenData(tokenDataList []TokenData) error {
	err := tdo.DB.Transaction(func(tx *gorm.DB) error {
		tx.CreateInBatches(tokenDataList, 2000)
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

func (tdo *TokenDataOperations) TruncateTokenData() error {
	if err := tdo.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", "token_data")); err.Error != nil {
		return err.Error
	}
	return nil
}
