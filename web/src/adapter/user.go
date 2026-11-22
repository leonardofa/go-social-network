package adapter

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"web/src/response"
)

// CreateUser call user service to create new user
func CreateUser(w http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	user, err := json.Marshal(map[string]string{
		"name":     request.FormValue("name"),
		"nick":     request.FormValue("nick"),
		"email":    request.FormValue("email"),
		"password": request.FormValue("password"),
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.ERROR(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)

}
