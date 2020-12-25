package main

import (
	"fmt"
	"os"
	"errors"
	"ncode/ncode"
)


func main(){

	userNationalID, error := getArgument()
	if (error != nil){
		os.Exit(1)
	}

	nID := ncode.NewNationalID(string(userNationalID))
	if nID.Validity == true{
		fmt.Println("This NationaID is Valid")
	} else{
		fmt.Println("This NationaID is Invalid")
	}

}


func getArgument() (string,error){
	input := os.Args[1:]
	if len(input) != 0 {
		return input[0],nil
	}else{
		fmt.Println("You must enter nationalID like this: \"go run main.go NATIONALID\"")
		return "",errors.New("invalid")
	}
}