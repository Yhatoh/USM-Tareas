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
puedojugarla
recibe una carta y determina si es que la carta se puede jugar o no.

pd: por ahora retorna -1 si es que la carta es negra debido a que aun no creamos la funcionalidad con 
las cartas negras.
*/
int puedojugarla(char* carta)
{
	char temp[30];
	DIR*d;

	d = opendir("./pozo");
	struct dirent *dir;
	if(d)
	{
		while((dir = readdir(d)) != NULL)
		{
			if(strcmp(".",dir->d_name) != 0 && strcmp("..",dir->d_name) != 0)
			{
				strcpy(temp,dir->d_name);
			}
		}
	}

	if(temp[0] == 'S')
	{
		if(carta[0] == 'S') return 1;

		else
		{
			if(carta[0] == 'R') return carta[3] == temp[4];

			else if(carta[2] != 'N') return carta[2] == temp[4];	

			else{
				printf("Aca va la variable que tiene el color actual,\n si es que alguien lo cambio con una negra\n" ); 
				return -1;		
			}
		}
	}

	else if(temp[0] == 'R')
	{
		if(carta[0] == 'R') return 1;

		else
		{
			if(carta[0] == 'S') return carta[4] == temp[3];
				

			else if(carta[0] == '+' && carta[1] == '2') return carta[2] == temp[3];	

			else if(carta[2] != 'N') return carta[2] == temp[3];

			else{
				printf("Aca va la variable que tiene el color actual,\n si es que alguien lo cambio con una negra\n" ); 
				return -1;
			}
		}
	}


	else if(temp[2] != 'N')
	{
		if(carta[0] == temp[0]) return 1;

		else if (carta[1] == temp[1] && carta[0] != '+') return 1;

		else
		{
			if(carta[0] == 'S') return carta[4] == temp[1];

			else if(carta[0] == 'R') return carta[3] == temp[1];
				
			else if(carta[0] == '+' && carta[1] == '2') return carta[2] == temp[1];

			else{
				printf("Aca va la variable que tiene el color actual,\n si es que alguien lo cambio con una negra\n" );
				return -1;
			}
		}
	}
	else return 1;
}