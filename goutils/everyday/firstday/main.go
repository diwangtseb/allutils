package main

import "fmt"

type IPersonService struct{
}

type PersonInterface interface{
    GetName()
    SetName()
}

func (personService *IPersonService) GetName(){
    fmt.Printf("shixianleGetName%v\n",personService)
}

func (personService *IPersonService) SetName(){
    fmt.Printf("shixianleSetName%v\n",personService)
}


func main(){
    iPersonService := IPersonService{}
    go iPersonService.GetName()
    go iPersonService.SetName()
}
