package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Estructuras para los datos
type Mascota struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Raza          string  `json:"raza"`
	Edad          int     `json:"edad"`
	Peso          float64 `json:"peso"`
	Frecuencia    string  `json:"frecuencia"`
	IDUsuario     int     `json:"id_usuario"`
	NombreUsuario string  `json:"nombre_usuario"`
}

func main() {
	// Configuración de la conexión a la base de datos
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ibowl_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexión exitosa a la base de datos")

	// Manejador HTTP para obtener las mascotas
	http.HandleFunc("/mascotas", func(w http.ResponseWriter, r *http.Request) {
		// Configuración CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Manejar preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Resto de tu código original...
		query := `SELECT m.id, m.nombre, m.raza, m.edad, m.peso, m.frecuencia, m.id_usuario, u.nombre as nombre_usuario
              FROM mascotas m JOIN usuarios u ON m.id_usuario = u.id`

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var mascotas []Mascota

		// Iterar sobre los resultados
		for rows.Next() {
			var m Mascota
			err := rows.Scan(&m.ID, &m.Nombre, &m.Raza, &m.Edad, &m.Peso, &m.Frecuencia, &m.IDUsuario, &m.NombreUsuario)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			mascotas = append(mascotas, m)
		}

		// Verificar errores después de iterar
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Configurar el encabezado para JSON
		w.Header().Set("Content-Type", "application/json")

		// Codificar a JSON y enviar la respuesta
		json.NewEncoder(w).Encode(mascotas)
	})

	// Iniciar el servidor
	fmt.Println("Servidor escuchando en http://localhost:8080/mascotas")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
