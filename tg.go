package main;

import (
	"os"
        "bufio"
	"fmt"
	"flag"
	"log"
	"strings"
)

type Turtle struct {
	x, y int
	pendown bool
}

var size = 20
var filepath = flag.String("f", "commands.txt", "filepath to commands file")
var end = false

func NewTurtle() *Turtle {
	t := new(Turtle)
	t.x = size/2
	t.y = size/2
//	pendown := true
	
	return t
}

func(t *Turtle) moveTurtle(dir string) {
	switch d := strings.TrimRight(dir, "\n"); d {
	case "up":
		t.y += 1
	case "down":
		t.y -= 1
	case "right":
		t.x += 1
	case "left":
		t.x -= 1
	case "end":
		end = true
	default:
		fmt.Println("No command")
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*filepath)
	if err != nil {
 		log.Fatal(err)
 	}

 	defer file.Close()

	scanner := bufio.NewScanner(file)
	t := NewTurtle()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		t.moveTurtle(scanner.Text())
		fmt.Printf("%d, %d\n",t.x,t.y)
		if end == true {
			break
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


