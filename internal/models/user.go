package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               string     `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name             string     `json:"name" gorm:"not null" validate:"required,min=1,max=50"`
	Email            string     `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	PasswordHash     string     `json:"-" gorm:"not null"`
	Admin            bool       `json:"admin" gorm:"default:false"`
	Activated        bool       `json:"activated" gorm:"default:false"`
	ActivationDigest string     `json:"-"`
	ActivatedAt      *time.Time `json:"activated_at"`
	ResetDigest      string     `json:"-"`
	ResetSentAt      *time.Time `json:"reset_sent_at"`
	RememberDigest   string     `json:"-"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// Associations
	Microposts           []Micropost    `json:"microposts,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ActiveRelationships  []Relationship `json:"-" gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	PassiveRelationships []Relationship `json:"-" gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) GravatarURL(size int) string {
	return GetGravatarURL(u.Email, size)
}

func (u *User) Following() ([]User, error) {
	// This would be implemented with proper database queries
	return []User{}, nil
}

func (u *User) Followers() ([]User, error) {
	// This would be implemented with proper database queries
	return []User{}, nil
}

func (u *User) Follow(other *User) error {
	// This would be implemented with proper database operations
	return nil
}

func (u *User) Unfollow(other *User) error {
	// This would be implemented with proper database operations
	return nil
}

func (u *User) IsFollowing(other *User) bool {
	// This would be implemented with proper database queries
	return false
}

func (u *User) Feed() ([]Micropost, error) {
	// This would return microposts from followed users + own microposts
	return []Micropost{}, nil
}
