package ds

import (
	// "gorm.io/gorm"
	"time"
)

const DRAFT string = "черновик"
const DELETED string = "удален"
const FORMED string = "сформирован"
const COMPLETED string = "завершен"
const REJECTED string = "отклонен"

type User struct {
	UUID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string `gorm:"size:30;not null" json:"name"`
	Login       string `gorm:"size:30;not null"`
	Password    string `gorm:"size:30;not null"`
	IsModerator bool   `gorm:"not null;default:false"`
}

type Cargo struct {
	UUID        string  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid"`
	Name        string  `gorm:"size:100;not null" form:"name" json:"name" binding:"required"`
	EnglishName string  `gorm:"size:100;not null" form:"en_name" json:"en_name" binding:"required"`
	Photo       *string `gorm:"size:100" json:"photo"`
	Category    string  `gorm:"size:50;not null" form:"category" json:"category" binding:"required"`
	Price       uint    `gorm:"not null" form:"price" json:"price" binding:"required"`       // Rubles
	Weight      float64 `gorm:"not null" form:"weight" json:"weight" binding:"required"`     // kg
	Capacity    float64 `gorm:"not null" form:"capacity" json:"capacity" binding:"required"` // m^3
	Description string  `gorm:"size:500;not null" form:"description" json:"description" binding:"required"`
	IsDeleted   bool    `gorm:"not null;default:false"`
}

type Flight struct {
	UUID           string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Status         string     `gorm:"size:20;not null"`
	CreationDate   time.Time  `gorm:"not null;type:timestamp"`
	FormationDate  *time.Time `gorm:"type:timestamp"`
	CompletionDate *time.Time `gorm:"type:timestamp"`
	ClientId       string     `gorm:"not null"`
	ModeratorId    *string    `json:"-"`
	RocketType     string     `gorm:"size:50"`

	Client    User
	Moderator *User
}

type FlightCargo struct {
	FlightId string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"flight_id"`
	CargoId  string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"cargo_id"`
	Quantity uint   `gorm:"not null" json:"quantity"`

	Flight *Flight `gorm:"foreignKey:FlightId" json:"flight"`
	Cargo  *Cargo  `gorm:"foreignKey:CargoId" json:"cargo"`
}
