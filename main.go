package main

import (
	"encoding/json" // Para codificar/decodificar JSON
	"fmt"           // Para imprimir texto en la respuesta
	"io"            // Para leer el cuerpo de la solicitud
	"log"           // Para registrar errores
	"net/http"      // Para manejar peticiones HTTP
	"strconv"       // Para convertir strings a enteros

	"github.com/gorilla/mux" // Router externo que permite trabajar con rutas con variables (como /task/{id})
)

// Estructura que representa una tarea
type task struct {
	ID      int    `json:"id"` // Etiquetas json: así es como se serializa a JSON
	Name    string `json:"name"`
	Content string `json:"content"`
}

var tasks = map[int]task{
	1: {ID: 1, Name: "Task One", Content: "Some Content"},
}

// Ruta de bienvenida
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome here!")
}

// Obtener todas las tareas
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result []task

	for _, t := range tasks {
		result = append(result, t)
	}

	json.NewEncoder(w).Encode(tasks)
}

// Crear una nueva tarea
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert a valid task", http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &newTask)

	newID := 1
	for id := range tasks {
		if id >= newID {
			newID = id + 1
		}
	}

	newTask.ID = newID
	tasks[newID] = newTask

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// Obtener una tarea por ID
func getTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)                     // Extrae los parámetros de la URL
	taskID, err := strconv.Atoi(vars["id"]) // Convierte el parámetro a entero

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, exists := tasks[taskID]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// Eliminar una tarea por ID
func deleteTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := tasks[taskID]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(tasks, taskID)
	fmt.Fprintf(w, "The task with ID %v has been removed successfully", taskID)
}

// Actualizar una tarea por ID
func updateTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := tasks[taskID]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updatedTask task
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Enter a valid task", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	updatedTask.ID = taskID
	tasks[taskID] = updatedTask

	fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
}

// Función principal: arranca el servidor y define las rutas
func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Rutas
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTaskById).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTaskById).Methods("DELETE")
	router.HandleFunc("/task/{id}", updateTaskById).Methods("PUT")

	// Arranca el servidor en el puerto 3000
	log.Fatal(http.ListenAndServe(":3000", router))
}
