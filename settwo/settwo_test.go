package settwo

import (
	"bytes"
	"testing"
)

func Test_PKCSPadding(t *testing.T) {
	text_test := "YELLOW SUBMARINE"
	result, err := PKCSPadding(text_test, 20)
	if err != nil {
		t.Fatalf("unexpected error : %v", err)
	}
	expected := append([]byte(text_test), bytes.Repeat([]byte{4}, 4)...)

	if !bytes.Equal(result, expected) {
		t.Errorf("incorrect result:\nexpected: %q\ngot: %q", expected, result)
	}
}
