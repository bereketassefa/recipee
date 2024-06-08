package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"example.com/m/model"
	"golang.org/x/crypto/bcrypt"

	"context"

	"example.com/m/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/machinebox/graphql"
	// "example.com/m/model"
	// "github.com/machinebox/graphql"
)

func Signup(c *gin.Context) {
	//get the email and pass of the req body
	var body struct {
		Input struct {
			FirstName string
			LastName  string
			Email     string
			Password  string
		}
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	if body.Input.Email == "" || body.Input.Password == "" || body.Input.FirstName == "" || body.Input.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please fill all the fields",
		})
		return
	}
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Input.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	fmt.Println(body.Input.Email, body.Input.Password, body.Input.FirstName, body.Input.LastName)
	//create  the user
	secondQuery := `mutation addUser($firstName: String! , $lastName: String! , $password: String! , $email: String!){
		insert_users_one(object: {firstName: $firstName, lastName: $lastName, password: $password , email: $email}) {
			id
		}
	  }
	`

	fmt.Println("=====> ", secondQuery)
	secondRequest := graphql.NewRequest(secondQuery)
	secondRequest.Var("firstName", body.Input.FirstName)
	secondRequest.Var("lastName", body.Input.LastName)
	secondRequest.Var("email", body.Input.Email)
	secondRequest.Var("password", string(hash))
	secondRequest.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	var SecondRespose model.CreatedUser
	requestError := initializers.CLIENT.Run(context.Background(), secondRequest, &SecondRespose)
	// err := initializers.CLIENT.Run(context.Background(), secondRequest, &SecondRespose)

	if requestError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": requestError.Error(),
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
func Login(c *gin.Context) {
	//get the email and pass of the req body
	var body struct {
		Input struct {
			Email    string
			Password string
		}
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body",
			"token":   "",
			"id":      "",
		})
		return
	}
	fmt.Println("req body", body.Input.Email, body.Input.Password)

	// create  the user
	fetchQuery := `
	query fetchUser ($email: String!) {
		users(where: {email: {_eq: $email}}) {
		  
		  password
		  id
		}
	  }

	`
	userRequest := graphql.NewRequest(fetchQuery)
	userRequest.Var("email", body.Input.Email)
	userRequest.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	var SecondRespose model.LoginSchema
	err := initializers.CLIENT.Run(context.Background(), userRequest, &SecondRespose)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"token":   "",
			"id":      "",
		})
		return
	}

	if len(SecondRespose.Users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is not found in the system",
			"token":   "",
		})
		return

	}
	//hash the password
	err = bcrypt.CompareHashAndPassword([]byte(SecondRespose.Users[0].Password), []byte(body.Input.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong password",
			"token":   "",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": SecondRespose.Users[0].Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-default-role":  "user",
			"x-hasura-allowed-roles": []string{"user", "admin"},
			"x-hasura-user-id":       SecondRespose.Users[0].Id,
		},
	})

	tokenString, tokenError := token.SignedString([]byte(os.Getenv("SECRET")))

	if tokenError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tokenError.Error(),
			"token":   "",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"message": "successfully loged in to the server",
		"id":      SecondRespose.Users[0].Id,
	})
}
func Upload(c *gin.Context) {
	uniqueID := uuid.New().ID()
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "first " + err.Error(),
		})
		return
	}
	parts := strings.Split(file.Filename, string(rune('.')))
	fmt.Println(len(parts))
	filename := fmt.Sprint(uniqueID) + "." + parts[len(parts)-1]
	err = c.SaveUploadedFile(file, "assets/"+filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success" + "http://localhost:3000/assets/" + filename,
		"url":     "http://localhost:3000/assets/" + filename,
	})
}
