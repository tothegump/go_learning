package animals

import "testing"

func TestCat(t *testing.T) {
	s := CatCall(2)
	if s != "miao.. miao.. " {
		t.Error("cat_call failed")
	}
}
