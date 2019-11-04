package wishlist

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// NewULID encapsulate the way to generate new ULIDs
func NewULID() string {
	t := time.Now().UTC()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)

	return id.String()
}
