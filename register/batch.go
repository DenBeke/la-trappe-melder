package register

import (
	"fmt"

	"gorm.io/gorm"
)

// Batch is just a stupid struct containing the La Trappe quadrupel Oak Aged batch number
type Batch struct {
	gorm.Model
	Batch uint `gorm:"primaryKey;autoIncrement:false"`
}

// AddBatch saves a newly detected batch number to the database
func (r *Register) AddBatch(batch uint) error {

	b := &Batch{Batch: batch}

	err := r.db.Create(b).Error
	if err != nil {
		return fmt.Errorf("couldn't insert batch into db: %w", err)
	}

	return nil

}

// GetBatches returns all batches from the database
func (r *Register) GetBatches() ([]*Batch, error) {

	batches := []*Batch{}

	err := r.db.Order("batch desc").Find(&batches).Error
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve batches from db: %w", err)
	}

	return batches, nil

}

// BatchExists checks whether a batch is known in our database
func (r *Register) BatchExists(batch uint) (bool, error) {

	batches := []*Batch{}

	err := r.db.Where(&Batch{Batch: batch}).Find(&batches).Error
	if err != nil {
		return false, fmt.Errorf("couldn't retrieve batches from db: %w", err)
	}

	return len(batches) == 1, nil
}
