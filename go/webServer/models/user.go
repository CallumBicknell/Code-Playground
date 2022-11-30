package models

import (
	"errors"
	"log"
	"time"

	"github.com/CallumBicknell/go-webServer/db"
	"github.com/CallumBicknell/go-webServer/forms"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `gorm:"type:uuid;primary_key;user_id;"`
	Username  string    `gorm:"username;"`
	Password  string    `gorm:"password;"`
	Pin       string    `gorm:"pin;"`
	Active    bool      `gorm:"default:false;active;"`
	UpdatedAt time.Time `gorm:"updated_at;"`
	CreatedAt time.Time `gorm:"updated_at;"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	db := db.GetDB()
	id := uuid.NewV4()
	user := User{
		ID:        id.String(),
		Username:  userPayload.Username,
		Password:  userPayload.Password, // TODO: Hash this before saving
		Pin:       userPayload.Pin,      // TODO: Hash this before saving
		Active:    true,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("TableUsers"),
	}
	if _, err := db.PutItem(params); err != nil { // TODO: Fix me
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	return &user, nil
}

// func (h User) GetByID(id string) (*User, error) {
// 	db := db.GetDB()
// 	params := &dynamodb.GetItemInput{
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"user_id": {
// 				S: aws.String(id),
// 			},
// 		},
// 		TableName:      aws.String("TableUsers"),
// 		ConsistentRead: aws.Bool(true),
// 	}
// 	resp, err := db.GetItem(params)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	var user *User
// 	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return user, nil
// }
