package handler

import (
	"NotesTracker/modules"
	"NotesTracker/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{Service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		notes := api.Group("/notes")
		{
			notes.POST("/", h.createNote)
			notes.GET("/", h.getAllNotes)
			notes.GET("/:id", h.getNoteById)
			notes.PUT("/:id", h.updateNote)
			notes.DELETE("/:id", h.deleteNote)
		}
	}

	return router
}

func (h *Handler) createNote(c *gin.Context) {
	var note modules.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	created := h.Service.CreateNote(note)
	c.JSON(201, created)
}

func (h *Handler) getAllNotes(c *gin.Context) {
	notes := h.Service.GetAllNotes()
	c.JSON(200, notes)
}

func (h *Handler) getNoteById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	note, err := h.Service.GetNoteById(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(200, note)
}

func (h *Handler) updateNote(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var note modules.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Service.UpdateNote(uint(id), note); err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Note updated"})
}

func (h *Handler) deleteNote(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteNote(uint(id)); err != nil {
		c.JSON(404, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Note deleted"})
}
