package utils

func CMToFeet(cm float64) (int, int) {
	// converter altura em centímetros para polegadas
	inches := int(cm / 2.54)

	// dividir as polegadas por 12 para obter o número de pés e o resto
	feet := inches / 12
	remainder := inches % 12

	return feet, remainder
}
