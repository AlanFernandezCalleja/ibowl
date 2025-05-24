package models

type Dispositivo struct {
	ID            string          `gorm:"primaryKey;column:dispositivo_id;size:50" json:"dispositivo_id"`
	Nombre        string          `gorm:"size:100" json:"nombre"`
	Ubicacion     string          `gorm:"size:100" json:"ubicacion"`
	Alimentaciones []Alimentacion `gorm:"foreignKey:DispositivoID" json:"alimentaciones"`
	Configuracion  *Configuracion `gorm:"foreignKey:DispositivoID" json:"configuracion"` // Cambiado a puntero
	Logs          []Log           `gorm:"foreignKey:DispositivoID" json:"logs"`
}

func (Dispositivo) TableName() string {
	return "dispositivos"
}