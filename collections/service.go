package collections

import (
	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	collectionService struct {
		collections []*model.Collection
	}
)

var (
	_ model.Service      = (*collectionService)(nil)
	_ CollectionsService = (*collectionService)(nil)
)

func NewCollectionService(collections []*model.Collection) model.Service {
	return &collectionService{
		collections: collections,
	}
}

func (c *collectionService) Start() error {
	return nil
}

func (c *collectionService) Stop() error {
	return nil
}

func (c *collectionService) Lookup(name string) *model.Collection {
	return slice.Head(slice.Filter(c.collections, func(it *model.Collection) bool {
		return it.Name == name
	}))
}
