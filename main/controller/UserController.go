package controller

import (
	"encoding/json"
	"fmt"
	"localhost/shbh/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.GetAll(w, r)
		case http.MethodPost:
			uc.AddUser(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		Id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.GetById(Id, w)
		case http.MethodPut:
			uc.UpdateUser(Id, w, r)
		case http.MethodDelete:
			uc.DeleteUser(Id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc UserController) GetById(Id int, w http.ResponseWriter) {
	u, err := models.GetUserById(Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

// this method called as post in request
func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User Object"))
		return
	}
	fmt.Println(u)
	u, err = models.AddUser(u)
	fmt.Println(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc UserController) UpdateUser(Id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User Object"))
		return
	}
	if u.Id != Id {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id Of Object should match the Id of Object"))
	}
	err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	encodeResponseAsJSON(u, w)
}

func (us UserController) DeleteUser(Id int, w http.ResponseWriter) {
	err := models.DeleteUserById(Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc UserController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	println(dec)
	var u models.User
	fmt.Println(u)
	err := dec.Decode(&u)
	fmt.Println(u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
