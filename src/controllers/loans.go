package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context) {
	db := config.InitConfig()
	var book models.Book
	var user models.User
	var body types.InputLoan

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	if errFBook := db.Where("isbn = ?", body.Isbn).First(&book).Error; errFBook != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}
	if errFUser := db.Where("nik = ?", body.Nik).First(&user).Error; errFUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	const day int = 7
	loans := models.Loan{
		Book_id:    book.ID,
		User_id:    user.ID,
		Note:       body.Note,
		Status:     "belum dikembalikan",
		Start_date: time.Now(),
		End_date:   time.Now().AddDate(0, 0, day), // AddDate(years,months,days)
	}
	books := models.Book{
		Stock: book.Stock - 1,
	}

	errCreateLoan := db.Create(&loans)
	if errCreateLoan.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreateLoan.Error})
		return
	}
	errUpdateBook := db.Model(&books).Where("id = ?", book.ID).Updates(books)
	if errUpdateBook.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errUpdateBook.Error})
		return
	}

	var createLoan models.Loan
	db.Where("id = ?", loans.ID).Preload("User").Preload("User.Auth").Preload("User.Role").Preload("Book").Preload("Book.Author").Preload("Book.Category").Preload("Book.Publisher").Preload("Book.Bookshelf").Preload("Book.Language").First(&createLoan)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": createLoan})
}

func GetAllLoans(c *gin.Context) {
	db := config.InitConfig()
	status := c.Query("status")
	var loans []models.Loan

	query := db.Preload("User").Preload("User.Auth").Preload("User.Role").Preload("Book").Preload("Book.Author").Preload("Book.Category").Preload("Book.Publisher").Preload("Book.Bookshelf").Preload("Book.Language")

	// Filter berdasarkan status jika status diberikan
	if status != "" {
		query = query.Where("LOWER(status) = LOWER(?)", status)
	}

	// Eksekusi query dan ambil data pinjaman
	errFLoans := query.Find(&loans)
	if errFLoans.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loans not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": loans})
}

func GetLoanById(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var loan models.Loan

	errFLoan := db.Where("id = ?", paramId).Preload("User").Preload("User.Auth").Preload("User.Role").Preload("Book").Preload("Book.Author").Preload("Book.Category").Preload("Book.Publisher").Preload("Book.Bookshelf").Preload("Book.Language").First(&loan)
	if errFLoan.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": loan})
}

func UpdateLoan(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var body types.UpdateLoan
	var loan models.Loan
	var book models.Book

	if errBind := c.ShouldBindJSON(&body); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errFLoan := db.Where("id = ?", paramId).First(&loan)
	if errFLoan.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found!"})
		return
	}
	db.Where("id = ?", loan.Book_id).First(&book)

	if loan != (models.Loan{}) {
		returnDate := time.Now()
		var dataLoan interface{}
		var dataBook interface{}

		if strings.ToLower(body.Status) == "dikembalikan" && loan.Return_date == nil {
			dataLoan = models.Loan{
				Status:      strings.ToLower(body.Status),
				Return_date: &returnDate,
			}
			dataBook = models.Book{
				Stock: book.Stock + 1,
			}
		}

		db.Model(&loan).Where("id = ?", paramId).Updates(dataLoan)
		db.Model(&book).Where("id = ?", loan.Book_id).Updates(dataBook)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": loan})
}

func DeleteLoan(c *gin.Context) {
	db := config.InitConfig()
	paramId := c.Param("id")
	var loan models.Loan

	errFLoan := db.Where("id = ?", paramId).First(&loan)
	if errFLoan.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found!"})
		return
	}

	db.Where("id = ?", paramId).Delete(&loan)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete loan successfully!"})
}
