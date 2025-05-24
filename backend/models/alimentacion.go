package models

import "time"

type Alimentacion struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	NivelComida       float64    `gorm:"column:nivel_comida;not null" json:"nivel_comida"`
	PesoBowl          float64    `gorm:"column:peso_bowl;not null" json:"peso_bowl"`
	MovimientoDetectado bool     `gorm:"column:movimiento_detectado;not null" json:"movimiento_detectado"`
	FechaHora         time.Time  `gorm:"column:fecha_hora;default:CURRENT_TIMESTAMP" json:"fecha_hora"`
	DispositivoID     string     `gorm:"column:dispositivo_id;size:50;not null" json:"dispositivo_id"`
	Dispositivo       Dispositivo `gorm:"foreignKey:DispositivoID;references:ID" json:"dispositivo"`
}

func (Alimentacion) TableName() string {
	return "alimentacion"
}