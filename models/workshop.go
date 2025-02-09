package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Workshop struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}

// DeleteWorkshop - Elimina un workshop por su ID
func DeleteWorkshop(db *gorm.DB, id uuid.UUID) error {
	result := db.Delete(&Workshop{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("Workshop not found")
	}
	return nil
}
