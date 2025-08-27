package cache

import (
	"errors"
	"sync"
)

var once sync.Once

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type cache struct {
	data map[string]interface{}
}

func (c *cache) Set(key string, value interface{}) error {
	c.data[key] = value
	return nil
}

func (c *cache) Get(key string) (interface{}, error) {
	value, ok := c.data[key]

	if !ok {
		return nil, errors.New("key:" + key + " not found")
	}

	return value, nil
}

func (c *cache) Delete(key string) error {
	_, ok := c.data[key]

	if !ok {
		return errors.New("key:" + key + " not found")
	}

	delete(c.data, key)

	return nil
}

var singleInstance Cache

func New() Cache {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &cache{data: make(map[string]interface{})}
			},
		)
	}

	return singleInstance
}
