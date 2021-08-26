package main

import "fmt"

// CoffeMachine represents an API to a hypotetical coffe maker
type CoffeeMachine struct {
	beanAmount   float64 // amount in ounces of beans to use
	grinderLevel int     // the granularity of the beans grinder
	waterTemp    int     // temperature of the water to use
	waterAmt     float64 // amount of water to use
	milkAmount   float64 // amount of milk to use
	addFoam      bool    // whether to add foam or not
}

func (c *CoffeeMachine) startCoffee(beanAmount float64, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans:", beanAmount, "and grind level:", grind)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans:", c.beanAmount, "beans at", c.grinderLevel)
	return true
}

func (c *CoffeeMachine) useMilk(amount float64) float64 {
	fmt.Println("Adding milk:", amount, "oz")
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp int) {
	fmt.Println("Setting water temp:", temp)
	c.waterTemp = temp
}

func (c *CoffeeMachine) useHotWater(amount float64) float64 {
	fmt.Println("Adding hot water:", amount, "oz")
	c.waterAmt = amount
	return amount
}

func (c *CoffeeMachine) doFoam(useFoam bool) {
	fmt.Println("Foam setting:", useFoam)
	c.addFoam = useFoam
}
