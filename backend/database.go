package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"backend/models" // Asegúrate que este import coincida con tu módulo

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB es la conexión global a la base de datos
var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() error {
	// Configuración de la conexión MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Configuración de GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Habilita logs SQL en desarrollo
	}

	// Establecer conexión
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return fmt.Errorf("error al conectar a MySQL: %w", err)
	}

	// Configurar conexión subyacente
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error al obtener DB() de GORM: %w", err)
	}

	// Configuración del pool de conexiones
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrar modelos
	if err := migrateModels(); err != nil {
		return fmt.Errorf("error en migración: %w", err)
	}

	log.Println("✅ Conexión a MySQL establecida correctamente")
	return nil
}

// migrateModels maneja la migración de todos los modelos
func migrateModels() error {
	modelos := []interface{}{
		&models.Usuario{},
		&models.Mascota{},
		&models.Dispositivo{},
		&models.Alimentacion{},
		&models.Configuracion{},
		&models.Log{},
		&models.Registro{},
	}

	// Deshabilitar claves foráneas temporalmente para evitar problemas
	if err := DB.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		return err
	}

	// Migrar cada modelo
	for _, modelo := range modelos {
		if err := DB.AutoMigrate(modelo); err != nil {
			return err
		}
	}

	// Reactivar claves foráneas
	if err := DB.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
		return err
	}

	return nil
}
