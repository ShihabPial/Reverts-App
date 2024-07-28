package main

import (
	"log"
	"net/http"
	mongoDB "revert_app/MongoDB"
)

func main() {
    //Set the client
    client := mongoDB.ConnectToMongoDB("mongodb://localhost:27017/")
    
    //Get the collection
    collection := client.Database("revert_app").Collection("users")

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            mongoDB.Insert(collection, w, r)
        case http.MethodGet:
            if r.URL.Query().Get("id") != "" {
                mongoDB.GetUserByID(collection, w, r)
            } else {
                mongoDB.GetAllUsers(collection, w, r)
            }
        case http.MethodPut:
            mongoDB.UpdateUser(collection, w, r)
        case http.MethodDelete:
            mongoDB.DeleteUser(collection, w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server starting on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}
