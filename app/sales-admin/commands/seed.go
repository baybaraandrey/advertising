package commands

import (
	"fmt"

	"github.com/baybaraandrey/advertising/business/data/schema"
	"github.com/baybaraandrey/advertising/foundation/database"
	"github.com/pkg/errors"
)

// Seed seeds database.
func Seed(cfg database.Config) error {
	db, err := database.Open(cfg)
	if err != nil {
		return errors.Wrap(err, "connect database")
	}
	defer db.Close()

	if err := schema.Seed(db); err != nil {
		return errors.Wrap(err, "seeds database")
	}

	fmt.Println("seeds complete")
	return nil
}
