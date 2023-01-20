package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Title    string `json:"Title" validate:"required"`
	SubTitle string `json:"SubTitle" validate:"required"`
	Text     string `json:"Text" validate:"required"`
}

type ServiceAgent struct {
	gorm.Model
	ID     string `gorm:"primary_key"`
	Name   string
	Status string
}

type Slot struct {
	gorm.Model
	ID             string `gorm:"primary_key"`
	Time           time.Time
	Quantity       uint
	Available      string
	ServiceAgent   string
	ServiceAgentID ServiceAgent `gorm:"foreignKey:ServiceAgent; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Ticket struct {
	gorm.Model
	ID             string `gorm:"primary_key"`
	TicketName     string `json:"ticket_name"`
	Description    string
	Type           string
	StartTime      time.Time
	EndTime        time.Time
	Quantity       uint
	Price          float32
	ServiceAgent   string       `json:"service_agent"`
	ServiceAgentID ServiceAgent `gorm:"foreignKey:ServiceAgent; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Slot           string
	// SlotID         Slot `gorm:"foreignKey:Slot; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Order struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Name     string
	Email    string
	WA       string
	Ticket   string
	TicketID Ticket `gorm:"foreignKey:Ticket; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
