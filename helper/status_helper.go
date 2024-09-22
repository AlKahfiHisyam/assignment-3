package helper

func WaterStatus(water int) string {
	if water <= 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	}
	return "Bahaya"
}

func WindStatus(wind int) string {
	if wind <= 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	}
	return "Bahaya"
}
