package repository

import (
	"NotesTracker/modules"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Create(modules.Note) modules.Note
	AllNotes() []modules.Note
	GetById(id uint) (modules.Note, error)
	Update(id uint, updatedNote modules.Note) error
	Delete(id uint) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(note modules.Note) modules.Note {
	r.db.Create(&note)
	return note
}

func (r *repository) AllNotes() []modules.Note {
	var notes []modules.Note
	r.db.Find(&notes)
	return notes
}

func (r *repository) GetById(id uint) (modules.Note, error) {
	var note modules.Note
	result := r.db.First(&note, id)
	if result.Error != nil {
		return modules.Note{}, result.Error
	}
	return note, nil
}

func (r *repository) Update(id uint, updatedNote modules.Note) error {
	var note modules.Note
	result := r.db.First(&note, id)
	if result.Error != nil {
		return result.Error
	}

	r.db.Model(&note).Updates(updatedNote)
	return nil
}

func (r *repository) Delete(id uint) error {
	result := r.db.Delete(&modules.Note{}, id)
	return result.Error
}
