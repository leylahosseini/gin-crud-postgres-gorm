package model

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`                                  // @primary
	Name  string `json:"name" gorm:"not null" binding:"required"`               // @required
	Email string `json:"email" gorm:"unique;not null" binding:"required,email"` // @required

}
