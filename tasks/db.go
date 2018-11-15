package tasks

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"time"
)

type DBRepository struct {
	db *bolt.DB
}

func NewDBManager(path string) (DBRepository, error) {
	db, err := bolt.Open(path, 0600, nil)
	return DBRepository{db: db,}, err
}

func (dbm DBRepository) Add(time time.Time, task Task) {
	bucketKey := KeyByDay(time)
	dbm.db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketKey))

		if err != nil {
			return err
		}

		return bucket.Put([]byte(task.Description), task.toJson())
	})
}

func (dbm DBRepository) Remove(time time.Time) {
	bucketKey := KeyByDay(time)
	dbm.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucketKey))
	})
}

func (dbm DBRepository) List(time time.Time) []Task {
	bucketKey := KeyByDay(time)
	tasks := make([]Task, 0)
	dbm.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketKey))
		if b == nil {
			return nil
		}
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			tasks = append(tasks, task)
		}
		return nil
	})

	return tasks
}
