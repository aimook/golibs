package v

import (
	"encoding/json"
	"testing"
)

func TestSuccess(t *testing.T) {
	v := Success()
	s1, _ := json.Marshal(v)
	t.Logf("s1.v==>%p", v)

	v.Data(map[string]interface{}{
		"name": "tom",
		"age":  12,
	})
	s2, _ := json.Marshal(v)
	t.Logf("s2.v==>%p", v)

	v.Message("modify message")
	s3, _ := json.Marshal(v)

	t.Logf("s3.v==>%p", v)
	t.Logf("s1=>%s,s1.addr=>%p", s1, s1)
	t.Logf("s1=>%s,s2.addr=>%p", s2, s2)
	t.Logf("s1=>%s,s3.addr=>%p", s3, s3)
	//t.Logf("s2=>%s", s2)
	//t.Logf("s3=>%s", s3)

	v2 := Fail(500)

	vs2, _ := json.Marshal(v2)
	t.Logf("%s", vs2)

	v2.Message("message")
	vs3, _ := json.Marshal(v2)
	t.Logf("%s", vs3)

	v2.Data("data")
	vs4, _ := json.Marshal(v2)
	t.Logf("%s", vs4)

}
