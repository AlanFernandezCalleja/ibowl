package models

type Usuario struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Nombre     string   `gorm:"size:100;not null" json:"nombre"`
	Email      string   `gorm:"size:100;not null;unique" json:"email"`
	Contrasena string   `gorm:"size:255;not null" json:"-"`
	Rol        string   `gorm:"type:enum('admin','usuario');default:'usuario'" json:"rol"`
	Mascotas   []Mascota `gorm:"foreignKey:UsuarioID" json:"mascotas"`
}

func (Usuario) TableName() string {
	return "usuarios"
}