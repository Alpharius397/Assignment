package main

import (
	"backend/utils"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	Message string `json:"message" default:"ok"`
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// Login godoc
// @Summary      Login API
// @Description  validate credentials and generate JWT tokens (access and refresh)
// @Accept       json
// @Produce      json
// @Param        user body UserLogin true "User Data"
// @Success      200  {object}  LoginSuccessResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /login [post]
func Login(g *gin.Context) {

	var userData UserLogin

	if err := g.ShouldBindJSON(&userData); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{Message: "Failed to parse body"})
		return
	}

	if DB == nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to establish connection to database"})
		return
	}

	query, err := DB.Prepare(`select ROWID, email, password from Users where user_name = ?`)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate database statement"})
		return
	}

	defer query.Close()

	var ROWID int
	var email string
	var password string

	err = query.QueryRow(userData.UserName).Scan(&ROWID, &email, &password)

	if err != nil {
		g.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Incorrect Username or Password found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userData.Password))

	if err != nil {
		g.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Incorrect Username or Password found"})
		return
	}

	access, err := utils.GetAccessToken(strconv.Itoa(ROWID), email)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create JWT token"})
		return
	}

	refresh, err := utils.GetRefreshToken(strconv.Itoa(ROWID), email)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create JWT token"})
		return
	}

	g.JSON(http.StatusOK, LoginSuccessResponse{Message: "ok", Access: access, Refresh: refresh})
}

type UserRegister struct {
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Aadhar          string `json:"aadhar"`
}

func (u *UserRegister) Validate() error {
	aadharRegex := regexp.MustCompile(`^\d{12}$`)

	if ok := aadharRegex.MatchString(u.Aadhar); !ok {
		return errors.New("invalid Aadhar Number found")
	}

	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return err
	}

	if u.ConfirmPassword != u.Password {
		return errors.New("password and confirm password must match")
	}

	return nil
}

type RegisterResponse struct {
	Message string `json:"message" default:"ok"`
}

// Register godoc
// @Summary      Register API
// @Description  Registers new user
// @Accept       json
// @Produce      json
// @Param        user body UserRegister true "User Data"
// @Success      200  {object}  RegisterResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /register [post]
func Register(g *gin.Context) {

	var userData UserRegister

	if err := g.ShouldBindJSON(&userData); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{Message: "Failed to parse body"})
		return
	}

	if err := userData.Validate(); err != nil {
		g.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("Invalid Data found. Error: %v", err)})
		return
	}

	if DB == nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to establish connection to database"})
		return
	}

	encryptedAadhar, err := utils.AesEncrypt([]byte(userData.Aadhar))

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to encrypted aadhar"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to encrypt password"})
		return
	}

	query, err := DB.Prepare(`insert into users(user_name, email, password, aadhar) values (?, ?, ?, ?)`)

	if err != nil {
		fmt.Println(err)
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate database statement"})
		return
	}

	defer query.Close()

	_, err = query.Exec(userData.UserName, userData.Email, string(hashedPassword), string(encryptedAadhar))

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to insert record into database"})
		return
	}

	g.JSON(http.StatusOK, RegisterResponse{Message: "ok"})
}

type RefreshResponse struct {
	Message string `json:"message" default:"ok"`
	Access  string `json:"access"`
}

// Refresh godoc
// @Summary      Refresh API
// @Description  Refresh user's access token
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Refresh Token"
// @Success      200  {object}  RefreshResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /refresh [post]
func Refresh(g *gin.Context) {
	authHeader := g.Request.Header.Get("Authorization")

	token, ok := strings.CutPrefix(authHeader, "Bearer ")

	if !ok {
		g.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Authorization is missing"})
		return
	}

	user, err := utils.ParseRefreshToken(token)

	if err != nil {
		g.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid JWT token received"})
		return
	}

	if DB == nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to establish connection to database"})
		return
	}

	query, err := DB.Prepare(`select user_name from Users where ROWID = ? and email = ?`)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate database statement"})
		return
	}

	defer query.Close()

	var user_name string

	err = query.QueryRow(user.UserId, user.Email).Scan(&user_name)

	if err != nil {
		g.JSON(http.StatusUnauthorized, ErrorResponse{Message: "UserID was not found"})
		return
	}

	access, err := utils.GetAccessToken(user.UserId, user.Email)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create JWT token"})
		return
	}

	g.JSON(http.StatusOK, RefreshResponse{Message: "ok", Access: access})
}

