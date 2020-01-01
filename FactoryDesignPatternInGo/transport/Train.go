package transport

// define a Car struct, the struct contain a string representing the type name
type Train struct {
	typeName string
}
//The Car struct implements the start() function
func (tr *Train) Start() {
	tr.typeName = " Train"
}
//The Car struct implements the GetMode() function
func (tr *Train) GetMode() string {
	return "I am a " + tr.typeName + " I travel on tracks!!!"
}
