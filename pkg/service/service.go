package service

import (
	"NotesTracker/modules"
	"NotesTracker/pkg/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) CreateNote(note modules.Note) modules.Note {
	return s.repo.Create(note)
}

func (s *Service) GetAllNotes() []modules.Note {
	return s.repo.AllNotes()
}

func (s *Service) GetNoteById(id int) (modules.Note, error) {
	return s.repo.GetById(id)
}

func (s *Service) UpdateNote(id int, updatedNote modules.Note) error {
	return s.repo.Update(id, updatedNote)
}

func (s *Service) DeleteNote(id int) error {
	return s.repo.Delete(id)
}
