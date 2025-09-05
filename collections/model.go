package collections

import "github.com/Meduzz/devserver/model"

type (
	CollectionsService interface {
		Lookup(name string) *model.Collection
	}
)
