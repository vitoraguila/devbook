package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))
}

func FindAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding all users!"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
