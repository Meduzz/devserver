package db

import (
	"github.com/etcd-io/bbolt"
)

var db *bbolt.DB

func Init() error {
	d, err := bbolt.Open("dev.db", 0666, nil)

	if err != nil {
		return nil
	}

	db = d
	return nil
}

func ReadBucket(bucket string, handler func(*bbolt.Bucket) error) error {
	return db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))

		return handler(bucket)
	})
}

func WriteBucket(bucket string, handler func(*bbolt.Bucket) error) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))

		if err != nil {
			return nil
		}

		return handler(bucket)
	})
}
