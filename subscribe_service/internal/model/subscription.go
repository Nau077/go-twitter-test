package model

type SubsModel struct {
	User     string `json:"user" binding:"required"`
	SubsUser string `json:"subsUser" binding:"required"`
}
