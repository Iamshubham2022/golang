// What This Program Covers:
// Create a File: Creates file.txt.
// Write to File: Adds content to the file.
// Append to File: Adds extra content without overwriting.
// Read File in Chunks: Reads file in 1024-byte chunks.
// Read Entire File: Reads the whole file in one go.
// Read Specific Lines: Reads file line by line with bufio.Scanner.
// Get Metadata: Displays file name, size, modification time, and permissions.
// Rename File: Renames file.txt to newfile.txt.
// Check File Existence: Verifies if a file exists.
// Copy File: Copies contents of newfile.txt to copy.txt.
// Delete File: Deletes newfile.txt.












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

    // file open for reading 
    f,er :=os.Open("file.txt")
    if(er!=nil){
        fmt.Println("error while opeing file",er)
    }
    defer f.Close()
    fmt.Println("file successfully open")

    // creating buffer storage for temp use;
    buffer :=make([]byte,1024)

    for{
        n,errors:=f.Read(buffer)
        if errors==io.EOF{
            break;
        }
        if errors!=nil{
            fmt.Println("error while storing in buffer")
        }
        // process the buffer 

        fmt.Println(string(buffer[:n]))
    }


    // read entire file but see if u don't have enough memory then it will through error ,
    // but in the above method that is read a chunk of data and then print it ,again and again don't use 
    // this for the long file oky

    conten ,eror := os.ReadFile("file.txt")
    if eror!=nil{
        fmt.Println("erorr when entire file reading ")
        return
    }
    fmt.Print(string(conten))

    


   
        // // 1. Create a File
        // file, err := os.Create("file.txt")
        // if err != nil {
        //     fmt.Println("Error creating file:", err)
        //     return
        // }
        // defer file.Close()
        // fmt.Println("File created successfully")
    
        // // 2. Write to the File
        // content := "Hello, I am learning Golang file handling.\n"
        // _, err = io.WriteString(file, content)
        // if err != nil {
        //     fmt.Println("Error writing to file:", err)
        //     return
        // }
        // fmt.Println("Content written successfully")
    
        // // 3. Append to the File
        // file, err = os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
        // if err != nil {
        //     fmt.Println("Error opening file for appending:", err)
        //     return
        // }
        // defer file.Close()
    
        // appendContent := "This line is appended to the file.\n"
        // _, err = file.WriteString(appendContent)
        // if err != nil {
        //     fmt.Println("Error appending to file:", err)
        //     return
        // }
        // fmt.Println("Content appended successfully")
    
        // // 4. Read the File (Chunk by Chunk)
        // file, err = os.Open("file.txt")
        // if err != nil {
        //     fmt.Println("Error opening file:", err)
        //     return
        // }
        // defer file.Close()
    
        // buffer := make([]byte, 1024)
        // fmt.Println("\nReading file chunk by chunk:")
        // for {
        //     n, err := file.Read(buffer)
        //     if err == io.EOF {
        //         break
        //     }
        //     if err != nil {
        //         fmt.Println("Error reading file:", err)
        //         return
        //     }
        //     fmt.Print(string(buffer[:n]))
        // }
    
        // // 5. Read the Entire File
        // fmt.Println("\n\nReading entire file:")
        // data, err := os.ReadFile("file.txt")
        // if err != nil {
        //     fmt.Println("Error reading file:", err)
        //     return
        // }
        // fmt.Println(string(data))
    
        // // 6. Read Specific Lines
        // fmt.Println("\nReading specific lines:")
        // file, err = os.Open("file.txt")
        // if err != nil {
        //     fmt.Println("Error opening file for line reading:", err)
        //     return
        // }
        // defer file.Close()
    
        // scanner := bufio.NewScanner(file)
        // lineNumber := 1
        // for scanner.Scan() {
        //     fmt.Printf("Line %d: %s\n", lineNumber, scanner.Text())
        //     lineNumber++
        // }
        // if err := scanner.Err(); err != nil {
        //     fmt.Println("Error reading lines:", err)
        // }
    
        // // 7. Get File Metadata
        // info, err := os.Stat("file.txt")
        // if err != nil {
        //     fmt.Println("Error getting file metadata:", err)
        //     return
        // }
        // fmt.Println("\nFile Metadata:")
        // fmt.Println("Name:", info.Name())
        // fmt.Println("Size:", info.Size(), "bytes")
        // fmt.Println("Last Modified:", info.ModTime())
        // fmt.Println("Permissions:", info.Mode())
    
        // // 8. Rename the File
        // err = os.Rename("file.txt", "newfile.txt")
        // if err != nil {
        //     fmt.Println("Error renaming file:", err)
        //     return
        // }
        // fmt.Println("\nFile renamed to 'newfile.txt' successfully")
    
        // // 9. Check if File Exists
        // if _, err := os.Stat("newfile.txt"); os.IsNotExist(err) {
        //     fmt.Println("File does not exist")
        // } else {
        //     fmt.Println("File exists")
        // }
    
        // // 10. Copy File Contents
        // source, err := os.Open("newfile.txt")
        // if err != nil {
        //     fmt.Println("Error opening source file:", err)
        //     return
        // }
        // defer source.Close()
    
        // dest, err := os.Create("copy.txt")
        // if err != nil {
        //     fmt.Println("Error creating destination file:", err)
        //     return
        // }
        // defer dest.Close()
    
        // _, err = io.Copy(dest, source)
        // if err != nil {
        //     fmt.Println("Error copying file:", err)
        //     return
        // }
        // fmt.Println("\nFile copied successfully to 'copy.txt'")
    
        // // 11. Delete a File
        // err = os.Remove("newfile.txt")
        // if err != nil {
        //     fmt.Println("Error deleting file:", err)
        //     return
        // }
        // fmt.Println("\nFile 'newfile.txt' deleted successfully")
    
    
}