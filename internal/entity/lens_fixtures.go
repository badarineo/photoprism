package entity

import (
	"time"
)

type LensMap map[string]Lens

func (m LensMap) Get(name string) Lens {
	if result, ok := m[name]; ok {
		return result
	}

	return *NewLens(name, "")
}

func (m LensMap) Pointer(name string) *Lens {
	if result, ok := m[name]; ok {
		return &result
	}

	return NewLens(name, "")
}

var LensFixtures = LensMap{
	"lens-f-380": {
		ID:              1000000,
		LensSlug:        "lens-f-380",
		LensModel:       "F380",
		LensMake:        "Apple",
		LensType:        "",
		LensOwner:       "Test",
		LensDescription: "",
		LensNotes:       "Notes",
		CreatedAt:       time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:       time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt:       nil,
	},
}

// CreateLensFixtures inserts known entities into the database for testing.
func CreateLensFixtures() {
	for _, entity := range LensFixtures {
		Db().Create(&entity)
	}
}
