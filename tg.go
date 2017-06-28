package main;

import (
        "bufio"
	"fmt"
	"flag"
	"math/rand"
	"log"
)

type Turtle struct {
	x, y int
	pendown bool
}

var size = 20
var filepath = flag.String("f", commands.txt, "filepath to commands file")
 

func NewTurtle() *Turtle {
	t := new(Turtle)
	t.x = size/2
	t.y = size/2
	pendown = true
}

func moveTurtle(string dir) *Turtle {
	switch d = strings.TrimRight(dir, "\n"); d {
	case "up":
		t.y += 1
	case "down":
		t.y -= 1
	case "right":
		t.x += 1
	case "left"
		t.x -= 1
	
}

func main() {
	flag.Parse()
	file, err := os.Open(filepath)
	if err != nil {
 		log.Fatal(err)
 	}

 	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


