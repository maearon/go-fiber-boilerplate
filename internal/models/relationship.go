package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Relationship struct {
	ID         string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FollowerID string    `json:"follower_id" gorm:"type:uuid;not null;index"`
	FollowedID string    `json:"followed_id" gorm:"type:uuid;not null;index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Associations
	Follower User `json:"follower,omitempty" gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Followed User `json:"followed,omitempty" gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}

func (r *Relationship) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}
	return nil
}
