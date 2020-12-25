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

int cant_mazo = 108;

/*
crear_carta 
recibe una dirección para crear una carta del UNO, toma este path y crea el archivo .txt
y crea el archivo correspondiente.
*/
void crear_carta(char* path)
{
	int fpa;
	fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
	close(fpa);
}

/*
eliminar_carta 
recibe una dirección para eliminar una carta del UNO
*/
void eliminar_carta(char* path)
{
	remove(path);
}

/*
mover_carta 
recibe una dirección de origen y de destino para luego eliminar una carta del origen
y crearla en el destino
*/
void mover_carta(char* origen, char* destino)
{
	crear_carta(destino);
	eliminar_carta(origen);
}

/*
cartamazo 
te retorna una carta al azar del mazo servira en el futuro para la implementación del juego
de todas formas fue utilizada para iniciar la partida
*/
char* cartaMazo()
{
	srand(time(NULL));
	int k,r;
	char* temp2;
	DIR*d;
	struct dirent *dir;


	temp2 = (char*)malloc(sizeof(char)*30);
	k = 0;
	r = rand()%cant_mazo;
	d = opendir("./mazo");


	if(d)
	{
		while((dir = readdir(d)) != NULL)
		{
			if(strcmp(".",dir->d_name) != 0 && strcmp("..",dir->d_name) != 0)
			{
				if(k == r)
				{
					strcpy(temp2,dir->d_name);		
				}
				k++;
			}
					
		}
		cant_mazo--;
	}
	return temp2;
}

/*
iniciar_partida 
reparte 7 cartas a cada jugador y pone una en el pozo
*/
void iniciar_partida()
{
	srand(time(NULL));
	int k,r;
	char temp[30],temp2[30];
	
	DIR*d;
	struct dirent *dir;

	mkdir("mano1", 0700);
	mkdir("mano2", 0700);
	mkdir("mano3", 0700);
	mkdir("mano4", 0700);

	robarXCartas(7,1);
	robarXCartas(7,2);
	robarXCartas(7,3);
	robarXCartas(7,4);
	
	mkdir("pozo", 0700);
	
	r = rand()%cant_mazo;
	d = opendir("./mazo");
	strcpy(temp, "pozo/");
	strcpy(temp2,"mazo/");
	k = 0;

	if(d)
	{
		while((dir = readdir(d)) != NULL)
		{
			if(strcmp(".",dir->d_name) != 0 && strcmp("..",dir->d_name) != 0)
			{
				if(k == r)
				{
					strcat(temp,dir->d_name);
					strcat(temp2,dir->d_name);
					mover_carta(temp2,temp);
				}
				k++;
			}
					
		}
		cant_mazo--;
	}
}

/*
robarXcartas 
mueve X cartas desde el mazo y hacia la mano de algun jugador
*/
void robarXCartas(int a, int mano)
{
	char* carta;
	char origen[30], destino[30], wow[30];

	sprintf(wow, "mano%d/",mano);
	for (int i = 0; i < a; ++i)
	{	
		strcpy(destino, wow);
		strcpy(origen, "mazo/");
		carta = cartaMazo();
		
		strcat(origen, carta);
		strcat(destino, carta);
		free(carta);
		mover_carta(origen, destino);
		
	}

}


/*
mazo
crea el estado inicial del mazo con las 108 cartas
*/
void mazo()
{
	int fpa;
	char num[12], path[1000], temp[12];
	char colo[2];
	char colores[9];

	mkdir("mazo", 0700);
	strcpy(colores, "RBYG");
	
	for (int z = 0; z < 2; ++z)
	{
		for (int i = 0; i < 4; ++i)
		{
			colo[0] = colores[i];
			colo[1] = '\0';
			for (int j = 0; j < 10; ++j)
			{
				if (z == 1)
				{
					if (j == 0)
					{
						j++;
					}
					
				}
				sprintf(num, "%d",j);
				sprintf(temp, "%d",z+1);
				strcat(num, colo);
				strcat(num, temp);
				strcat(num, ".txt");
				strcpy(path, "mazo/");
				strcat(path, num);
				fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
				close(fpa);
			}


		strcpy(path, "mazo/+2");
		strcat(path, colo);
		strcat(path, "1.txt");
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);
		strcpy(path, "mazo/+2");
		strcat(path, colo);
		strcat(path, "2.txt");
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);



		strcpy(path,"mazo/+4N");
		sprintf(temp, "%d", i+1);
		strcat(temp, ".txt");
		strcat(path, temp);
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);


		strcpy(path, "mazo/Rev");
		strcat(path, colo);
		strcat(path, "1.txt");
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);
		strcpy(path, "mazo/Rev");
		strcat(path, colo);
		strcat(path, "2.txt");
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);


		strcpy(path, "mazo/Skip");
		strcat(path, colo);
		strcat(path, "1.txt");
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);
		strcpy(path, "mazo/Skip");
		strcat(path, colo);
		strcat(path, "2.txt");
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);

		strcpy(path,"mazo/CCN");
		sprintf(temp, "%d", i+1);
		strcat(temp, ".txt");
		strcat(path, temp);
		fpa = open(path, O_RDWR | O_APPEND | O_CREAT);
		close(fpa);

		}
	}
}

/*
obtenercartas
esta funcion retorna las cartas presentas en cualquier lugar tanto mazo, pozo o alguna mano
*/
char** obtenercartas(char* carpeta)
{
	char** cartas = (char**)malloc(sizeof(char*)*109);
	DIR*d;
	struct dirent *dir;
	int i = 0;
	char path[12];
	strcpy(path,"./");
	strcat(path,carpeta);
	d = opendir(path);

	if(d)
	{
		while((dir = readdir(d)) != NULL)
		{
			if(strcmp(".",dir->d_name) != 0 && strcmp("..",dir->d_name) != 0)
			{
				cartas[i] = (char*)malloc(sizeof(char)*9);
				strcpy(cartas[i],dir->d_name);
				i++;
			}
		}
	}
	cartas[i] = (char*)malloc(sizeof(char)*9);
	strcpy(cartas[i],"STOP");
	return cartas;
}