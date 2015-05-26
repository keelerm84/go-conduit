package differential

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/keelerm84/go-conduit/conduit"
)

// Revision defines what a Phabricator revision looks like, including who
// reviewed it and a link to the revision.
type Revision struct {
	ID        string   `json:"id"`
	Phid      string   `json:"phid"`
	Title     string   `json:"title"`
	URI       string   `json:"uri"`
	Reviewers []string `json:"reviewers"`
	Diffs     []string `json:"diffs"`
}

// RevisionQuery contains various supported fields by which a user might want
// to search.
type RevisionQuery struct {
	Conduit   conduit.Connection `json:"__conduit__"`
	Status    string             `json:"status"`
	Reviewers []string           `json:"reviewers"`
}

// Search queries the API for revisions using the criteria provided in the
// RevisionQuery struct.
func (q *RevisionQuery) Search() []Revision {
	searchParams, _ := json.Marshal(q)

	v := url.Values{}
	v.Set("params", string(searchParams))
	v.Set("output", "json")

	resp, _ := http.PostForm(q.Conduit.Host+"/api/differential.query", v)

	result := struct {
		Results []Revision `json:"result"`
	}{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result.Results
}
