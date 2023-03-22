package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	var Person Person
	err := json.NewDecoder(r.Body).Decode(&Person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	db, err := Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE tb_crud SET name = ?, phone = ? WHERE id = ? ", Person.Name, Person.Phone, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	db, err := Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	_, err = db.Exec("DELETE FROM tb_crud WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success"))
}


func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	var Person Person
	err := json.NewDecoder(r.Body).Decode(&Person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tb_crud (name, phone) VALUES (?, ?)", Person.Name, Person.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success"))
}


func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	db, err := Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	data, err := db.Query("SELECT * FROM tb_crud")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result []Person

	for data.Next() {
		var each = Person{}
		var err = data.Scan(&each.Id, &each.Name, &each.Phone)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, each)
	}

	hasil, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	db, err := Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var result = Person{}
	err = db.QueryRow("SELECT * FROM tb_crud WHERE id = (?)", id).Scan(&result.Id, &result.Name, &result.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hasil, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}