package main

import "fmt"

func main(){
	str1 := []string{"a","b","c"}
	str2 := str1[1:]
	str2[1] = "New"
	fmt.Println(str1)

	str2 = append(str2,"x","y","z")
	fmt.Println(str1)
}