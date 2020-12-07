package main

import "fmt"

/*
Goal: train outer loops

Task: You have  several vehicle factories, one of them produces tesla cars,
If factory produces tesla (desiredNames list), need to check tesla vehicle:
- check engine is electric
- check autopilot is exist
- check 1-60 mph less than 5 seconds

Use:
- Outer loops

*/

const (
	gasolineEngine = "benzin"
	electricEngine = "electric"
	fastCarGoesIn  = 5.0
)

// Car is a typical car properties
type Car struct {
	Name         string
	Engine       string
	Autopilot    bool
	ToSixtyMiles float32
}

// Factory is a typical car factory
type Factory struct {
	Name string
	Cars []Car
}

func teslaChecker(vehicle *Car) error {
	customErr := func(whatsWrong string) error {
		return fmt.Errorf("Requirement %s does not meet expectations of the Tesla: i need electro + autopilot + really fast!", whatsWrong)
	}

	if vehicle.Engine != electricEngine {
		return customErr("Engine")
	}
	if !vehicle.Autopilot {
		return customErr("Autopilot")
	}

	if vehicle.ToSixtyMiles > fastCarGoesIn {
		return customErr("Speed")
	}

	return nil
}

func main() {
	allCars := []Car{
		Car{
			Name:         "gm-foo",
			Engine:       gasolineEngine,
			Autopilot:    false,
			ToSixtyMiles: 6.7,
		},
		Car{
			Name:         "gm-bar",
			Engine:       electricEngine,
			Autopilot:    false,
			ToSixtyMiles: 4.4,
		},
		Car{
			Name:         "gm-xpeng-teslahuesla",
			Engine:       electricEngine,
			Autopilot:    true,
			ToSixtyMiles: 8.5,
		},
		Car{
			Name:         "xpeng-teslafake",
			Engine:       gasolineEngine,
			Autopilot:    true,
			ToSixtyMiles: 8.5,
		},
		Car{
			Name:         "xpeng-tesla",
			Engine:       electricEngine,
			Autopilot:    true,
			ToSixtyMiles: 4.4,
		},
	}
	carsAmount := len(allCars)

	// getting factories
	generalMotors := Factory{
		Name: "General Motors",
		Cars: allCars[:carsAmount-2],
	}
	xPeng := Factory{
		Name: "XPeng Motors",
		Cars: allCars[2:],
	}
	allFactories := []Factory{generalMotors, xPeng}

	fmt.Printf("Here is GM cars %v\n", generalMotors)
	fmt.Printf("Here is Xpeng cars %v\n", xPeng)

	desiredCarsList := []string{"xpeng-tesla", "xpeng-teslafake", "gm-xpeng-teslahuesla"}

	for _, factory := range allFactories {
	outer: // here is label to return for factories cars looking
		for _, vehicle := range factory.Cars {
			for _, desiredCar := range desiredCarsList {
				fmt.Printf("Looking for %s, they offered me: %s \n", desiredCar, vehicle.Name)

				if desiredCar == vehicle.Name {
					fmt.Printf("Hurray, i found %s, need to check it ========>>>>>>>\n", desiredCar)
					err := teslaChecker(&vehicle)
					if err != nil {
						fmt.Printf("%s its fucking fake: %s....\n\n", desiredCar, err)
					} else {
						fmt.Printf("%s suites me!!! but need to look smth else =)\n\n", desiredCar)
						continue outer
					}
				}
			}

		}
	}

}
