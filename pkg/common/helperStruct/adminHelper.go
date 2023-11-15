package helper

type ReportData struct {
	UserId uint   ` json:"userid" validate:"required"`
	Reason string ` json:"reason" validate:"required"`
}

type CreateSeller struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password"`
}

type ReportParams struct {
	Status string `json:"status"`
	Year   int    `json:"year"`
	Month  int    `json:"month"`
	Week   int    `json:"week"`
	Day    int    `json:"day"`
	Date1  string `json:"date1"`
	Date2  string `json:"date2"`
}
