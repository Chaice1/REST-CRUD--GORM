package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Chaice/Postgres+GO/models"
	"github.com/Chaice/Postgres+GO/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *gin.Context) {
	book := Book{}

	err := context.BindJSON(&book)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "request failed"})
		return
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "could not create the book"})
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "book has created"})

}

func (r *Repository) GetBooks(context *gin.Context) {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "could not get books"})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "books fetched successfully",
		"data":    bookModels})
}
func (r *Repository) DeleteBook(context *gin.Context) {
	bookModel := models.Books{}
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "id cannot be empty",
		})
		return
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "could not delete book",
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "books delete successfully",
	})

}
func (r *Repository) GetBookByID(context *gin.Context) {
	id := context.Param("id")
	bookModel := &models.Books{}
	if id == "" {
		context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "id cannot be empty",
		})
		return
	}
	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "could not get the book",
		})
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "book get successfully",
	})
}
func (r *Repository) SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")
	api.POST("/create_books", r.CreateBook)
	api.DELETE("delete_book/:id", r.DeleteBook)
	api.GET("/get_books/:id", r.GetBookByID)
	api.GET("/books", r.GetBooks)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := gin.Default()

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	r := Repository{
		DB: db,
	}
	r.SetupRoutes(app)
	app.Run(":8080")
}
