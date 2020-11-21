// Your Custom Header

package migrate

import (
	"context"
	"fmt"
	"io"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql/schema"
)

var (
	// WithGlobalUniqueID sets the universal ids options to the migration.
	// If this option is enabled, ent migration will allocate a 1<<32 range
	// for the ids of each entity (table).
	// Note that this option cannot be applied on tables that already exist.
	WithGlobalUniqueID = schema.WithGlobalUniqueID
	// WithDropColumn sets the drop column option to the migration.
	// If this option is enabled, ent migration will drop old columns
	// that were used for both fields and edges. This defaults to false.
	WithDropColumn = schema.WithDropColumn
	// WithDropIndex sets the drop index option to the migration.
	// If this option is enabled, ent migration will drop old indexes
	// that were defined in the schema. This defaults to false.
	// Note that unique constraints are defined using `UNIQUE INDEX`,
	// and therefore, it's recommended to enable this option to get more
	// flexibility in the schema changes.
	WithDropIndex = schema.WithDropIndex
	// WithFixture sets the foreign-key renaming option to the migration when upgrading
	// ent from v0.1.0 (issue-#285). Defaults to false.
	WithFixture = schema.WithFixture
	// WithForeignKeys enables creating foreign-key in schema DDL. This defaults to true.
	WithForeignKeys = schema.WithForeignKeys
)

// Schema is the API for creating, migrating and dropping a schema.
type Schema struct {
	drv         dialect.Driver
	universalID bool
}

// NewSchema creates a new schema client.
func NewSchema(drv dialect.Driver) *Schema { return &Schema{drv: drv} }

// Create creates all schema resources.
func (s *Schema) Create(ctx context.Context, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %v", err)
	}
	return migrate.Create(ctx, Tables...)
}

// WriteTo writes the schema changes to w instead of running them against the database.
//
// 	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
//		log.Fatal(err)
// 	}
//
func (s *Schema) WriteTo(ctx context.Context, w io.Writer, opts ...schema.MigrateOption) error {
	drv := &schema.WriteDriver{
		Writer: w,
		Driver: s.drv,
	}
	migrate, err := schema.NewMigrate(drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %v", err)
	}
	return migrate.Create(ctx, Tables...)
}
