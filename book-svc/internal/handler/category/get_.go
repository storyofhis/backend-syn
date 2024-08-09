package category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/storyofhis/book-svc/internal/repository/model/category"
)

func GetCategoryById(id int) (*category.Category, error) {
	url := fmt.Sprintf("http://localhost:8082/v1/categories/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[ERROR] : Failed to fetch category", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("[ERROR] : Non-OK HTTP status:", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch category, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] : Failed to read response body", err)
		return nil, err
	}

	var category category.Category
	if err := json.Unmarshal(body, &category); err != nil {
		log.Println("[ERROR] : Failed to unmarshal category", err)
		return nil, err
	}

	return &category, nil
}
