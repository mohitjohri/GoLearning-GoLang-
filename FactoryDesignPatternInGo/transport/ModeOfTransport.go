package transport

import (
	"errors"
	"fmt"
)

type Transport interface {
	Start()
	GetMode() string
}

const (
	CAR = iota
	PLANE
	TRAIN
)

func CreateTransport() (Transport, error) {
	//Request the user to enter the transport type
	fmt.Println("Enter preferred transport type")
	fmt.Println("0: Car ")
	fmt.Println("1: Plane")
	fmt.Println("2: Train")

	var myType int
	fmt.Scan(&myType)

	switch myType {
	case CAR:
		return new(Car), nil
	case PLANE:
		return new(Plane), nil
	case TRAIN:
		return new(Train), nil
	default:
		return nil, errors.New("Invalid transport Type")
	}
}
