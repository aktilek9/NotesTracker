package repository

import (
	"NotesTracker/modules"
	"errors"
)

type Repository struct {
	db []modules.Note
}

func NewRepository() *Repository {
	return &Repository{db: make([]modules.Note, 0)}
}

func (r *Repository) Create(note modules.Note) modules.Note {
	if len(r.db) == 0 {
		note.ID = 1
	} else {
		note.ID = r.db[len(r.db)-1].ID + 1
	}
	r.db = append(r.db, note)

	return note
}

func (r *Repository) AllNotes() []modules.Note {
	return r.db
}

func (r *Repository) GetById(id int) (modules.Note, error) {

	for _, value := range r.db {
		if value.ID == id {
			return value, nil
		}
	}
	return modules.Note{}, errors.New("note not found")
}

func (r *Repository) Delete(id int) error {

	var index int = -1
	for idx, value := range r.db {
		if value.ID == id {
			index = idx
			break
		}
	}

	if index != -1 {
		r.db = append(r.db[:index], r.db[index+1:]...)
		return nil
	}
	return errors.New("note not found")
}
