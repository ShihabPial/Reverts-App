package mongoDB

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "revert_app/model"
    common "revert_app/common"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectToMongoDB(url string) *mongo.Client {
	// Set up a client options
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

// Example function to create a sample Person
func CreateSamplePerson() model.Person {
    return model.Person{
        FirstName:   "Fayek",
        LastName:    "Ahmed",
        Age:         20,
        Gender:      "Male",
        Location:    model.Location{City: "New York City", State: "New York"},
        Email:       "fayek.ahmed@example.com",
        PhoneNumber: "1234567890",
    }
}

// Handler for POST request to create a new user
func Insert(collection *mongo.Collection,w http.ResponseWriter, r *http.Request) {
    var user model.Person
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := collection.InsertOne(context.Background(), user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user.ID = result.InsertedID.(primitive.ObjectID)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}



// Upodate the sample document with additional fields
func Update(collection *mongo.Collection, sampleDocument bson.M) {
	// Define the update document with additional fields
	update := bson.M{
		"$set": bson.M{
			"Email":       "shihabpial1998@gmail.com",
			"PhoneNumber": "6314820631",
		},
	}

	// Update all documents to include the new fields
	result, err := collection.UpdateMany(context.TODO(), bson.M{}, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and modified %v documents.\n", result.MatchedCount, result.ModifiedCount)
}

func GetAllUsers(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) {
    var users []model.Person
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user model.Person
        cursor.Decode(&user)
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

//Update user
func UpdateUser(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, common.Missing_ID, http.StatusBadRequest)
        return
    }

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, common.Invalid_ID, http.StatusBadRequest)
        return
    }

    var user model.Person
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    filter := bson.M{"_id": objectId}
    update := bson.M{
        "$set": user,
    }

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if result.MatchedCount == 0 {
        http.Error(w, common.No_User_Found, http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

//Delete User
func DeleteUser(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, common.Missing_ID, http.StatusBadRequest)
        return
    }

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, common.Invalid_ID, http.StatusBadRequest)
        return
    }

    filter := bson.M{"_id": objectId}
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, common.No_User_Found, http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

//Get User by ID
func GetUserByID(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, common.Missing_ID, http.StatusBadRequest)
        return
    }

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, common.Invalid_ID, http.StatusBadRequest)
        return
    }

    var user model.Person
    filter := bson.M{"_id": objectId}
    err = collection.FindOne(context.Background(), filter).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, common.No_User_Found, http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}


