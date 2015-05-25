package differential

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/keelerm84/go-conduit/conduit"
)

type Revision struct {
	Id        string   `json:"id"`
	Phid      string   `json:"phid"`
	Title     string   `json:"title"`
	Uri       string   `json:"uri"`
	Reviewers []string `json:"reviewers"`
	Diffs     []string `json:"diffs"`
}

type RevisionQuery struct {
	Conduit   conduit.Connection `json:"__conduit__"`
	Status    string             `json:"status"`
	Reviewers []string           `json:"reviewers"`
}

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
