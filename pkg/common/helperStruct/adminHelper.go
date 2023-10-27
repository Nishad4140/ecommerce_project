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
