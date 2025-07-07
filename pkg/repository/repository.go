package repository

import (
	"NotesTracker/modules"
	"errors"
)

type repository struct {
	db []modules.Note
}

type Repository interface {
	Create(modules.Note) modules.Note
	AllNotes() []modules.Note
	GetById(id int) (modules.Note, error)
	Update(id int, updatedNote modules.Note) error
	Delete(id int) error
}

func NewRepository() *repository {
	return &repository{db: make([]modules.Note, 0)}
}

func (r *repository) Create(note modules.Note) modules.Note {
	if len(r.db) == 0 {
		note.ID = 1
	} else {
		note.ID = r.db[len(r.db)-1].ID + 1
	}
	r.db = append(r.db, note)

	return note
}

func (r *repository) AllNotes() []modules.Note {
	return r.db
}

func (r *repository) GetById(id int) (modules.Note, error) {

	for _, value := range r.db {
		if value.ID == id {
			return value, nil
		}
	}
	return modules.Note{}, errors.New("note not found")
}

func (r *repository) Update(id int, updatedNote modules.Note) error {

	for i, value := range r.db {
		if value.ID == id {
			updatedNote.ID = id
			r.db[i] = updatedNote
			return nil
		}
	}

	return errors.New("note not found")
}

func (r *repository) Delete(id int) error {

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
