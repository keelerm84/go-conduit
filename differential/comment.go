package differential

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/keelerm84/go-conduit/conduit"
)

// CommentsQuery contains various supported fields by which a user might want
// to search.
type CommentsQuery struct {
	Conduit conduit.Connection `json:"__conduit__"`
	IDs     []string           `json:"ids"`
}

// Comment defines what a Phabricator comment looks like, including when it was
// created and what action the owner took to generate it.
type Comment struct {
	ID          string `json:"id"`
	DateCreated string `json:"dateCreated"`
	Action      string `json:"action"`
	Content     string `json:"content"`
}

// Search queries the API for comments using the criteria provided in the
// CommentsQuery struct.
func (q *CommentsQuery) Search() map[string][]Comment {
	searchParams, _ := json.Marshal(q)

	v := url.Values{}
	v.Set("params", string(searchParams))
	v.Set("output", "json")

	resp, _ := http.PostForm(q.Conduit.Host+"/api/differential.getrevisioncomments", v)

	result := struct {
		Results map[string][]Comment `json:"result"`
	}{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result.Results
}
