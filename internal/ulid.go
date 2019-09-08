package wishlist

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// UlidGen encapsulate the way to generate ulids
func UlidGen() string {
	t := time.Now().UTC()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)

	return id.String()
}
