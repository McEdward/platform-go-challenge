// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/frontend"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	log.Println("Hello, the application has started")
	viper.SetConfigFile("config.env")
	viper.ReadInConfig()
	r := mux.NewRouter()

	h := frontend.New(r)

	h.SetRoutes()

	//Use of dynamic port and ip address from a confic file
	//This allows the inspector to adjust it in case the port is being used by another program
	ip := fmt.Sprintf("%s", viper.Get("IP"))
	port := fmt.Sprintf("%s", viper.Get("PORT"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Println("Server running on " + ip + ":" + port)
	srv := &http.Server{
		Handler:      r,
		Addr:         ip + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

/*
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", FavoriteHandler)

}
*/

func FavoriteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End")
}
