package register

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Subscription contains an email subscription
type Subscription struct {

	// gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// actual fields
	UUID      string `gorm:"primaryKey;not null:default:null;autoIncrement:false"`
	Name      string `gorm:"not null:default:null"`
	Email     string `gorm:"unique;not null;default:null"`
	Confirmed bool   `gorm:"default:false"`
}

// Subscribe creates a new subscription with the given name and email
func (r *Register) Subscribe(name string, email string) (*Subscription, error) {

	s := &Subscription{
		UUID:  uuid.New().String(),
		Name:  name,
		Email: email,
	}

	err := r.db.Create(s).Error
	if err != nil {
		return nil, fmt.Errorf("couldn't insert subscription into db: %w", err)
	}

	return s, nil
}

// UnSubscribe removes the email from the register
func (r *Register) UnSubscribe(email string) error {

	s := &Subscription{}

	err := r.db.Where(&Subscription{Email: email}).Take(&s).Error
	if err != nil {
		return fmt.Errorf("couldn't find email subscription in db: %w", err)
	}

	err = r.db.Delete(s).Error
	if err != nil {
		return fmt.Errorf("couldn't delete email subscription from db: %w", err)
	}

	return nil
}

// GetAllSubscriptions returns all subscriptions
func (r *Register) GetAllSubscriptions() ([]*Subscription, error) {
	s := []*Subscription{}

	err := r.db.Find(&s).Error
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve subscriptions from db: %w", err)
	}

	return s, nil
}

// ConfirmSubscription confirms the subscription that matches the given id
func (r *Register) ConfirmSubscription(uuid string) (*Subscription, error) {

	s := &Subscription{}

	err := r.db.Where(&Subscription{UUID: uuid}).Take(&s).Error
	if err != nil {
		return nil, fmt.Errorf("couldn't find email subscription in db: %w", err)
	}

	s.Confirmed = true

	err = r.db.Save(s).Error
	if err != nil {
		return nil, fmt.Errorf("couldn't confirm email subscription in db: %w", err)
	}

	return s, nil
}
