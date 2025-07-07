package service

import (
	"NotesTracker/modules"
	"NotesTracker/pkg/repository"
)

type Service interface {
	CreateNote(note modules.Note) modules.Note
	GetAllNotes() []modules.Note
	GetNoteById(id int) (modules.Note, error)
	UpdateNote(id int, updatedNote modules.Note) error
	DeleteNote(id int) error
}

type service struct {
	repo repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repo: repository}
}

func (s *service) CreateNote(note modules.Note) modules.Note {
	return s.repo.Create(note)
}

func (s *service) GetAllNotes() []modules.Note {
	return s.repo.AllNotes()
}

func (s *service) GetNoteById(id int) (modules.Note, error) {
	return s.repo.GetById(id)
}

func (s *service) UpdateNote(id int, updatedNote modules.Note) error {
	return s.repo.Update(id, updatedNote)
}

func (s *service) DeleteNote(id int) error {
	return s.repo.Delete(id)
}
