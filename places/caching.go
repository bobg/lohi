package places

import (
	"context"
	"encoding/json"

	"github.com/bobg/errors"
	"github.com/boltdb/bolt"
)

type CachingService struct {
	db   *bolt.DB
	next Service
}

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

func (s *CachingService) Close() error {
	return s.db.Close()
}

var bucketName = []byte("places")

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
