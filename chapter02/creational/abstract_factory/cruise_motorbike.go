package abstractfactory

type CruiseMotorbike struct {}

func (c *CruiseMotorbike) NumWheels() int {
	return 2
}

func (c *CruiseMotorbike) NumSeats() int {
	return 2
}

func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}