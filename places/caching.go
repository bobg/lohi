package places

import (
	"context"
	"encoding/json"

	"github.com/bobg/errors"
	"github.com/boltdb/bolt"
)

// CachingService is a Service that permanently caches place lookups in a bolt database.
type CachingService struct {
	db   *bolt.DB
	next Service
}

// NewCachingService creates a new CachingService that "wraps" the given Service and caches lookups in the given filename.
func NewCachingService(filename string, next Service) (*CachingService, error) {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, errors.Wrap(err, "opening cache database")
	}
	return &CachingService{
		db:   db,
		next: next,
	}, nil
}

// Close closes the cache database.
func (s *CachingService) Close() error {
	return s.db.Close()
}

var bucketName = []byte("places")

// GetPlace looks up a [Place] by ID.
// If the place is in the cache, it is returned.
// Otherwise the wrapped Service is called to fetch it,
// and the result is stored in the cache before being returned.
func (s *CachingService) GetPlace(ctx context.Context, id string) (*Place, error) {
	var place *Place
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			var err error
			b, err = tx.CreateBucket(bucketName)
			if err != nil {
				return errors.Wrap(err, "creating bucket")
			}
		}
		if v := b.Get([]byte(id)); v != nil {
			return json.Unmarshal(v, &place)
		}

		var err error
		place, err = s.next.GetPlace(ctx, id)
		if err != nil {
			return errors.Wrapf(err, "in cache miss for place %s", id)
		}
		v, err := json.MarshalIndent(place, "", "  ")
		if err != nil {
			return errors.Wrap(err, "marshaling place")
		}
		err = b.Put([]byte(id), v)
		return errors.Wrapf(err, "adding place %s to cache", id)
	})
	return place, err
}
