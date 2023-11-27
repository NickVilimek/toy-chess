package users

import (
	"log"
)

func GetUsers(email string) (*[]User, error) {
	log.Println("[user-service.GetUsers] Getting users")

	filters := make(map[string]interface{})

	if email != "" {
		log.Printf("[user-service.GetUsers] Filtering by email: %s", email)
		filters["email"] = email
	}

	users, err := GetMongoUsers(&filters)

	if err != nil {
		log.Println("[user-service.GetUsers] Error from repo")
		return nil, err
	}

	log.Println("[user-service.GetUsers] Returning successfully")
	return users, nil

}
