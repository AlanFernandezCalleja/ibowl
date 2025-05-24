package models

type Mascota struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Nombre      string  `gorm:"size:100" json:"nombre"`
	Raza        string  `gorm:"size:100" json:"raza"`
	Edad        int     `gorm:"type:int" json:"edad"`
	Peso        float64 `gorm:"type:double" json:"peso"`
	Frecuencia  string  `gorm:"size:100" json:"frecuencia"`
	UsuarioID   uint    `gorm:"column:id_usuario" json:"id_usuario"`
	Usuario     Usuario `gorm:"foreignKey:UsuarioID;references:ID" json:"usuario"`
}

func (Mascota) TableName() string {
	return "mascotas"
}