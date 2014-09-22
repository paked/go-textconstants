package texco

import "testing"

func TestAddConstant(t *testing.T) {
	AddConstant("TEST", "YOLO")
	if len(consts) < 1 {
		t.Fatalf("Could not add test constant")
	}
}

func TestPassFile(t *testing.T) {
	expecting := "YOLO"
	received, _ := PassFile("sample.txt")
	if expecting != received {
		t.Fatalf("Did not receive expected, instead got", received)
	}
}

func TestPassString(t *testing.T) {
	expecting := "YOLO"
	source := "=TEST="
	receieved := PassString(source)

	if expecting != receieved {
		t.Fatalf("Did not receive expected, instead got", receieved)
	}
}
