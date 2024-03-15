package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"pustaka-api/types"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	r := c.Request
	maxSize := int64(5 * 1024 * 1024) // 5 MB

	dir, errDir := os.Getwd()
	if errDir != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDir.Error()})
		return
	}

	// get file
	if err := r.ParseMultipartForm(1024); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	file, fileHeader, errFile := r.FormFile("file")
	if errFile != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFile.Error()})
		return
	}
	defer file.Close()

	// type & size validations
	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only image files are allowed!"})
		return
	}
	if fileHeader.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large!"})
		return
	}

	// generate UUID for the filename & save file
	fileExt := filepath.Ext(fileHeader.Filename)
	filename := uuid.New().String() + fileExt
	fileLocation := filepath.Join(dir, "files/img", filename)

	targetFile, errCreate := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
		return
	}
	defer targetFile.Close()

	if _, errCopy := io.Copy(targetFile, file); errCopy != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCopy.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": gin.H{"filename": filename}})
}

func GetFile(c *gin.Context) {
	filename := c.Param("filename")

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fileLocation := filepath.Join(dir, "files/img", filename)

	// check file
	_, err = os.Stat(fileLocation)
	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found!"})
		return
	}

	c.File(fileLocation)
}

func UpdateFile(c *gin.Context) {
	filename := c.Param("filename")
	maxSize := int64(5 * 1024 * 1024) // 5 MB

	dir, errDir := os.Getwd() // get location
	if errDir != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDir.Error()})
		return
	}

	// delete old file
	oldFileLocation := filepath.Join(dir, "files/img", filename)
	if errFileLocation := os.Remove(oldFileLocation); errFileLocation != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFileLocation.Error()})
		return
	}

	// get newFile
	newFile, newFileHeader, errFile := c.Request.FormFile("file")
	if errFile != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFile.Error()})
		return
	}
	defer newFile.Close()

	// type & size validations
	contentType := newFileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only image files are allowed!"})
		return
	}
	if newFileHeader.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large!"})
		return
	}

	// Generate UUID for the filename & save new file
	fileExt := filepath.Ext(newFileHeader.Filename)
	newFilename := uuid.New().String() + fileExt
	newFileLocation := filepath.Join(dir, "files/img", newFilename)

	newTargetFile, errCreate := os.OpenFile(newFileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
		return
	}
	defer newTargetFile.Close()

	if _, errCopy := io.Copy(newTargetFile, newFile); errCopy != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCopy.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "filename": newFilename})
}

func DeleteFile(c *gin.Context) {
	filename := c.Param("filename")

	dir, errDir := os.Getwd() // get location
	if errDir != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDir.Error()})
		return
	}

	// delete old file
	fileLocation := filepath.Join(dir, "files/img", filename)
	if errFileLocation := os.Remove(fileLocation); errFileLocation != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errFileLocation.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "File deleted successfully"})
}

