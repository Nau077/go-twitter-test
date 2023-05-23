package model

type SubsModel struct {
	Id     string `json:"id" binding:"required"`
	SubsId string `json:"subsId" binding:"required"`
}
