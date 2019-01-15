package handler

import (
	"github.com/butlerwang/golang_yaml_rest_api_server/lib"
	"github.com/butlerwang/golang_yaml_rest_api_server/memdb"
	"github.com/butlerwang/golang_yaml_rest_api_server/model"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func CreateApp(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("X-token"); token != "123456" {
		http.Error(w, "Wrong token", http.StatusForbidden)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	app := new(model.App)

	err = yaml.Unmarshal(body, &app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := lib.StripSpace(app.Title)
	memdb.Save(id, app)

	w.Header().Set("Location", r.URL.Path+"/"+id)
	w.WriteHeader(http.StatusCreated)

	bytes, err := yaml.Marshal(app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeYamlResponse(w, bytes)
}

func GetApp(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("X-token"); token != "123456" {
		http.Error(w, "Wrong token", http.StatusForbidden)
	}

	vars := mux.Vars(r)
	id := lib.StripSpace(vars["name"])

	app, ok := memdb.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := yaml.Marshal(app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeYamlResponse(w, bytes)
}

func GetApps(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("X-token"); token != "123456" {
		http.Error(w, "Wrong token", http.StatusForbidden)
	}

	apps, ok := memdb.GetAll()
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := yaml.Marshal(apps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeYamlResponse(w, bytes)
}

func UpdateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := lib.StripSpace(vars["name"])

	_, ok := memdb.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	app := new(model.App)
	err = yaml.Unmarshal(body, &app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	memdb.Save(id, app)
}

func DeleteApp(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("X-token"); token != "123456" {
		http.Error(w, "Wrong token", http.StatusForbidden)
	}

	vars := mux.Vars(r)
	id := lib.StripSpace(vars["name"])

	_, ok := memdb.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	memdb.Remove(id)
	w.Write([]byte("OK"))
}

func SearchApp(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("X-token"); token != "123456" {
		http.Error(w, "Wrong token", http.StatusForbidden)
	}

	vars := mux.Vars(r)
	id := lib.StripSpace(vars["name"])

	apps, ok := memdb.Search(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := yaml.Marshal(apps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeYamlResponse(w, bytes)
}

func writeYamlResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "text/x-yaml; charset=UTF-8")
	w.Write(bytes)
}
