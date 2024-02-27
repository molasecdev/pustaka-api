package controllers

import (
	"net/http"
	"pustaka-api/config"
	"pustaka-api/src/models"

	"github.com/gin-gonic/gin"
)

func GetBooksCount(c *gin.Context) {
	db := config.InitConfig()

	var booksCountByCategory []struct {
		Category string `json:"category"`
		Count    int    `json:"count"`
	}

	db.Table("books").Select("categories.category, count(*) as count").
		Joins("JOIN categories ON books.category_id = categories.id").
		Group("categories.category").Scan(&booksCountByCategory)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": booksCountByCategory})
}

func GetActiveLoanSummary(c *gin.Context) {
	db := config.InitConfig()
	var loan models.Loan

	var activeLoansCount int64
	var lateLoansCount int64
	var returnedLoansCount int64

	// Menghitung jumlah peminjaman aktif berdasarkan status
	db.Model(&loan).Where("status = ?", "belum dikembalikan").Count(&activeLoansCount)
	db.Model(&loan).Where("status = ?", "terlambat").Count(&lateLoansCount)
	db.Model(&loan).Where("status = ?", "dikembalikan").Count(&returnedLoansCount)

	// Mengembalikan data sebagai JSON
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"Peminjam aktif":        activeLoansCount,
			"Peminjam terlambat":    lateLoansCount,
			"Peminjam dikembalikan": returnedLoansCount,
		},
	})
}

func GetUserStatistics(c *gin.Context) {
	db := config.InitConfig()
	var loan models.Loan
	var user models.User

	var registeredUsersCount int64
	var totalBooksBorrowed int64
	var averageBorrowedBooksPerUser float64

	// Menghitung jumlah pengguna terdaftar
	db.Model(&user).Count(&registeredUsersCount)

	// Menghitung total buku yang telah dipinjam oleh semua pengguna
	db.Model(&loan).Count(&totalBooksBorrowed)

	// Menghitung rata-rata buku yang dipinjam per pengguna
	db.Table("loans").Select("COUNT(user_id) / COUNT(DISTINCT user_id) AS avg_books_per_user").Scan(&averageBorrowedBooksPerUser)

	// Mengembalikan data sebagai JSON
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"User terdaftar":          registeredUsersCount,
			"Total buku dipinjam":     totalBooksBorrowed,
			"Rata-rata buku dipinjam": averageBorrowedBooksPerUser,
		},
	})
}
