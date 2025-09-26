package main

import "fmt"

type Trash struct {
	name string
	material string
}

type TrashCan struct {
	color string
	material string
	content []Trash
}

func sort(trash Trash, green TrashCan, yellow TrashCan, brown TrashCan) (TrashCan, TrashCan, TrashCan) {
	if trash.material == "metal" {
		green.content = append(green.content, trash)
	} else if trash.material == "plastic" {
		yellow.content = append(yellow.content, trash)
	} else {
		brown.content = append(brown.content, trash)
	}

	return green, yellow, brown
}

func main() {
	var emptyContent []Trash

	green := TrashCan{"green", "metal", emptyContent}
	yellow := TrashCan{"yellow", "plastic", emptyContent}
	brown := TrashCan{"brown", "other", emptyContent}

	bottle := Trash{"bottle", "plastic"}

	green, yellow, brown = sort(bottle, green, yellow, brown)

	fmt.Println(green)
	fmt.Println(yellow)
	fmt.Println(brown)
}