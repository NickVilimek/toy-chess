package users

import (
	"log"
	"net/http"
	"toy-chess/lib"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//func GetUsersRoute(c *gin.Context) {
//log.Println("[user-routes.GetUsersRoute] Calling Users Routes")

//email := c.Query("email")

//user, err := GetUsers(email)
//if err != nil {
//log.Printf("[user-routes.GetUsersRoute] Error getting users - %v\n", err)
//c.IndentedJSON(http.StatusBadRequest, nil)
//}

//log.Println("[user-routes.GetUsersRoute] Returning Successful get users query")

//c.IndentedJSON(http.StatusOK, user)
//}

/* Temp Mock Functions */

var user1Id, _ = primitive.ObjectIDFromHex("5f87924388202b76c32d8651")
var user2Id, _ = primitive.ObjectIDFromHex("5f87924388202b76c32d8652")
var user3Id, _ = primitive.ObjectIDFromHex("5f87924388202b76c32d8653")

var users = []User{
	{ID: user1Id, Email: "fake-email-1@mail.com", FirstName: "Fake", LastName: "One"},
	{ID: user2Id, Email: "fake-email-2@mail.com", FirstName: "Fake", LastName: "Two"},
	{ID: user3Id, Email: "fake-email-3@mail.com", FirstName: "Fake", LastName: "Three"},
}

func GetUsersRoute(c *gin.Context) {

	log.Println("[user-routes.GetUsersRoute] Calling Users Routes")

	email := c.Query("email")
	if email == "" {
		tempUsers := make([]User, len(users))
		copy(tempUsers, users)

		log.Println("[user-routes.GetUsersRoute] Returning successful user query with no filters")

		c.IndentedJSON(http.StatusOK, tempUsers)
		return
	}

	log.Printf("[user-routes.GetUsersRoute] email is %s\n", email)

	tempUsers := []User{}
	for _, value := range users {
		if value.Email == email {
			tempUsers = append(tempUsers, value)
		}
	}

	log.Println("[user-routes.GetUsersRoute] Returning Successful get users query")

	c.IndentedJSON(http.StatusOK, tempUsers)
	return

}

func GetUserByIdRoute(c *gin.Context) {

	log.Println("[user-routes.GetUserByIdRoute] Calling Users Routes")

	idParam := c.Param("id")
	if idParam == "" {
		message := "Please pass id in the route"
		log.Println("[user-routes.GetUserByIdRoute] Missing id param")
		c.IndentedJSON(http.StatusBadRequest, map[string]interface{}{"message": message})
		return
	}

	mongoId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		log.Println("[user-routes.GetUserByIdRoute] id param not a valid bson id")
		c.IndentedJSON(http.StatusBadRequest, map[string]interface{}{"message": "Id provided is not valid"})
		return
	}

	var user *User
	foundUser := false

	for _, value := range users {
		if value.ID == mongoId {
			foundUser = true
			user = &value
		}
	}

	if foundUser == false {
		log.Println("[user-routes.GetUserByIdRoute] No user found")
		c.IndentedJSON(http.StatusBadRequest, map[string]interface{}{"message": "No user found for the provided id"})
		return
	}

	c.IndentedJSON(http.StatusOK, &user)
}

func CreateUserRoute(c *gin.Context) {

	log.Println("[user-routes.CreateUserRoute] Calling Users Routes")

	var user User

	// Bind the JSON request body to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid JSON in request"})
		return
	}

	// Validate the User struct using the validator package
	if err := lib.GlobalValidatorInstance.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		return
	}

	user.ID = primitive.NewObjectID()

	users = append(users, user)

	c.JSON(http.StatusOK, user)

}
