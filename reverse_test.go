/*
 * Test of the Reverse function
 * more about testing on :
 * https://golang.org/pkg/testing/
 */
package stringutil

// Package to import to make the unit test
import "testing"

//naming : Testxxx where xxx name of the function to test
func TestReverse(t *testing.T) {
	// built an array of structure for the data set
	// (in value ,want or expected value)
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
