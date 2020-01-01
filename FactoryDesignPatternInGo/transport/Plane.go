package transport

// define a Car struct, the struct contain a string representing the type name
type Plane struct {
	typeName string
}

//The Car struct implements the start() function
func (pl *Plane) Start() {
	pl.typeName = " Plane"
}

//The Car struct implements the GetMode() function
func (pl *Plane) GetMode() string {
	return "I am a " + pl.typeName + " I travel in the air!!!"
}
