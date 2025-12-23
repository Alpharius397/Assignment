package main

import (
	"backend/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Init script to setup sqlite3 db and insert required tables, relies on an init.sql script
func InitDB(initSqlPath string, dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return err
	}

	defer db.Close()

	initSql, err := os.Open(initSqlPath)

	if err != nil {
		return err
	}

	defer initSql.Close()

	queries, err := io.ReadAll(initSql)

	if err != nil {
		return err
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(string(queries)); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

// Opens connection to sqlite3 db
func GetDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Seed db with sample records for testing purpose using seed.json
func SeedDb() error {

	if(DB == nil) {
		return errors.New("failed to establish connection to database")
	}

	seedFile, err := os.Open(SeedJson)

	if err != nil {
		return err
	}

	defer seedFile.Close()

	data, err := io.ReadAll(seedFile)

	if err != nil {
		return err
	}

	var users []UserRegister

	err = json.Unmarshal(data, &users)

	if err != nil {
		return err
	}

	mainQuery := `insert into users(user_name, email, password, aadhar) values `

	var values []any
	var placeholders []string

	for _, u := range users {
		encryptedAadhar, err := utils.AesEncrypt([]byte(u.Aadhar))

		if err != nil {
			continue
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

		if err != nil {
			continue
		}

		placeholders = append(placeholders,
			"(? , ?, ?, ?)")

		values = append(values,
			u.UserName,
			u.Email,
			hashedPassword,
			encryptedAadhar,
		)
	}

	tx, err := DB.Begin()

	if err != nil {
		return err
	}

	sqlQuery, err := tx.Prepare(mainQuery + strings.Join(placeholders, ","))

	if err != nil {
		return err
	}

	defer sqlQuery.Close()

	if _, err := sqlQuery.Exec(values...); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
