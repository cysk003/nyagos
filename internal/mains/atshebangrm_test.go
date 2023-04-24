package mains

import (
	"testing"

	"golang.org/x/text/transform"
)

func TestAtShebangRemover(t *testing.T) {
	source := `@hogehoge
@ahahaha
@ihihihi
0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
2345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
3456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012
4567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123
5678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234
6789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345
7890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456
8901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567
9012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
2345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
3456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012
4567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123
5678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234
6789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345
7890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456
8901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567
9012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
@ihihihi
4567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123`
	expect := `


0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
2345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
3456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012
4567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123
5678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234
6789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345
7890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456
8901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567
9012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
2345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901
3456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012
4567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123
5678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234
6789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345
7890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456
8901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567
9012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
@ihihihi
4567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123`

	result, _, err := transform.String(&_AtShebangFilter{}, source)
	if err != nil {
		t.Fatal(err.Error())
	}
	if result != expect {
		t.Fatalf("expect `%s` but `%s`", expect, result)
	}
}