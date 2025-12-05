package main
import "fmt"

type RailroadWidthChecker interface {
	CheckRailsWidth() int
}

type Railroad struct {
	Width int
}

func (rail *Railroad) IsCorrectSizeTrain(train RailroadWidthChecker) bool {
	return train.CheckRailsWidth() == rail.Width
}

type Train struct {
	TrainWidth int
}

func (t *Train) CheckRailsWidth() int {
	return t.TrainWidth
}

func main() {
	railroad := Railroad{Width: 10}
	passengerTrain := &Train{TrainWidth: 10}
	cargoTrain := &Train{TrainWidth: 15}

	canPassengerTrainPass := railroad.IsCorrectSizeTrain(passengerTrain)
	canCargoTrainPass := railroad.IsCorrectSizeTrain(cargoTrain)

	fmt.Printf("Can passenger train pass? %t\n", canPassengerTrainPass) 
	fmt.Printf("Can cargo train pass? %t\n", canCargoTrainPass)
	
	fmt.Printf("Passenger: %v, Cargo: %v\n", canPassengerTrainPass, canCargoTrainPass)
}
