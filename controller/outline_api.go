package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      int    `json:"port"`
	Method    string `json:"method"`
	AccessURL string `json:"accessUrl"`
}

func CreateNewUser(name string) (string, bool) {
	s, err := GetUsers()
	if err != nil {
		log.Println(":::GetUsers:::\n" + err.Error())
		return "API call error, Please try again.", false
	}
	if strings.Contains(s, name) {
		return "User is already assigned.", false
	} else {

		user, err := RequestNewUser()
		if err != nil {
			log.Println(":::RequestNewUser:::\n" + err.Error())
			return "API call error, Please try again.", false
		} else {
			user.Name = name
			err = RenameUser(&user)
			if err != nil {
				log.Println(":::RenameUser:::\n" + err.Error())
				return "API call error, Please try again.", false
			}
			return user.AccessURL, true
		}

	}
}

func GetUsers() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, os.Getenv("SECRET_URL")+"/access-keys/", strings.NewReader(""))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

func RequestNewUser() (User, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, os.Getenv("SECRET_URL")+"/access-keys/", strings.NewReader(""))
	if err != nil {
		return User{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}
	var user User
	err = json.Unmarshal(body, &user)
	return user, err
}

func RenameUser(r *User) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, os.Getenv("SECRET_URL")+"/access-keys/"+r.ID+"/name", strings.NewReader(`{"name": "`+r.Name+`"}`))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
