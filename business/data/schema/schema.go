package schema

import (
	"github.com/dimiro1/darwin"
	"github.com/jmoiron/sqlx"
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sqlx.DB) error {
	driver := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})
	d := darwin.New(driver, migrations, nil)
	return d.Migrate()
}

// migrations contains the queries needed to construct the database schema.
// Entries should never be removed from this slice once they have been ran in
// production.
//
// Using constants in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add users",
		Script: `
CREATE TABLE users (
	uuid          UUID,
	name          TEXT,
	email         TEXT UNIQUE NOT NULL,
	phone         TEXT UNIQUE NOT NULL,
	roles         TEXT[],
	password_hash TEXT,
	created TIMESTAMP,
	updated TIMESTAMP,
	PRIMARY KEY (uuid)
);`},
	{
		Version:     2,
		Description: "Add advert category",
		Script: `
CREATE TABLE categories (
	uuid		UUID,
	name		TEXT UNIQUE NOT NULL,
	created TIMESTAMP,
	updated TIMESTAMP,
	PRIMARY KEY (uuid)
);`},
	{
		Version:     3,
		Description: "Add advert",
		Script: `
CREATE TABLE adverts (
	uuid		  UUID,
	user_uuid     UUID REFERENCES users(uuid) NOT NULL,
	category_uuid UUID REFERENCES categories(uuid) NOT NULL,
	title         VARCHAR(256) NOT NULL,
	description   TEXT NOT NULL,
	location      TEXT NOT NULL DEFAULT '',
	price         INTEGER NOT NULL,
	created TIMESTAMP,
	updated TIMESTAMP,
	PRIMARY KEY (uuid)
);`},
}
