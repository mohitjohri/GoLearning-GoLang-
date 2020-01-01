package main

import(
	"FactoryDesignPatternInGo/transport"
	"fmt"
)

func main() {

	myTransport, err:= transport.CreateTransport()

	//if no errors start the appliance then print it's purpose
	if err == nil {
		myTransport.Start()
		fmt.Println(myTransport.GetMode())
	} else {
		//if error encountered, print the error
		fmt.Println(err)
	}
}
