package types

import "time"

type InputLoan struct {
	Nik  string `json:"nik" binding:"required"`
	Isbn string `json:"isbn" binding:"required"`
	Note string `json:"note" binding:"required"`
}

type UpdateLoan struct {
	Status string `json:"status" binding:"required"`
}

type LoanDetails struct {
	FullName   string     `json:"full_name"`
	Title      string     `json:"title"`
	StartDate  string     `json:"start_date"`
	EndDate    string     `json:"end_date"`
	Note       string     `json:"note"`
	Status     string     `json:"status"`
	ReturnDate *time.Time `json:"return_date"`
	Penalty    int        `json:"penalty"`
}
