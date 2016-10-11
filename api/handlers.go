package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/polypmer/mappyr/database"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	//err = json.NewEncoder(w).Encode
	fmt.Fprintln(w, "Index")
}

/* DB Read */

func ShowComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	c, err := database.ReadComment(db, id)
	if err != nil {
		w.Header().Set("Content-Type",
			"application/json;charset=UTF-8")
		w.WriteHeader(http.StatusNotFound) // Doesn't exist
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		w.Header().Set("Content-Type",
			"application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(c)
		if err != nil {
			fmt.Fprintf(w, "Error JSON encoding %s", err)
		}
	}
}

func ShowComments(w http.ResponseWriter, r *http.Request) {
	comments, err := database.ReadComments(db)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		fmt.Fprintf(w, "Error JSON encoding %s", err)
	}
}

/* DB Write */
// NewComment Writes a new comment to db
// If no user is specified, use 0 for id
func NewComment(w http.ResponseWriter, r *http.Request) {
	// This is taking a POST method
	var comment database.Comment
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}
	err = r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal, stick into my struct
	err = json.Unmarshal(body, &comment)
	// Unable to marshal?:
	if err != nil {
		// if shit don't work out
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) //422
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	}
	// Now I have a comment json object
	comment.Date = time.Now()
	// Check for Authentication, and in that case,
	// Add the user to the comment

	id, err := database.WriteComment(db, comment)
	if err != nil {
		fmt.Println(err)
	}
	comment.Id = id
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// Feed it back into the response Writer
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		fmt.Println(err)
	}
}
