package main 



func main(){
    file1 := &file{
        name:"file1",
    }
    file2 := &file{
        name:"file2",
    }
    file3 := &file{
        name:"file3",
    }
    floder1 := &floder{
        name:"floder1",
    }
    floder1.add(file1)
    floder1.add(file2)
    floder1.add(file3)
    floder2 := &floder{
        name:"floder2",
    }
    floder2.add(floder1)
    floder2.search("rose")
}
