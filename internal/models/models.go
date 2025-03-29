package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"uniqueIndex; not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"` //Exclude from JSON response
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`

	Blogs []Blog `gorm:"foreignKey:UserID" json:"blogs"`

	BookMarks []Blog    `gorm:"many2many:book_marks;joinForeignKey:UserID;joinReferences:BlogID" json:"bookmarks"`
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments"`
}

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	Claps     int       `json:"claps"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`

	BookMarkedBy []User    `gorm:"many2many:book_marks;joinForeignKey:BlogID;joinReferences:UserID" json:"bookmarked_by"`
	Comments     []Comment `gorm:"foreignKey:BlogID" json:"comments"`
}

type BookMark struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	BlogID    uuid.UUID `json:"blog_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	Blog      Blog      `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE;" json:"blog"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	BlogID    uuid.UUID `json:"blog_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	Blog      Blog      `gorm:"foreignKey:BlogID;constraint:OnDelete:CASCADE;" json:"blog"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

func (bm *BookMark) BeforeCreate(tx *gorm.DB) (err error) {
	if bm.ID == uuid.Nil {
		bm.ID = uuid.New()
	}
	return
}
