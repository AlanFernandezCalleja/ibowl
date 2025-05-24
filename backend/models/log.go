package models

import "time"

type Log struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TipoEvento    string    `gorm:"column:tipo_evento;size:50" json:"tipo_evento"`
	Descripcion   string    `gorm:"column:descripcion;type:text" json:"descripcion"`
	FechaHora     time.Time `gorm:"column:fecha_hora;default:CURRENT_TIMESTAMP" json:"fecha_hora"`
	DispositivoID string    `gorm:"column:dispositivo_id;size:50" json:"dispositivo_id"`
	Dispositivo   Dispositivo `gorm:"foreignKey:DispositivoID;references:ID" json:"dispositivo"`
}

func (Log) TableName() string {
	return "logs"
}