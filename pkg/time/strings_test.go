package t

import "testing"

func TestToString(t *testing.T) {

	var obj interface{}
	obj = 123445454
	t.Log(ToString(obj))

	var i64 int64
	i64 = 111
	t.Log(ToString(i64))

	var b1 bool
	b1 = true
	t.Log(ToString(b1))

}
