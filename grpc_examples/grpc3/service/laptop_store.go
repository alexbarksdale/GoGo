package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexbarksdale/GoGo/grpc_examples/pb"
	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemLaptopStore() *InMemLaptopStore {
	return &InMemLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (s *InMemLaptopStore) Save(laptop *pb.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Laptop{}
	if err := copier.Copy(other, laptop); err != nil {
		return fmt.Errorf("cannot copy laptop: %w", err)
	}

	s.data[other.Id] = other
	return nil

}
