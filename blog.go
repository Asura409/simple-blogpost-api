package main
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Posts struct {
	ID int `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// Global variable to hold posts
// In a real application, you would use a database to store posts
var posts [] Posts

func add_posts(w http.ResponseWriter, r *http.Request) {
	var new_post Posts
	err := json.NewDecoder(r.Body).Decode(&new_post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
		posts = append(posts, new_post)
		fmt.Fprintf(w, "Post added successfully")
		json.NewEncoder(w).Encode(posts)
}

func edit_posts(w http.ResponseWriter, r *http.Request) {
	var updated_post Posts
	err := json.NewDecoder(r.Body).Decode(&updated_post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}	
	found := false
	for index, post := range posts {
		if post.ID == updated_post.ID {
			found = true
			posts[index] = updated_post
			fmt.Fprintf(w, "Post updated successfully")
			json.NewEncoder(w).Encode(posts)
		}
	}
	if !found {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
}

func delete_post(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)       // Extracts variables from the path
    idStr := vars["id"]       // Gets the value mapped to {id}
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

	found := false
	for index, post := range posts {
		if post.ID == id {
			found = true
			posts = append (posts[:index], posts[index+1:]...)
			fmt.Fprintf(w, "Post deleted successfully")
			json.NewEncoder(w).Encode(posts)	
		}
		if !found {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
	}
	}

func get_posts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if len (posts) == 0 {
		http.Error(w, "NO available posts", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(posts)
}

func get_post_by_id(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, post := range posts {
		if fmt.Sprintf("%d", post.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}
