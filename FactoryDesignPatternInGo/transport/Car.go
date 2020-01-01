package transport

// define a Car struct, the struct contain a string representing the type name
type Car struct {
	typeName string
}

//The Car struct implements the start() function
func (cr *Car) Start() {
	cr.typeName = " Car"
}

//The Car struct implements the GetMode() function
func (cr *Car) GetMode() string {
	return "I am a " + cr.typeName + " I travel on Road!!!"
}
