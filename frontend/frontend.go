package frontend

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//This data type is a conbination of all the different asset types that can be used and the
//data should be processed in accordance to it's type ideally
type datast struct {
	//charts
	XValues []float64
	YValues []float64

	//audience
	gender        string
	birth_country string
	age_group     string
	daily_hours   float32
	purchases     int64
}

type favorite struct {
	Name string
	Desc string
	Data datast
	Type string
}

type Handler struct {
	*mux.Router
}

type data struct {
	data []favorite
}

var datal []favorite
var favs data

//Test input data
var test1 = map[string]favorite{
	"1st asset": {
		Name: "1st asset",
		Desc: "40% of millenials spend more than 3hours on social media daily",
		Data: datast{
			XValues: []float64{},
			YValues: []float64{},
		},
		Type: "insight",
	},
	"2nd asset": {
		Name: "2nd asset",
		Desc: "40% of millenials spend more than 3hours on social media daily",
		Data: datast{
			XValues:       []float64{},
			YValues:       []float64{},
			gender:        "male",
			birth_country: "Poland",
			age_group:     "18-32",
			daily_hours:   10,
			purchases:     1,
		},
		Type: "audience",
	},
}

func New(r *mux.Router) *Handler {

	return &Handler{
		Router: r,
	}
}

func (h *Handler) SetRoutes() {
	h.HandleFunc("/", h.favorites).Methods("GET")
	h.HandleFunc("/add", h.addFavorite).Methods("POST")
	h.HandleFunc("/update/{assetName}", h.addFavorite).Methods("POST")
	h.HandleFunc("/delete/{assetName}", h.delFavorite).Methods("GET")
}

func (h *Handler) favorites(w http.ResponseWriter, r *http.Request) {
	log.Println(test1)
	jsn, err := json.Marshal(test1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(jsn))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
func (h *Handler) addFavorite(w http.ResponseWriter, r *http.Request) {
	type res struct {
		Status  string
		Message string
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var fav favorite
	json.Unmarshal(reqBody, &fav)

	//This section could be easily expanded to cover all the types and make sure they are always in the correct format
	if fav.Name == "" || fav.Type == "" {
		err_msg := res{
			Status:  "Failed",
			Message: "Content Format error",
		}
		log.Println("Updated results:", test1)
		jsn, err := json.Marshal(err_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
		return
	}

	log.Println(fav)

	//Ensure the new favourite name does not already exit
	_, ok := test1[fav.Name]

	log.Println(fav.Name)
	// If the key exists
	if ok {
		err_msg := res{
			Status:  "Failed",
			Message: "Asset name already exist",
		}
		log.Println("Updated results:", test1)
		jsn, err := json.Marshal(err_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	} else {
		test1[fav.Name] = fav
		log.Println("Updated results:", test1)
		success_msg := res{
			Status:  "Success",
			Message: "Asset added successfully",
		}

		jsn, err := json.Marshal(success_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	}
}

func (h *Handler) updateFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	type res struct {
		Status  string
		Message string
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var fav favorite
	json.Unmarshal(reqBody, &fav)

	log.Println(fav)

	//This section could be easily expanded to cover all the types and make sure they are always in the correct format
	if fav.Name == "" || fav.Type == "" {
		err_msg := res{
			Status:  "Failed",
			Message: "Content Format error",
		}
		log.Println("Updated results:", test1)
		jsn, err := json.Marshal(err_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
		return
	}

	//Ensure the new favourite name does not already exit
	_, ok := test1[vars["assetName"]]

	log.Println(vars["assetName"])
	// If the key exists
	if !ok {
		err_msg := res{
			Status:  "Failed",
			Message: "Asset name does not exist",
		}
		log.Println("Updated results:", test1)
		jsn, err := json.Marshal(err_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	} else {
		delete(test1, fav.Name)
		test1[fav.Name] = fav
		log.Println("Updated results:", test1)
		success_msg := res{
			Status:  "Success",
			Message: "Asset added successfully",
		}

		jsn, err := json.Marshal(success_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	}
}

func (h *Handler) delFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	type res struct {
		Status  string
		Message string
	}

	_, ok := test1[vars["assetName"]]

	log.Println(vars["assetName"])
	// If the key exists
	if !ok {
		err_msg := res{
			Status:  "Failed",
			Message: "Asset name does not exist",
		}
		log.Println("Updated results:", test1)
		jsn, err := json.Marshal(err_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	} else {
		name := vars["assetName"]
		delete(test1, name)
		log.Println("Updated results:", test1)
		success_msg := res{
			Status:  "Success",
			Message: "Asset deleted successfully",
		}

		jsn, err := json.Marshal(success_msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(jsn))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	}
}
