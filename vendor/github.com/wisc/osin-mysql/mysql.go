// Package mysql is a osin storage implementation for mysql.
package mysql

import (
	"database/sql"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/go-errors/errors"
	"github.com/ory-am/common/pkg"
	"log"
	"time"
)

var schemas = []string{`CREATE TABLE IF NOT EXISTS client (
	id           varchar(255) NOT NULL PRIMARY KEY,
	secret 		 varchar(255) NOT NULL,
	extra 		 varchar(255) NOT NULL,
	redirect_uri varchar(255) NOT NULL
)`, `CREATE TABLE IF NOT EXISTS authorize (
	client       varchar(255) NOT NULL,
	code         varchar(255) NOT NULL PRIMARY KEY,
	expires_in   int(10) NOT NULL,
	scope        varchar(255) NOT NULL,
	redirect_uri varchar(255) NOT NULL,
	state        varchar(255) NOT NULL,
	extra 		 varchar(255) NOT NULL,
	created_at   timestamp NOT NULL
)`, `CREATE TABLE IF NOT EXISTS access (
	client        varchar(255) NOT NULL,
	authorize     varchar(255) NOT NULL,
	previous      varchar(255) NOT NULL,
	access_token  varchar(255) NOT NULL PRIMARY KEY,
	refresh_token varchar(255) NOT NULL,
	expires_in    int(10) NOT NULL,
	scope         varchar(255) NOT NULL,
	redirect_uri  varchar(255) NOT NULL,
	extra 		  varchar(255) NOT NULL,
	created_at    timestamp NOT NULL
)`, `CREATE TABLE IF NOT EXISTS refresh (
	token         varchar(255) NOT NULL PRIMARY KEY,
	access        varchar(255) NOT NULL
)`}

// Storage implements interface "github.com/RangelReale/osin".Storage and interface "github.com/felipeweb/osin-mysql/storage".Storage
type Storage struct {
	db *sql.DB
}

// New returns a new mysql storage instance.
func New(db *sql.DB) *Storage {
	return &Storage{db}
}

// CreateSchemas creates the schemata, if they do not exist yet in the database. Returns an error if something went wrong.
func (s *Storage) CreateSchemas() error {
	for k, schema := range schemas {
		if _, err := s.db.Exec(schema); err != nil {
			log.Printf("Error creating schema %d: %s", k, schema)
			return err
		}
	}
	return nil
}

// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
// to avoid concurrent access problems.
// This is to avoid cloning the connection at each method access.
// Can return itself if not a problem.
func (s *Storage) Clone() osin.Storage {
	return s
}

// Close the resources the Storage potentially holds (using Clone for example)
func (s *Storage) Close() {
}

// GetClient loads the client by id
func (s *Storage) GetClient(id string) (osin.Client, error) {
	row := s.db.QueryRow("SELECT id, secret, redirect_uri, extra FROM client WHERE id=?", id)
	var c osin.DefaultClient
	var extra string

	if err := row.Scan(&c.Id, &c.Secret, &c.RedirectUri, &extra); err == sql.ErrNoRows {
		return nil, pkg.ErrNotFound
	} else if err != nil {
		return nil, errors.New(err)
	}
	c.UserData = extra
	return &c, nil
}

// UpdateClient updates the client (identified by it's id) and replaces the values with the values of client.
func (s *Storage) UpdateClient(c osin.Client) error {
	data, err := assertToString(c.GetUserData())
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("UPDATE client SET (secret, redirect_uri, extra) = (?, ?, ?) WHERE id=?", c.GetSecret(), c.GetRedirectUri(), data, c.GetId()); err != nil {
		return errors.New(err)
	}
	return nil
}

// CreateClient stores the client in the database and returns an error, if something went wrong.
func (s *Storage) CreateClient(c osin.Client) error {
	data, err := assertToString(c.GetUserData())
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("INSERT INTO client (id, secret, redirect_uri, extra) VALUES (?, ?, ?, ?)", c.GetId(), c.GetSecret(), c.GetRedirectUri(), data); err != nil {
		return errors.New(err)
	}
	return nil
}

// RemoveClient removes a client (identified by id) from the database. Returns an error if something went wrong.
func (s *Storage) RemoveClient(id string) (err error) {
	if _, err = s.db.Exec("DELETE FROM client WHERE id=?", id); err != nil {
		return errors.New(err)
	}
	return nil
}

// SaveAuthorize saves authorize data.
func (s *Storage) SaveAuthorize(data *osin.AuthorizeData) (err error) {
	extra, err := assertToString(data.UserData)
	if err != nil {
		return err
	}

	if _, err = s.db.Exec(
		"INSERT INTO authorize (client, code, expires_in, scope, redirect_uri, state, created_at, extra) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		data.Client.GetId(),
		data.Code,
		data.ExpiresIn,
		data.Scope,
		data.RedirectUri,
		data.State,
		data.CreatedAt,
		extra,
	); err != nil {
		return errors.New(err)
	}
	return nil
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	var data osin.AuthorizeData
	var extra string
	var cid string
	if err := s.db.QueryRow("SELECT client, code, expires_in, scope, redirect_uri, state, created_at, extra FROM authorize WHERE code=? LIMIT 1", code).Scan(&cid, &data.Code, &data.ExpiresIn, &data.Scope, &data.RedirectUri, &data.State, &data.CreatedAt, &extra); err == sql.ErrNoRows {
		return nil, pkg.ErrNotFound
	} else if err != nil {
		return nil, errors.New(err)
	}
	data.UserData = extra

	c, err := s.GetClient(cid)
	if err != nil {
		return nil, err
	}

	if data.ExpireAt().Before(time.Now()) {
		return nil, errors.Errorf("Token expired at %s.", data.ExpireAt().String())
	}

	data.Client = c
	return &data, nil
}

