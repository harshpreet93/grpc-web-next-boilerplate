// Your Custom Header

package predicate

import (
	"github.com/facebook/ent/dialect/sql"
)

// User is the predicate function for user builders.
type User func(*sql.Selector)
