package easy

type ParkingSystem struct {
	slots [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slots: [3]int{big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType-1] <= 0 {
		return false
	}

	this.slots[carType-1]--
	return true
}
