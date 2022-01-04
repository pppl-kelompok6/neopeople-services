package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neopeople-service/database"
	_helper "neopeople-service/helper_"
	"neopeople-service/model"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)

	var user model.Team
	var dbUser model.Team

	json.Unmarshal(requestBody, &user)

	// error from body json
	err := json.NewDecoder(r.Body).Decode(&user)
	if dbUser.Email != "" {
		fmt.Println("line 24")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	//  check if email exist
	database.Connector.Where("email = ?", user.Email).First(&dbUser)
	database.Connector.Where("username = ?", user.Username).First(&dbUser)

	if dbUser.Email != "" {
		fmt.Println("Email has already taken!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email has already taken!")
		return
	}
	if dbUser.Username != "" {
		fmt.Println("Username has already taken!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email has already taken!")
		return
	}

	// hashing password
	user.Password, err = _helper.GenerateHashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error while hashing Password")
		return
	}

	errDB := database.Connector.Create(&user).Error

	if errDB != nil {
		fmt.Println(errDB)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errDB)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var auth model.Authentication
	var user model.Team

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		fmt.Println("Something went wrong!", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid input!")
	}
	// fmt.Println(auth)
	database.Connector.Where("email = ?", auth.Email).First(&user)
	// fmt.Println(user.Email)
	if user.Email == "" {
		fmt.Println("Email not found!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username or Password is incorrect!")
		return
	}

	check := _helper.CheckPasswordHash(auth.Password, user.Password)

	if !check {
		fmt.Println("Username or Password is incorrect")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username or Password is incorrect!")
		return
	}

	validToken, err := _helper.GenerateJWT(auth.Email, user.Position)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Failed to generate token")
		return
	}

	var token model.Token

	token.Email = user.Email
	token.Role = user.Position
	token.TokenString = validToken

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
