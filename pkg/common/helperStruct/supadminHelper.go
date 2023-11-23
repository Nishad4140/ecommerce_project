package helper

type BlockData struct {
	UserId uint   ` json:"userid" validate:"required"`
	Reason string ` json:"reason" validate:"required"`
}

type AdminData struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Mobile   string `json:"mobile" validate:"required"`
	Password string `json:"password" validate:"required"`
}
