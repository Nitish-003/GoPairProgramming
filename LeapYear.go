package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main(){
	var a int64;
	scanner:= bufio.NewScanner(os.Stdin);
	fmt.Println("Enter the Year");
	scanner.Scan()
	a,_ = strconv.ParseInt(scanner.Text(),10,64);
	if(a % 4 == 0){
		if(a % 100 == 0){
			if( a % 400 == 0){
				fmt.Println("It is a leap Year")
				return;
			}
		}
		fmt.Println("It is a leap Year")
		return;
	}
	fmt.Println("It is not a leap year");
}