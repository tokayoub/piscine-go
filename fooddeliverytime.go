package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}
	if "burger" == order {
		return burger.preptime
	} else if "chips" == order {
		return chips.preptime
	} else if "nuggets" == order {
		return nuggets.preptime
	}
	return 404
}
