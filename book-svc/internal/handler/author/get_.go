package author

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/storyofhis/book-svc/internal/repository/model/author"
)

func GetAuthorById(id int) (*author.Author, error) {
	url := fmt.Sprintf("http://localhost:8083/v1/authors/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[ERROR] : Failed to fetch author", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("[ERROR] : Non-OK HTTP status:", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch author, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] : Failed to read response body", err)
		return nil, err
	}

	var author author.Author
	if err := json.Unmarshal(body, &author); err != nil {
		log.Println("[ERROR] : Failed to unmarshal author", err)
		return nil, err
	}
	return &author, nil
}