func ImportExcelBooks(c *gin.Context) {
	db := config.InitConfig()

	filePath := "/home/maulana/Desktop/pustaka-api/files/excel/File Master Create Books.xlsx"
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path should be : 'files/excel/File Master Create Books.xlsx'"})
	}

	sheet1Name := "Sheet1"

	type M map[string]interface{}
	rows := make([]M, 0)

	startRow := 3
	dataLength := 1000
	for i := startRow; i <= dataLength; i++ {
		row := M{
			"title":       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"description": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			"stock":       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),
			"isbn":        xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)),
			"year":        xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", i)),
			"pages":       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i)),
			"category":    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", i)),
			"author":      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i)),
			"publisher":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", i)),
			"language":    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", i)),
			"bookshelf":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", i)),
		}
		// hanya untuk cek data yang tersedia
		if row["title"] != "" {
			rows = append(rows, row)
		}
	}

	for _, val := range rows {
		var findAuthor models.Author
		var findPublisher models.Publisher
		var findCategory models.Category
		var findBookshelf models.Bookshelfs
		var findLanguage models.Language
		var existingBook models.Book

		isbnFloat, _ := strconv.ParseFloat(val["isbn"].(string), 64)
		yearFloat, _ := strconv.ParseFloat(val["year"].(string), 64)
		pagesFloat, _ := strconv.ParseFloat(val["pages"].(string), 64)
		stockFloat, _ := strconv.ParseFloat(val["stock"].(string), 64)
		title := val["title"].(string)
		description := val["description"].(string)
		isbn := strconv.FormatFloat(isbnFloat, 'f', -1, 64)
		year := strconv.FormatFloat(yearFloat, 'f', -1, 64)
		pages := int(pagesFloat)
		stock := int(stockFloat)

		errFAuthor := db.Where("LOWER(author) = LOWER(?)", val["author"]).First(&findAuthor)
		errFPublisher := db.Where("LOWER(publisher) = LOWER(?)", val["publisher"]).First(&findPublisher)
		errFCategory := db.Where("LOWER(category) = LOWER(?)", val["category"]).First(&findCategory)
		errFBookshelf := db.Where("LOWER(bookshelf) = LOWER(?)", val["bookshelf"]).First(&findBookshelf)
		errFLanguage := db.Where("LOWER(language) = LOWER(?)", val["language"]).First(&findLanguage)
		db.Where("isbn = ?", isbn).First(&existingBook)

		// jika data tidak ada maka akan auto create
		if errFAuthor.Error != nil {
			if findAuthor.Author != val["author"] {
				author := models.Author{
					Author: val["author"].(string),
				}
				db.Create(&author)
				findAuthor = author
			}
		}
		if errFPublisher.Error != nil {
			if findPublisher.Publisher != val["publisher"] {
				publisher := models.Publisher{
					Publisher: val["publisher"].(string),
				}
				db.Create(&publisher)
				findPublisher = publisher
			}
		}
		if errFCategory.Error != nil {
			if findCategory.Category != val["category"] {
				category := models.Category{
					Category: val["category"].(string),
				}
				db.Create(&category)
				findCategory = category
			}
		}
		if errFBookshelf.Error != nil {
			if findBookshelf.Bookshelf != val["bookshelf"] {
				bookshelf := models.Bookshelfs{
					Bookshelf: val["bookshelf"].(string),
				}
				db.Create(&bookshelf)
				findBookshelf = bookshelf
			}
		}
		if errFLanguage.Error != nil {
			if findLanguage.Language != val["language"] {
				language := models.Language{
					Language: val["language"].(string),
				}
				db.Create(&language)
				findLanguage = language
			}
		}

		// jika book sudah ada, maka akan di skip
		if existingBook.ID != uuid.Nil {
			continue
		}

		book := models.Book{
			Title:        title,
			Description:  description,
			Stock:        stock,
			Isbn:         isbn,
			Year:         year,
			Pages:        pages,
			Author_id:    findAuthor.ID,
			Publisher_id: findPublisher.ID,
			Category_id:  findCategory.ID,
			Bookshelf_id: findBookshelf.ID,
			Language_id:  findLanguage.ID,
		}
		db.Create(&book)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "File imported successfully!"})
}

func ExportLoans(c *gin.Context) {
	db := config.InitConfig()
	var loans []models.Loan
	var datas []types.LoanDetails

	query := db.Preload("User").Preload("User.Auth").Preload("User.Role").Preload("Book").Preload("Book.Author").Preload("Book.Category").Preload("Book.Publisher").Preload("Book.Bookshelf").Preload("Book.Language")

	errFLoans := query.Find(&loans)
	if errFLoans.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loans not found!"})
		return
	}

	for _, value := range loans {
		var loanDetail types.LoanDetails
		loanDetail.FullName = value.User.Firstname + " " + value.User.Lastname
		loanDetail.Title = value.Book.Title
		loanDetail.StartDate = (value.Start_date).String()
		loanDetail.EndDate = (value.End_date).String()
		loanDetail.Note = value.Note
		loanDetail.Status = value.Status
		loanDetail.ReturnDate = value.Return_date
		loanDetail.Penalty = value.Penalty
		datas = append(datas, loanDetail)
	}

	xlsx := excelize.NewFile()
	sheetName := "List loans"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

	xlsx.SetCellValue(sheetName, "A2", "Full Name")
	xlsx.SetCellValue(sheetName, "B2", "Book Title")
	xlsx.SetCellValue(sheetName, "C2", "Start Date")
	xlsx.SetCellValue(sheetName, "D2", "End Date")
	xlsx.SetCellValue(sheetName, "E2", "Note")
	xlsx.SetCellValue(sheetName, "F2", "Status")
	xlsx.SetCellValue(sheetName, "G2", "Return Date")
	xlsx.SetCellValue(sheetName, "H2", "Penalty")

	errFilter := xlsx.AutoFilter(sheetName, "A2", "H2", "")
	if errFilter != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errFilter.Error()})
		return
	}

	for i, loan := range datas {
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(i+3), loan.FullName)
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(i+3), loan.Title)
		xlsx.SetCellValue(sheetName, "C"+strconv.Itoa(i+3), loan.StartDate)
		xlsx.SetCellValue(sheetName, "D"+strconv.Itoa(i+3), loan.EndDate)
		xlsx.SetCellValue(sheetName, "E"+strconv.Itoa(i+3), loan.Note)
		xlsx.SetCellValue(sheetName, "F"+strconv.Itoa(i+3), loan.Status)
		xlsx.SetCellValue(sheetName, "G"+strconv.Itoa(i+3), loan.ReturnDate)
		xlsx.SetCellValue(sheetName, "H"+strconv.Itoa(i+3), loan.Penalty)
	}

	now := time.Now().Unix()
	errSave := xlsx.SaveAs("./files/excel/loans-" + strconv.FormatInt(now, 10) + ".xlsx")
	if errSave != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errSave.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Export loans data to excel file successfully"})
}
