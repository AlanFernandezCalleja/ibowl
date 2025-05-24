package models

type Configuracion struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	DispositivoID  string     `gorm:"column:dispositivo_id;size:50" json:"dispositivo_id"`
	Dispositivo    *Dispositivo `gorm:"foreignKey:DispositivoID" json:"dispositivo"` // Cambiado a puntero
	HoraProgramada string     `gorm:"column:hora_programada;type:time" json:"hora_programada"`
	PorcionGramos  float64    `gorm:"column:porcion_gramos" json:"porcion_gramos"`
	Activo         bool       `gorm:"column:activo;default:true" json:"activo"`
}

func (Configuracion) TableName() string {
	return "configuraciones"
}