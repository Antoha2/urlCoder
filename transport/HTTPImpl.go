package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	service "github.com/antoha2/urlCoder/service"
	"github.com/gin-gonic/gin"
)

func (wImpl *webImpl) StartHTTP() error {

	router := gin.Default()
	router.POST("/api/:token", wImpl.redirectHandler)
	router.POST("/add", wImpl.handlerAddLongUrl)
	router.POST("/gen", wImpl.handlerGenTokens)
	router.Run()
	return nil
}

//log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!2")
// wImpl.server = &http.Server{Addr: ":8180"}
// fmt.Println("Server is listening :8180 ...")
// wImpl.server.ListenAndServe()

// http.HandleFunc("/", wImpl.handlerGenTokens)

//r := mux.NewRouter() //I'm using Gorilla Mux, but it could be any other library, or even the stdlib

//r.Methods("POST").Path("/create").Handler(createHandler)
// r.Methods("POST").Path("/auth/sign-up/admin").Handler(signUpAdminHandler)
// r.Methods("POST").Path("/auth/sign-up/user").Handler(signUpUserHandler)
// r.Methods("POST").Path("/auth/deleteUser").Handler(deleteUserHandler)
// r.Methods("POST").Path("/auth/updateUser").Handler(updateUserHandler)

////////////////////////////////////////
// http.HandleFunc("/api", wImpl.redirectHandler)
// http.HandleFunc("/add", wImpl.handlerAddLongUrl)
// http.HandleFunc("/genTokens", wImpl.handlerGenTokens)
// ///////////////////////////////////
// wImpl.server = &http.Server{Addr: ":8180"}
// fmt.Println("Server is listening :8180 ...")
// wImpl.server.ListenAndServe()
/////////////////////////////////////

func (wImpl *webImpl) redirectHandler(c *gin.Context) {
	url := new(service.ServUrl)
	url.Token = c.Param("token")
	err := wImpl.service.ServRedirect(url)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, url.Long_url)
}

func (wImpl *webImpl) handlerAddLongUrl(c *gin.Context) {

	url := new(service.ServUrl)
	//url.Long_url=c.Param("long_url")
	if err := c.BindJSON(&url); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err := wImpl.service.AddLongUrl(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, url.Token)

}

func (wImpl *webImpl) handlerGenTokens(c *gin.Context) {

	err := wImpl.service.ServGenTokens()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "генерация прошла успешно")

}

/*
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

	err := wImpl.service.ServGenTokens()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	str := "генерация прошла успешно"
	json, err := json.Marshal(str)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)

}
*/

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
