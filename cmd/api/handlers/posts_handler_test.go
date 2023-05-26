package handlers

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type EndToEndSuit struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {

	suite.Run(t, new(EndToEndSuit))
}

func (s *EndToEndSuit) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuit) TestPostNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/post/5334")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	
	s.JSONEq(`{"status":"ok", "data":[]}`, string(b))
}
