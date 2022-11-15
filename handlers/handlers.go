package handlers

import (
	"API/data"
	"API/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welecome to GOLANG API")
}

// Post
func PostRoute(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	reqBody, err := ioutil.ReadAll(r.Body)

	//Error case
	if err != nil {
		fmt.Fprintf(w, "Inserted invalid data")
	}

	//Read of JSON
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(data.TasksData)
	data.TasksData = append(data.TasksData, newTask)

	//Return
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// Get
func GetRoute(w http.ResponseWriter, r *http.Request) {
	//Return
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.TasksData)
}

// Put
func PutRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var updatedTask models.Task

	taskID, err := strconv.Atoi(vars["id"])

	//Error case
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	//Data error case
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}

	//Read json
	json.Unmarshal(reqBody, &updatedTask)

	//Return
	for i, t := range data.TasksData {
		if t.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)

			updatedTask.ID = t.ID
			data.TasksData = append(data.TasksData, updatedTask)

			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
		}
	}
}

// Delete
func DeleteRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	//Error case
	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	//Return
	for i, t := range data.TasksData {
		if t.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskID)
		}
	}
}
