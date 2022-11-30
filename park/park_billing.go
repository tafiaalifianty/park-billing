package park

type VehicleType int

const (
	Car VehicleType = iota
	Motorcycle
	Car_First        = 7000
	Car_After        = 5000
	Car_Extra        = 50000
	Motorcycle_First = 3000
	Motorcycle_After = 2000
	Motorcycle_Extra = 20000
)

func CalculateParkingBill(vehicleType VehicleType, duration int) float32 {
	var price int

	if vehicleType == 0 {
		if duration >= 24 {
			price = ((duration - 1) * Car_After) + Car_Extra + Car_First
		} else if duration < 24 && duration > 1 {
			price = ((duration - 1) * Car_After) + Car_First
		} else if duration > 0 && duration <= 1 {
			price = Car_First
		} else {
			price = 0
		}
	} else if vehicleType == 1 {
		if duration >= 24 {
			price = ((duration - 1) * Motorcycle_After) + Motorcycle_Extra + Motorcycle_First
		} else if duration < 24 && duration > 1 {
			price = ((duration - 1) * Motorcycle_After) + Motorcycle_First
		} else if duration > 0 && duration <= 1 {
			price = 3000
		} else {
			price = 0
		}
	}
	return float32(price)
}
