package main

import "fmt"

type floder struct {
    components []component
    name string
}

func (f *floder) search(keywrods string){
    fmt.Printf("search %s's %s\n",f.name,keywrods)
    for _,ele := range(f.components){
        ele.search(keywrods)
    }
}

func (f *floder) add(c component){
    f.components = append(f.components, c)
}

type component interface{
    search(string)
}
