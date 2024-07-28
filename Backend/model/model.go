package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
    City  string `bson:"City"`
    State string `bson:"State"`
}

type Person struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    FirstName   string             `bson:"FirstName"`
    LastName    string             `bson:"LastName"`
    Age         int                `bson:"Age"`
    Gender      string             `bson:"Gender"`
    Location    Location           `bson:"Location"`
    Email       string             `bson:"Email"`
    PhoneNumber string             `bson:"PhoneNumber"`
}
