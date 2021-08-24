#include <stdio.h>
#include <malloc.h>


typedef struct IPhoneServices{
    int (*getName)(int);
    int (*setName)(int);
} INewPhoneServices;

int getName(){
    printf("%s\n","getName");
    return 0;
}

int setName(){
    printf("%s\n","setName");
    return 1;
}


int main(){
    INewPhoneServices iNewPhoneServices = {getName,setName};
    iNewPhoneServices.getName = getName;
    iNewPhoneServices.setName = setName;
    iNewPhoneServices.getName(0);
    iNewPhoneServices.setName(1);
    return 0;
}
