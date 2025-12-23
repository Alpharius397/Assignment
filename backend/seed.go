package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type SeedData struct {
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Aadhar          string `json:"aadhar"`
}

// generate random valid aadhar ID
func generateAadhar() string {
	b := make([]byte, 12)
	for i := range b {
		b[i] = byte('0' + rand.Intn(10))
	}
	return string(b)
}

// generate a sample password
func generatePassword(id int) string {
	return fmt.Sprintf("Pass%d", id)
}

// generate a sample user struct
func generateUser(id int) SeedData {
	password := generatePassword(id)
	first := "User"
	last := strconv.Itoa(id)

	return SeedData{
		UserName:        fmt.Sprintf("%s_%s", strings.ToLower(first), strings.ToLower(last)),
		Email:           fmt.Sprintf("%s%s@example.com", strings.ToLower(first), strings.ToLower(last)),
		Password:        password,
		Aadhar:          generateAadhar(),
	}
}

// generate a array of sample user
func generateUsers(count int) []SeedData {
	users := make([]SeedData, 0, count)
	for i := 1; i <= count; i++ {
		users = append(users, generateUser(i))
	}
	return users
}

// generates a list of 100 users and saves it to seed.json
func generateRandomUsers() error {

	users := generateUsers(100)

	out, err := json.MarshalIndent(users, "", "  ")

	if err != nil {
		return err
	}

	file, err := os.OpenFile("seed.json", os.O_CREATE | os.O_WRONLY | os.O_EXCL, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(out)

	if err != nil {
		return err
	}

	return nil
}
