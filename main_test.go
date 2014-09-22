package texco

import (
	"testing"
)

func TestAddConstant(t *testing.T) {
	AddConstant("TEST", "YOLO")
	AddConstant("NUM", 33)
	AddConstant("NONO", "NOTTEST")
	if len(consts) < 1 {
		t.Fatalf("Could not add test constant")
	}
}

func TestPassFile(t *testing.T) {
	expecting := "YOLO NOTTEST"
	received, _ := PassFile("sample.txt")
	if expecting != received {
		t.Fatalf("Did not receive expected, instead got", received)
	}
}

func TestPassString(t *testing.T) {
	expecting := "YOLO"
	source := "|TEST?"
	receieved := PassString(source)

	if expecting != receieved {
		t.Fatalf("Did not receive expected, instead got", receieved)
	}
}

func TestPassJavascript(t *testing.T) {
	expecting := "3"
	source := "|1 + 2?"
	receieved := PassString(source)

	if expecting != receieved {
		t.Fatalf("Did not receive expected, instead got", receieved)
	}
}

func TestPassJavascriptAndConstant(t *testing.T) {
	expecting := "36"
	source := "|3 + NUM?"
	receieved := PassString(source)

	if expecting != receieved {
		t.Fatalf("Did not receive expected, instead got", receieved)
	}
}

// func TestPullString(t *testing.T) {
// 	// source := "hi =onlythistextisvisible=i'm bob"
// 	// a := strings.Index(source, "=")
// 	// part1 := source[a+1:]
// 	// b := strings.Index(part1, "=")
// 	// part2 := part1[:b]
// 	// log.Println(part1)
// 	// log.Println(part2)

// 	// log.Println("FINAL = "])
// }
