package differential

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/keelerm84/go-conduit/conduit"
)

type diffResult struct {
	Results map[string]Diff `json:"result"`
}

type Diff struct {
	DateCreated string `json:dateCreated`
}

type DiffQuery struct {
	Conduit conduit.Connection `json:"__conduit__"`
	Ids     []string           `json:"ids"`
}

func (q *DiffQuery) Search() map[string]Diff {
	searchParams, _ := json.Marshal(q)

	v := url.Values{}
	v.Set("params", string(searchParams))
	v.Set("output", "json")

	resp, _ := http.PostForm(q.Conduit.Host+"/api/differential.querydiffs", v)
	body, _ := ioutil.ReadAll(resp.Body)

	var result diffResult
	json.Unmarshal(body, &result)

	return result.Results
}
