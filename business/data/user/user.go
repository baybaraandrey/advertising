// Package user contains user related CRUD functionality.
package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/baybaraandrey/advertising/foundation/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

// User manages the set of API's for user access.
type User struct {
	log *log.Logger
	db  *sqlx.DB
}

// New constructs a User for api access.
func New(log *log.Logger, db *sqlx.DB) User {
	return User{
		log: log,
		db:  db,
	}
}

// Create inserts a new user into the database.
func (u User) Create(ctx context.Context, traceID string, nu NewUser, now time.Time) (Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "internal.data.user.create")
	defer span.End()

	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return Info{}, errors.Wrap(err, "generating password hash")
	}

	usr := Info{
		ID:           uuid.New().String(),
		Name:         nu.Name,
		Email:        nu.Email,
		Phone:        nu.Phone,
		PasswordHash: hash,
		Roles:        nu.Roles,
		Created:      now.UTC(),
		Updated:      now.UTC(),
	}

	const q = `INSERT INTO users
		(uuid, name, email, phone, password_hash, roles, created, updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	u.log.Printf("%s : %s : query : %s", traceID, "user.Create",
		database.Log(q, usr.ID, usr.Name, usr.Email, usr.Phone, usr.PasswordHash, usr.Roles, usr.Created, usr.Updated),
	)

	if _, err = u.db.ExecContext(ctx, q, usr.ID, usr.Name, usr.Email, usr.Phone, usr.PasswordHash, usr.Roles, usr.Created, usr.Updated); err != nil {
		return Info{}, errors.Wrap(err, "inserting user")
	}

	return usr, nil
}

// Query retrieves a list of existing users from the database.
func (u User) Query(ctx context.Context, traceID string) ([]Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.query")
	defer span.End()

	const q = `SELECT * FROM users`

	log.Printf("%s : %s : query : %s", traceID, "user.Query",
		database.Log(q),
	)

	users := []Info{}
	if err := u.db.SelectContext(ctx, &users, q); err != nil {
		return nil, errors.Wrap(err, "selecting users")
	}

	return users, nil
}

// QueryByID gets the specified user from the database.
func (u User) QueryByID(ctx context.Context, traceID string, userID string) (Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.querybyid")
	defer span.End()

	if _, err := uuid.Parse(userID); err != nil {
		return Info{}, ErrInvalidID
	}

	const q = `SELECT * FROM users WHERE uuid = $1`

	u.log.Printf("%s : %s : query : %s", traceID, "user.QueryByID",
		database.Log(q, userID),
	)

	var usr Info
	if err := u.db.GetContext(ctx, &usr, q, userID); err != nil {
		if err == sql.ErrNoRows {
			return Info{}, ErrNotFound
		}
		return Info{}, errors.Wrapf(err, "selecting user %q", userID)
	}

	return usr, nil
}
