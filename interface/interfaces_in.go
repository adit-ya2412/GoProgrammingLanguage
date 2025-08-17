package main

import "log"

type Vehcile interface {
	numberOfWheels() string
	fuelType() string
}

type Car struct {
	color     string
	typeOfCar string
}

type Bike struct {
	color       string
	enginePower string
}

func main() {

    mercedesSclass:=Car{
		color: "Black",
		typeOfCar: "Sedan",
	}

	yamhaR15:=Bike{
		color: "Blue",
		enginePower: "400cc",
	}

    printInfo(mercedesSclass)
	printInfo(yamhaR15)

}

func (c Car) numberOfWheels()string{
	return "4"
}
func (c Car) fuelType()string{
	return "Petrol"
}

func (b Bike) numberOfWheels()string{
	return "2"
}

func (b Bike) fuelType()string{
	return "Petrol"
}

func printInfo(vh Vehcile) {
	log.Println("Number of wheels in vehcile ",vh.numberOfWheels())
	log.Println("Fuel used in vehcile ",vh.fuelType())
}
