package category

import "time"

// CategoryInfo represent an individual category.
type CategoryInfo struct {
	ID      string    `db:"uuid" json:"id"`
	Name    string    `db:"name" json:"name"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

//  NewCategory ...
type NewCategory struct {
	Name string `json:"name" validate:"required"`
}

// UpdateCategory ...
type UpdateCategory struct {
	Name *string `json:"name"`
}
