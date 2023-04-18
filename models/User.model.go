package models

type User struct {
	BaseModel
	Email string `gorm:"column:email;type:varchar;size:255;uniqueIndex;not null;unique" json:"email" bson:"email"`
	Photo string `gorm:"column:photo" json:"photo" bson:"photo"`
}

type UserResponse struct {

}
