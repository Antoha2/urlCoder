package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	service "github.com/antoha2/urlCoder/service"
)

func (wImpl *webImpl) StartHTTP() error {

	http.HandleFunc("/add", wImpl.handlerAddLongUrl)
	http.HandleFunc("/genTokens", wImpl.handlerGenTokens)
	// http.HandleFunc("/", wImpl.handlerGenTokens)

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

func (wImpl *webImpl) handlerAddLongUrl(w http.ResponseWriter, r *http.Request) {
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

	err = wImpl.service.AddLongUrl(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	str := fmt.Sprintf("id-(%v) url-(%v) , token-(%v)", url.Id, url.Long_url, url.Token)
	json, err := json.Marshal(str)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)

}

func (wImpl *webImpl) handlerGenTokens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	//Q := new(service.Quantity)
	// err := wImpl.DecoderQ(r, Q)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	//q := Q.Q
	err := wImpl.service.ServGenTokens()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	str := fmt.Sprintf("генерация прошла успешно")
	json, err := json.Marshal(str)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)

}

//декодеры JSON
func (wImpl *webImpl) Decoder(r *http.Request, url *service.ServUrl) error { //unit *service.Service

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, url)
	if err != nil {
		fmt.Println("can't unmarshal !!!!!: ", err.Error())
		return err
	}
	return nil
}

// func (wImpl *webImpl) DecoderQ(r *http.Request) error { //unit *service.Service

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(body, q)
// 	if err != nil {
// 		fmt.Println("can't unmarshal !!!!!: ", err.Error())
// 		return err
// 	}
// 	return nil
// }
