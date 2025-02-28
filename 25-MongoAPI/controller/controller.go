package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abhijeetsharan/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const connectionString = "mongodb+srv://abhijeetsharan1510:abhi6093@cluster0.shyoc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Connect to MongoDB
func init() {
	clientOptions := options.Client().ApplyURI(connectionString) // ✅ Corrected
	client, err := mongo.Connect(context.TODO(), clientOptions) // ✅ Fixed argument type
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance created!")
}

// Insert one movie
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie:", inserted.InsertedID)
}

// Update one movie
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated one movie:", result.ModifiedCount)
}

// Delete one movie
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted one movie:", deleteResult.DeletedCount)
}

// Delete all movies
func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all movies:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all movies
func getAllMovies() []bson.M { // ✅ Changed primitive.M to bson.M
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M
	for cur.Next(context.Background()) {
		var movie bson.M
		if err := cur.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) // ✅ Now both are bson.M
	}
	defer cur.Close(context.Background())
	return movies
}

// API Handlers

// Get all movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// Create a new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// Mark movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie marked as watched", "id": params["id"]})
}

// Delete a movie
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted", "id": params["id"]})
}

// Delete all movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(map[string]int64{"deleted_count": count})
}
