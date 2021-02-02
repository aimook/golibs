package api

import (
	"encoding/json"
	"testing"
)

func Test_Create(t *testing.T) {
	r := ResultCreate()
	s1, _ := json.Marshal(r)
	t.Logf("%s", s1)

	r.WithFail()
	s2, _ := json.Marshal(r)
	t.Logf("%s", s2)
}
