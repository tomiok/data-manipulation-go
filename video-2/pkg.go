package main
/*
	Media => promedio

	Mediana => un arreglo ordenado -> buscamos el valor del medio
	[24, 27, 31]
	[24, 27, 31, 45]


	Moda => el valor que mas se repite
	[1, 1, 4, 5, 9, 9, 1]
	1
	1 IS IN THE MAP?
		if true => 1 appears one more time
	if false
		add 1 in to the map with value 1

map definition:
	1 => 3
	4 => 1
	5 => 1
	9 => 2

	Media ponderada => cada valor va a estar ponderado

							ponderacion			Estudiante X			Estudiante Y
	tarea						10				90							85
	examen final				30				85							96
	examenes orales				25				75							70
	trabajos practicos			20				96							100
	asistencia					15				100							99
 */