// RemoveAuthorize revokes or deletes the authorization code.
func (s *Storage) RemoveAuthorize(code string) (err error) {
	if _, err = s.db.Exec("DELETE FROM authorize WHERE code=?", code); err != nil {
		return errors.New(err)
	}
	return nil
}

// SaveAccess writes AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (s *Storage) SaveAccess(data *osin.AccessData) (err error) {
	prev := ""
	authorizeData := &osin.AuthorizeData{}

	if data.AccessData != nil {
		prev = data.AccessData.AccessToken
	}

	if data.AuthorizeData != nil {
		authorizeData = data.AuthorizeData
	}

	extra, err := assertToString(data.UserData)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return errors.New(err)
	}

	if data.RefreshToken != "" {
		if err := s.saveRefresh(tx, data.RefreshToken, data.AccessToken); err != nil {
			return err
		}
	}

	if data.Client == nil {
		return errors.New("data.Client must not be nil")
	}

	_, err = tx.Exec("INSERT INTO access (client, authorize, previous, access_token, refresh_token, expires_in, scope, redirect_uri, created_at, extra) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", data.Client.GetId(), authorizeData.Code, prev, data.AccessToken, data.RefreshToken, data.ExpiresIn, data.Scope, data.RedirectUri, data.CreatedAt, extra)
	if err != nil {
		if rbe := tx.Rollback(); rbe != nil {
			return errors.New(rbe)
		}
		return errors.New(err)
	}

	if err = tx.Commit(); err != nil {
		return errors.New(err)
	}

	return nil
}

// LoadAccess retrieves access data by token. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *Storage) LoadAccess(code string) (*osin.AccessData, error) {
	var extra, cid, prevAccessToken, authorizeCode string
	var result osin.AccessData

	if err := s.db.QueryRow(
		"SELECT client, authorize, previous, access_token, refresh_token, expires_in, scope, redirect_uri, created_at, extra FROM access WHERE access_token=? LIMIT 1",
		code,
	).Scan(
		&cid,
		&authorizeCode,
		&prevAccessToken,
		&result.AccessToken,
		&result.RefreshToken,
		&result.ExpiresIn,
		&result.Scope,
		&result.RedirectUri,
		&result.CreatedAt,
		&extra,
	); err == sql.ErrNoRows {
		return nil, pkg.ErrNotFound
	} else if err != nil {
		return nil, errors.New(err)
	}

	result.UserData = extra
	client, err := s.GetClient(cid)
	if err != nil {
		return nil, err
	}

	result.Client = client
	result.AuthorizeData, _ = s.LoadAuthorize(authorizeCode)
	prevAccess, _ := s.LoadAccess(prevAccessToken)
	result.AccessData = prevAccess
	return &result, nil
}

// RemoveAccess revokes or deletes an AccessData.
func (s *Storage) RemoveAccess(code string) (err error) {
	_, err = s.db.Exec("DELETE FROM access WHERE access_token=?", code)
	if err != nil {
		return errors.New(err)
	}
	return nil
}

// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *Storage) LoadRefresh(code string) (*osin.AccessData, error) {
	row := s.db.QueryRow("SELECT access FROM refresh WHERE token=? LIMIT 1", code)
	var access string
	if err := row.Scan(&access); err == sql.ErrNoRows {
		return nil, pkg.ErrNotFound
	} else if err != nil {
		return nil, errors.New(err)
	}
	return s.LoadAccess(access)
}

// RemoveRefresh revokes or deletes refresh AccessData.
func (s *Storage) RemoveRefresh(code string) error {
	_, err := s.db.Exec("DELETE FROM refresh WHERE token=$1", code)
	if err != nil {
		return errors.New(err)
	}
	return nil
}

// Makes easy to create a osin.DefaultClient
func (s *Storage) CreateClientWithInformation(id string, secret string, redirectUri string, userData interface{}) osin.Client {
	return &osin.DefaultClient{
		Id:          id,
		Secret:      secret,
		RedirectUri: redirectUri,
		UserData:    userData,
	}
}

func (s *Storage) saveRefresh(tx *sql.Tx, refresh, access string) (err error) {
	_, err = tx.Exec("INSERT INTO refresh (token, access) VALUES (?, ?)", refresh, access)
	if err != nil {
		if rbe := tx.Rollback(); rbe != nil {
			return errors.New(rbe)
		}
		return errors.New(err)
	}
	return nil
}

func assertToString(in interface{}) (string, error) {
	var ok bool
	var data string
	if in == nil {
		return "", nil
	} else if data, ok = in.(string); ok {
		return data, nil
	} else if str, ok := in.(fmt.Stringer); ok {
		return str.String(), nil
	}
	return "", errors.Errorf(`Could not assert "%v" to string`, in)
}
