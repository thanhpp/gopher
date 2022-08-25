package filestorage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JSONStorage struct {
	lock sync.RWMutex
	path string
	f    *os.File
	mem  map[string]interface{}
}

func NewJSONStorage(path string) (*JSONStorage, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("new json storage - open file %s error: %w", path, err)
	}

	s := &JSONStorage{
		path: path,
		f:    f,
		mem:  make(map[string]interface{}),
	}

	return s, nil
}

func (s *JSONStorage) Get(key string) (interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, ok := s.mem[key]
	if !ok {
		return nil, fmt.Errorf("json storage - get - key %s not found", key)
	}

	return v, nil
}

func (s *JSONStorage) Set(key string, value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.mem[key] = value

	if err := s.flush(); err != nil {
		return err
	}

	return nil
}

func (s *JSONStorage) Del(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.mem, key)

	if err := s.flush(); err != nil {
		return err
	}

	return nil
}

// flush not concurrent safe.
func (s *JSONStorage) flush() error {
	data, err := json.Marshal(s.mem)
	if err != nil {
		return fmt.Errorf("json storage - marshal mem error: %w", err)
	}

	stat, err := s.f.Stat()
	if err != nil {
		return fmt.Errorf("json storage - get file stat error: %w", err)
	}

	if stat.Size() != 0 {
		if err := s.f.Truncate(0); err != nil {
			return fmt.Errorf("json storage - truncate file error: %w", err)
		}
	}

	if _, err := s.f.Write(data); err != nil {
		return fmt.Errorf("json storage - write file error: %w", err)
	}

	return nil
}
