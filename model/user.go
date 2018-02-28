package model

type User struct {
	Base
	Name         string  `json:"name" bson:"name"`
	Email        string  `json:"email" bson:"email"`
	Password     string  `json:"password" bson:"-"`
	passwordHash string  `bson:"password_hash"`
	CurrentDots  float32 `json:"current_dots" bson:"current_dots"`
	TotalDots    float32 `json:"total_dots" bson:"total_dots"`
	Team         string  `json:"team" bson:"team"`
	UserType     string  `json:"user_type" bson:"user_type"`
	Position     string  `json:"position" bson:"position"`
}
