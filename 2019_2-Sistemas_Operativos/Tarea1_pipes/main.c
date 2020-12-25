#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <time.h>
#include <sys/errno.h>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <dirent.h>

#include "archivos.h"
#include "juego.h"

/*
un main de prueba para realizar testear las funciones realizadas para este entregable
*/

int main()
{
	char** cartas;
	int j;

	mazo();
	iniciar_partida();
	cartas = obtenercartas("mano1");

	printf("Cartas en la mano del jugador 1:\n");
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		printf("%s\n",cartas[j]);
		j++;
	}
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		free(cartas[j]);
		j++;
	}
	free(cartas);

	cartas = obtenercartas("mano2");
	printf("Cartas en la mano del jugador 2:\n");
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		printf("%s\n",cartas[j]);
		j++;
	}
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		free(cartas[j]);
		j++;
	}
	free(cartas);
	cartas = obtenercartas("mano3");
	printf("Cartas en la mano del jugador 3:\n");
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		printf("%s\n",cartas[j]);
		j++;
	}
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		free(cartas[j]);
		j++;
	}
	free(cartas);
	cartas = obtenercartas("mano4");
	printf("Cartas en la mano del jugador 4:\n");
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		printf("%s\n",cartas[j]);
		j++;
	}
	j = 0;
	while(strcmp(cartas[j],"STOP") != 0){
		free(cartas[j]);
		j++;
	}
	free(cartas);

	return 0;
}