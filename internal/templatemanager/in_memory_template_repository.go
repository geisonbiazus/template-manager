package templatemanager

import (
	"errors"
	"strconv"
)

type InMemoryTemplateRepository struct {
	Data   map[string]Template
	lastID int
}

func NewInMemoryTemplateRepository() *InMemoryTemplateRepository {
	return &InMemoryTemplateRepository{
		Data: make(map[string]Template),
	}
}

func (r *InMemoryTemplateRepository) Create(t Template) (Template, error) {
	r.lastID++
	t.ID = strconv.Itoa(r.lastID)
	r.Data[t.ID] = t
	return t, nil
}

func (r *InMemoryTemplateRepository) FindByID(id string) (Template, error) {
	t, found := r.Data[id]
	if !found {
		return t, errors.New("Template not found")
	}
	return t, nil
}
