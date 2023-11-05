package ds

import (
	// "gorm.io/gorm"
	"time"
)

type User struct {
	UserId      uint   `gorm:"primarykey"`
	Name        string `gorm:"size:30;not null"`
	Login       string `gorm:"size:30;not null"`
	Password    string `gorm:"size:30;not null"`
	IsModerator bool   `gorm:"not null"`
}

type Cargo struct {
	CargoId     uint    `gorm:"primaryKey;not null;autoIncrement:false"`
	Name        string  `gorm:"size:100;not null"`
	EnglishName string  `gorm:"size:100;not null"`
	Photo       string  `gorm:"size:100;not null"`
	Category    string  `gorm:"size:50;not null"`
	Price       uint    `gorm:"not null"` // Rubles
	Weight      float32 `gorm:"not null"` // kg
	Capacity    float32 `gorm:"not null"` // m^3
	Description string  `gorm:"size:500;not null"`
	IsDeleted   bool    `gorm:"not null"`
}

type Flight struct {
	FlightId       uint       `gorm:"primaryKey"`
	Status         string     `gorm:"size:50;not null"`
	CreationDate   time.Time  `gorm:"not null;type:timestamp"`
	FormationDate  *time.Time `gorm:"type:timestamp"`
	CompletionDate *time.Time `gorm:"type:timestamp"`
	ClientId       uint       `gorm:"not null"`
	ModeratorId    *uint
	RocketType     *string `gorm:"size:50"`

	Client    User `gorm:"foreignKey:ClientId"`
	Moderator User `gorm:"foreignKey:ModeratorId"`
}

type FlightCargo struct {
	FlightId uint `gorm:"primaryKey;not null;autoIncrement:false"`
	CargoId  uint `gorm:"primaryKey;not null;autoIncrement:false"`
	Quantity uint `gorm:"not null"`

	Flight *Flight `gorm:"foreignKey:FlightId"`
	Cargo  *Cargo  `gorm:"foreignKey:CargoId"`
}
