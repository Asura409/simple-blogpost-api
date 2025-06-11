import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Posts struct {
	ID int
	author string
	title string
	content string
}

var posts [] Posts

func add_posts (w http.ResponseWriter, r *http.Request) {
	var new_post Posts
	err := json.NewDecoder(r.Body).Decode(&new_post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
		posts = append(posts, new_posts)