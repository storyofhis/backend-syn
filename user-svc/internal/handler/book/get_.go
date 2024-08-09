package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http" // Ensure time package is imported if needed

	"github.com/storyofhis/auth-svc/internal/repository/model/book"
)

func GetBookById(id int) (*book.Book, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/books/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var book book.Book
	if err := json.Unmarshal(body, &book); err != nil {
		return nil, err
	}

	return &book, nil
}
