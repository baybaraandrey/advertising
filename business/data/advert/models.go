package advert

import (
	"time"

	"github.com/baybaraandrey/advertising/business/data/category"
	"github.com/baybaraandrey/advertising/business/data/user"
)

// Info represent an individual advert.
type Info struct {
	ID          string `db:"uuid" json:"id"`
	UserID      string `db:"user_uuid" json:"user_uuid"`
	CategoryID  string `db:"category_uuid" json:"category_uuid"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Location    string `db:"location" json:"location"`
	Price       int    `db:"price" json:"price"`
	IsActive    bool   `db:"is_active" json:"is_active"`

	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

// AdvertInfo represent an individual advert.
type AdvertInfo struct {
	ID          string `db:"uuid" json:"id"`
	UserID      string `db:"user_uuid" json:"user_uuid"`
	CategoryID  string `db:"category_uuid" json:"category_uuid"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Location    string `db:"location" json:"location"`
	Price       int    `db:"price" json:"price"`
	IsActive    bool   `db:"is_active" json:"is_active"`

	Created      time.Time             `db:"created" json:"created"`
	Updated      time.Time             `db:"updated" json:"updated"`
	UserInfo     user.UserInfo         `db:"user" json:"user"`
	CategoryInfo category.CategoryInfo `db:"category" json:"category"`
}

// NewAdvert represent an individual advert.
type NewAdvert struct {
	UserID      string `db:"user_uuid" json:"user_uuid"`
	CategoryID  string `db:"category_uuid" json:"category_uuid"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Location    string `db:"location" json:"location"`
	Price       int    `db:"price" json:"price"`
}