type ProfileResponse struct {
	Message  string `json:"message" default:"ok"`
	UserName string `json:"user_name"`
	UserId   int    `json:"id"`
	Email    string `json:"email"`
	Aadhar   string `json:"aadhar"`
}

// GetProfile godoc
// @Summary      Profile API
// @Description  Returns signed-in user's info
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Access Token"
// @Success      200  {object}  ProfileResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /profile [get]
func GetProfile(g *gin.Context) {
	_user, _ := g.Get("User")

	user, ok := _user.(AuthUser)

	if !ok {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "User was not found"})
		return
	}

	if DB == nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to establish connection to database"})
		return
	}

	userQuery, err := DB.Prepare(`select ROWID, aadhar from Users where ROWID = ?`)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate database statement"})
		return
	}

	defer userQuery.Close()

	var ROWID int
	var aadhar string

	err = userQuery.QueryRow(user.UserId).Scan(&ROWID, &aadhar)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch data from database"})
		return
	}

	decrypted, err := utils.AesDecrypt(aadhar)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to decrypt data"})
		return
	}

	g.JSON(http.StatusOK, ProfileResponse{Message: "ok", UserName: user.UserName, UserId: ROWID, Email: user.Email, Aadhar: decrypted})
}

type UsersList struct {
	ROWID    int    `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Aadhar   string `json:"aadhar"`
}

type DataResponse struct {
	Message string      `json:"message" default:"ok"`
	Data    []UsersList `json:"data"`
	Total   int         `json:"total"`
}

// GetData godoc
// @Summary      Data API
// @Description  Returns list of user info with pagination
// @Accept       json
// @Produce      json
// @Param        offset query number false "Offset"
// @Param        limit query number false "Limit"
// @Param        raw query boolean false "Raw"
// @Param        Authorization header string true "JWT Access Token"
// @Success      200  {object}  DataResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /get-data [get]
func GetData(g *gin.Context) {
	_user, _ := g.Get("User")

	_, ok := _user.(AuthUser)

	if !ok {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "User was not found"})
		return
	}

	if DB == nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to establish connection to database"})
		return
	}

	countQuery, err := DB.Prepare(`select count(ROWID) from users;`)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate database statement"})
		return
	}

	defer countQuery.Close()

	var totalCount int = 0

	err = countQuery.QueryRow().Scan(&totalCount)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch data from database"})
		return
	}

	userQuery, err := DB.Prepare(`select ROWID, user_name, email, aadhar from users limit ? offset ?`)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to generate database statement"})
		return
	}

	defer userQuery.Close()

	offset := utils.GetParams(g.Request.URL.Query(), "offset", 0)
	limit := utils.GetParams(g.Request.URL.Query(), "limit", 1)
	raw := !(g.Request.URL.Query().Get("raw") == "false")

	row, err := userQuery.Query(limit, offset)

	if err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch data from database"})
		return
	}

	defer row.Close()

	dataList := make([]UsersList, 0)

	for row.Next() {

		var ROWID int
		var userName string
		var email string
		var aadhar string

		if err = row.Err(); err != nil {
			g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch data from database"})
			return
		}

		if err = row.Scan(&ROWID, &userName, &email, &aadhar); err != nil {
			g.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch data from database"})
			return
		}

		if raw {
			dataList = append(dataList, UsersList{ROWID: ROWID, UserName: userName, Email: email, Aadhar: aadhar})

		} else {
			decrypted, err := utils.AesDecrypt(aadhar)

			if err != nil {
				continue
			}

			dataList = append(dataList, UsersList{ROWID: ROWID, UserName: userName, Email: email, Aadhar: decrypted})
		}
	}

	g.JSON(http.StatusOK, DataResponse{Message: "ok", Data: dataList, Total: totalCount})
}
