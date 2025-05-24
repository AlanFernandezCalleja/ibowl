package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileContent struct {
	Nombre        string `json:"nombre"`
	Edad          int    `json:"edad"`
	Correo        string `json:"correo"`
	Password      string `json:"password"`
	FechaCreacion string `json:"fecha_creacion"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Pedir datos
	fmt.Print("Nombre: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Edad: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Edad inválida")
		return
	}

	fmt.Print("Correo: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Contraseña: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Fecha de creación
	today := time.Now().Format("2006-01-02 15:04:05")

	// Crear estructura
	fileData := FileContent{
		Nombre:        name,
		Edad:          age,
		Correo:        email,
		Password:      password,
		FechaCreacion: today,
	}

	// Verificar si el archivo ya existe
	fileName := "usuarios.json"
	var users []FileContent

	// Leer archivo si existe
	if _, err := os.Stat(fileName); err == nil {
		// Si existe, leer el archivo
		fileBytes, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error al leer el archivo:", err)
			return
		}

		// Deserializar el archivo JSON a una lista de usuarios
		err = json.Unmarshal(fileBytes, &users)
		if err != nil {
			fmt.Println("Error al deserializar el archivo:", err)
			return
		}
	}

	// Agregar el nuevo usuario a la lista
	users = append(users, fileData)

	// Convertir la lista de usuarios a JSON
	jsonBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Escribir los datos actualizados al archivo
	err = os.WriteFile(fileName, jsonBytes, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}

	fmt.Println("Archivo actualizado exitosamente como:", fileName, "en ", today)
}
