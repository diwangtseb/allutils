#include<stdio.h>
#include<stdlib.h>

int main(int argc,char **argv){
    if (argc<3)
        return 1;
    char* file1 = argv[1];
    char* file2 = argv[2];
    FILE *filer,*filew;
    char ch;
    //读取源文件
    if((filer=fopen(file1,"r"))==NULL)
        printf("file isn't exist!\n");
    if((filew=fopen(file2,"w"))==NULL)
        printf("file has been created\n");
    while(((ch=fgetc(filer))!=EOF))
        fputc(ch,filew);
    fclose(filer);
    fclose(filew);
    return 0;

}
