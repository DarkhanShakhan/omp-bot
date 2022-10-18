package online

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/education"
)

const (
	NONEXIST = "online course with such id doesn't exist"
)

type OnlineService interface {
	Describe(onlineID uint64) (*education.Online, error)
	List(cursor uint64, limit uint64) ([]education.Online, error)
	Create(education.Online) (uint64, error)
	Update(onlineID uint64, online education.Online) error
	Remove(onlineID uint64) (bool, error)
}

type DummyOnlineService struct {
	maxID    uint64
	freeIDs  []uint64
	Entities map[uint64]education.Online
}

func NewDummyOnlineService() *DummyOnlineService {
	return &DummyOnlineService{maxID: 1, freeIDs: []uint64{}, Entities: map[uint64]education.Online{}}
}

func (d *DummyOnlineService) Describe(onlineID uint64) (*education.Online, error) {
	online, ok := d.Entities[onlineID]
	if !ok {
		return nil, errors.New(NONEXIST)
	}
	return &online, nil
}

func (d *DummyOnlineService) List(cursor uint64, limit uint64) ([]education.Online, error) {
	list := make([]education.Online, len(d.Entities))
	ix := 0
	for _, online := range d.Entities {
		list[ix] = online
		ix++
	}
	return list, nil
}

func (d *DummyOnlineService) Create(online education.Online) (uint64, error) {
	var id uint64
	if len(d.freeIDs) != 0 {
		id = d.freeIDs[0]
		d.freeIDs = d.freeIDs[1:]
	} else {
		id = d.maxID + 1
		d.maxID++
	}
	online.SetID(id)
	d.Entities[id] = online
	return id, nil
}

func (d *DummyOnlineService) Update(onlineID uint64, online education.Online) error {
	if _, ok := d.Entities[onlineID]; ok {
		d.Entities[onlineID] = online
		return nil
	}
	return errors.New(NONEXIST)
}

func (d *DummyOnlineService) Remove(onlineID uint64) (bool, error) {
	if _, ok := d.Entities[onlineID]; ok {
		delete(d.Entities, onlineID)
		d.freeIDs = append(d.freeIDs, onlineID)
		return true, nil
	}
	return false, errors.New(NONEXIST)
}
