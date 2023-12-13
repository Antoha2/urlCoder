package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	service "github.com/antoha2/urlCoder/service"
)

func (wImpl *webImpl) StartHTTP() error {

	http.HandleFunc("/create", wImpl.handlerLongUrl)

	//r := mux.NewRouter() //I'm using Gorilla Mux, but it could be any other library, or even the stdlib

	//r.Methods("POST").Path("/create").Handler(createHandler)
	// r.Methods("POST").Path("/auth/sign-up/admin").Handler(signUpAdminHandler)
	// r.Methods("POST").Path("/auth/sign-up/user").Handler(signUpUserHandler)
	// r.Methods("POST").Path("/auth/deleteUser").Handler(deleteUserHandler)
	// r.Methods("POST").Path("/auth/updateUser").Handler(updateUserHandler)

	wImpl.server = &http.Server{Addr: ":8180"}
	fmt.Println("Server is listening :8180 ...")
	wImpl.server.ListenAndServe()

	// wImpl.server = &http.Server{Addr: config.HTTPAddr}
	// log.Printf(" Запуск HTTP-сервера на http://127.0.0.1%s\n", wImpl.server.Addr) //:8180

	// if err := http.ListenAndServe(wImpl.server.Addr, r); err != nil {
	// 	log.Println(err)
	// }

	return nil
}

func (wImpl *webImpl) handlerLongUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	url := new(service.ServUrl)
	err := wImpl.Decoder(r, url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//fmt.Println("web ", url)

	err = wImpl.service.LongUrl(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	str := fmt.Sprintf("выполнено id-(%v) url-(%v) ", url.Id, url.Long_url)
	json, err := json.Marshal(str)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)

}

//декодеры JSON
func (wImpl *webImpl) Decoder(r *http.Request, unit *service.ServUrl) error { //unit *service.Service

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, unit)
	if err != nil {
		fmt.Println("can't unmarshal !!!!!: ", err.Error())
		return err
	}
	return nil
}
