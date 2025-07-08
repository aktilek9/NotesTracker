package service

import (
	"NotesTracker/modules"
	"NotesTracker/pkg/repository"
)

type Service interface {
	CreateNote(note modules.Note) modules.Note
	GetAllNotes() []modules.Note
	GetNoteById(id uint) (modules.Note, error)
	UpdateNote(id uint, updatedNote modules.Note) error
	DeleteNote(id uint) error
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

func (s *service) GetNoteById(id uint) (modules.Note, error) {
	return s.repo.GetById(id)
}

func (s *service) UpdateNote(id uint, updatedNote modules.Note) error {
	return s.repo.Update(id, updatedNote)
}

func (s *service) DeleteNote(id uint) error {
	return s.repo.Delete(id)
}
