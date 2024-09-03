package helper

import "fmt"


//Error
func ReturnError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
