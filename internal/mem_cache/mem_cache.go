package mem_cache

import "encoding/json"

type Cache struct {
	storage map[string][]byte
}

func NewCache() *Cache {
	return &Cache{
		storage: map[string][]byte{},
	}
}

func (c *Cache) Read(key string, v any) (bool, error) {
	data, ok := c.storage[key]
	if !ok {
		return false, nil
	}

	err := json.Unmarshal(data, v)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Cache) Write(key string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.storage[key] = data

	return nil
}
