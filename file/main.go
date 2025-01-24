package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
    // create file

	file,err:=os.Create("file.txt")
    if err!=nil{
        fmt.Println("error occur during file creation ")
    }
    defer file.Close() // os tells us that,free my resource which required during the file ceartion
    fmt.Println("successfully created the file");

    
    // write something in file
    content :="hello i am shubh as web developer in bluestock"
    // _,error:=io.WriteString(file,content)   // return two thing (byte, err)
    _,error:=io.WriteString(file,content +"\n")   // return two thing (byte, err)
        // if i want to that cursor goes in next line so i simply add content+"\n" instead of content

    if error!=nil{
        fmt.Println("error while writing in file",error)
    }


    // read file 

    


    // fmt.Println(file);
}