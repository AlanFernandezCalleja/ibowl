package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/models" // Importamos todos los modelos

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	// Conectar a MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a MySQL:", err)
	}

	// Migrar todos los modelos
	if err := migrateModels(); err != nil {
		log.Fatal("Error en migración:", err)
	}

	// Configurar rutas
	r := mux.NewRouter()

	// Rutas para Mascotas
	r.HandleFunc("/mascotas", crearMascota).Methods("POST")
	//r.HandleFunc("/mascotas", listarMascotas).Methods("GET")
	//r.HandleFunc("/mascotas/{id}", obtenerMascota).Methods("GET")
	//r.HandleFunc("/mascotas/{id}", actualizarMascota).Methods("PUT")
	//r.HandleFunc("/mascotas/{id}", eliminarMascota).Methods("DELETE")

	// Rutas para Usuarios
	r.HandleFunc("/usuarios", crearUsuario).Methods("POST")
	r.HandleFunc("/usuarios", listarUsuarios).Methods("GET")

	// Rutas para Alimentación
	r.HandleFunc("/alimentacion", registrarAlimentacion).Methods("POST")
	r.HandleFunc("/alimentacion/{dispositivo_id}", obtenerRegistrosAlimentacion).Methods("GET")

	// Iniciar servidor
	port := os.Getenv("SERVER_PORT")
	log.Printf("API corriendo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// migrateModels realiza la migración de todos los modelos
// func migrateModels() error {
// 	modelos := []interface{}{
// 		&models.Usuario{},
// 		&models.Mascota{},
// 		&models.Dispositivo{},
// 		&models.Alimentacion{},
// 		&models.Configuracion{},
// 		&models.Log{},
// 		&models.Registro{},
// 	}

// 	return db.AutoMigrate(modelos...)
// }





// Handlers para Mascotas (similar a los anteriores, pero usando models.Mascota)
func crearMascota(w http.ResponseWriter, r *http.Request) {
	var mascota models.Mascota
	if err := json.NewDecoder(r.Body).Decode(&mascota); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&mascota)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mascota)
}






// Handlers para Usuarios
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash de contraseña debería ir aquí antes de crear
	result := db.Create(&usuario)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// No devolver la contraseña en la respuesta
	usuario.Contrasena = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	result := db.Select("id, nombre, email, rol").Find(&usuarios) // Excluir contraseñas
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// Handlers para Alimentación
func registrarAlimentacion(w http.ResponseWriter, r *http.Request) {
	var registro models.Alimentacion
	if err := json.NewDecoder(r.Body).Decode(&registro); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&registro)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registro)
}

func obtenerRegistrosAlimentacion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dispositivoID := vars["dispositivo_id"]

	var registros []models.Alimentacion
	result := db.Where("dispositivo_id = ?", dispositivoID).Find(&registros)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registros)
}
