package main

import "fmt"
import "os"

func main(){
	defer fmt.Println("Hai")
	fmt.Println("Selamat Datang")
	os.Exit(1)
	fmt.Println("Selamat Datang")
	
}