package differential

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/keelerm84/go-conduit/conduit"
)

// Diff defines what a Phabricator diff entry looks like.
type Diff struct {
	DateCreated string `json:dateCreated`
}

// DiffQuery contains various supported fields by which a user might want to
// search.
type DiffQuery struct {
	Conduit conduit.Connection `json:"__conduit__"`
	IDs     []string           `json:"ids"`
}

// Search queries the API for diffs using the criteria provided in the
// DiffQuery struct.
func (q *DiffQuery) Search() map[string]Diff {
	searchParams, _ := json.Marshal(q)

	v := url.Values{}
	v.Set("params", string(searchParams))
	v.Set("output", "json")

	resp, _ := http.PostForm(q.Conduit.Host+"/api/differential.querydiffs", v)

	result := struct {
		Results map[string]Diff `json:"result"`
	}{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result.Results
}
