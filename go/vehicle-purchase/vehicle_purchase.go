package purchase

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true
	case
		"bike", "stroller", "e-scooter":
		return false
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	sentence := " is clearly the better choice."
	if option1 < option2 {
		return option1 + sentence
	}
	return option2 + sentence
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return originalPrice * 0.8
	case age < 10 && age >= 3:
		return originalPrice * 0.7
	case age >= 10:
		return originalPrice * 0.5
	}
	return originalPrice
}
