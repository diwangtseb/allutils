package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(keywords string) {
	fmt.Printf("search %s's %s\n",f.name,keywords)
}

