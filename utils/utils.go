package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"strings"
	"time"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	env := os.Getenv("ENV")
	fmt.Println(env)

	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	fmt.Println(".env." + env + ".local")

	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	fmt.Println(".env." + env)

	err := godotenv.Load()
	HandleErr(err)
}

func Validation(values []models.Validation) bool {
	username := regexp.MustCompile("^([A-Za-z0-9]{5,})+$")
	email := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")
	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}

func PrepareToken(user *models.User, expiry int64) string {
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		fmt.Println(exists)
	}
	tokenContent := jwt.MapClaims{
		"user_id":   user.ID,
		"user_uuid": user.Uuid,
		"expiry":    expiry,
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(jwtKey))
	HandleErr(err)
	return token
}

func ValidateRequestToken(c *fiber.Ctx) bool {
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		fmt.Println(exists)
		return false
	}

	jwtToken := c.Get("Authorization")
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	cleanJWTHeader := strings.Split(cleanJWT, ".")[0]
	cleanJWTPayload := strings.Split(cleanJWT, ".")[1]
	cleanJWTSecret := strings.Split(cleanJWT, ".")[2]
	_, err := jwt.DecodeSegment(cleanJWTHeader)
	if err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			panic("\nbase64 input is corrupt, check service Key")
		}
		panic(err)
	}
	_, err = jwt.DecodeSegment(cleanJWTPayload)
	if err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			panic("\nbase64 input is corrupt, check service Key")
		}
		panic(err)
	}
	_, err = jwt.DecodeSegment(cleanJWTSecret)
	fmt.Println(err)
	if err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			panic("\nbase64 input is corrupt, check service Key")
		}
		panic(err)
	}

	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	HandleErr(err)
	//HandleErrRequest(err)

	now := time.Now()
	expiry := tokenData["expiry"].(float64)

	expired := now.After(time.Unix(int64(expiry), 0))
	if expired {
		return false
	}
	if !token.Valid {
		return false
	}

	return true
}
