package differential

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/keelerm84/go-conduit/conduit"
)

type commentsResult struct {
	Results map[string][]Comment `json:"result"`
}

type CommentsQuery struct {
	Conduit conduit.Connection `json:"__conduit__"`
	Ids     []string           `json:"ids"`
}

type Comment struct {
	Id          string `json:"id"`
	DateCreated string `json:"dateCreated"`
	Action      string `json:"action"`
	Content     string `json:"content"`
}

func (q *CommentsQuery) Search() map[string][]Comment {
	searchParams, _ := json.Marshal(q)

	v := url.Values{}
	v.Set("params", string(searchParams))
	v.Set("output", "json")

	resp, _ := http.PostForm(q.Conduit.Host+"/api/differential.getrevisioncomments", v)
	body, _ := ioutil.ReadAll(resp.Body)

	var result commentsResult
	json.Unmarshal(body, &result)

	return result.Results
}
