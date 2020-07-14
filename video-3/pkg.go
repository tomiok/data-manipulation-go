package main
/*
	Validacion de un modelo.

	Variables continuas (temperatura, porcentajes, cantidad de personas, desviacion, ganancias)
	Varibales discretas (estados: activo, inacativo, interrumpido, etc)

	Valores observados			valores pronosticados			margen de error
		10.5							8							3
		12.66							7.88						2.33
		09.2							6							1.76
		11.3							9.03						2


	MAE (Mean absolut error) error absoluto medio

	Podemos comparar el resultado con la media de los valores observados.

	formula:
		MAE = sum(abs(Vo - Vp)) / N

	N es el total de valores que tenemos.
 */
