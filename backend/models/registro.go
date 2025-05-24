package models

import "time"

type Registro struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Distancia  float64   `gorm:"type:float" json:"distancia"`
	Movimiento bool      `gorm:"type:tinyint(1)" json:"movimiento"`
	Fecha      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"fecha"`
}

func (Registro) TableName() string {
	return "registros"
}