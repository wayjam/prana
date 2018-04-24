// Package sqlmigr provides primitives and functions to work with SQL
// sqlmigrs.
package sqlmigr

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/phogolabs/parcello"
)

//go:generate counterfeiter -fake-name MigrationRunner -o ../fake/MigrationRunner.go . MigrationRunner
//go:generate counterfeiter -fake-name MigrationProvider -o ../fake/MigrationProvider.go . MigrationProvider
//go:generate counterfeiter -fake-name MigrationGenerator -o ../fake/MigrationGenerator.go . MigrationGenerator

var (
	format = "20060102150405"
	min    = time.Date(1, time.January, 1970, 0, 0, 0, 0, time.UTC)
)

// FileSystem provides with primitives to work with the underlying file system
type FileSystem = parcello.FileSystem

// MigrationRunner runs or reverts a given sqlmigr item.
type MigrationRunner interface {
	// Run runs a given sqlmigr item.
	Run(item *Migration) error
	// Revert reverts a given sqlmigr item.
	Revert(item *Migration) error
}

// MigrationProvider provides all items.
type MigrationProvider interface {
	// Migrations returns all sqlmigr items.
	Migrations() ([]Migration, error)
	// Insert inserts executed sqlmigr item in the sqlmigrs table.
	Insert(item *Migration) error
	// Delete deletes applied sqlmigr item from sqlmigrs table.
	Delete(item *Migration) error
	// Exists returns true if the sqlmigr exists
	Exists(item *Migration) bool
}

// MigrationGenerator generates a migration item file.
type MigrationGenerator interface {
	// Create creates a new sqlmigr.
	Create(m *Migration) error
	// Write creates a new sqlmigr for given content.
	Write(m *Migration, content *Content) error
}

// Content represents a migration content.
type Content struct {
	// UpCommand is the content for upgrade operation.
	UpCommand io.Reader
	// DownCommand is the content for rollback operation.
	DownCommand io.Reader
}

// Migration represents a single migration record.
type Migration struct {
	// Id is the primary key for this sqlmigr
	ID string `db:"id"`
	// Description is the short description of this sqlmigr.
	Description string `db:"description"`
	// CreatedAt returns the time of sqlmigr execution.
	CreatedAt time.Time `db:"created_at"`
}

// Filename returns the item filename
func (m Migration) Filename() string {
	return fmt.Sprintf("%s_%s.sql", m.ID, m.Description)
}

// Parse parses a given file path to a sqlmigr item.
func Parse(path string) (*Migration, error) {
	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	parts := strings.SplitN(name, "_", 2)
	parseErr := fmt.Errorf("Migration '%s' has an invalid file name", path)

	if len(parts) != 2 {
		return nil, parseErr
	}

	if _, err := time.Parse(format, parts[0]); err != nil {
		return nil, parseErr
	}

	return &Migration{
		ID:          parts[0],
		Description: parts[1],
	}, nil
}