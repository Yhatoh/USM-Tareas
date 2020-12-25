#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>
#include "fun.h"
#include "TDALista.h"

int flag = 0;

/*
Funcion que genera una lista con datos al azar.
Recibe un int que representa la cantidad de datos que se desea.
Retorna la lista creada.
*/
void* generarSolucion(int n){
	if (flag == 0){
		srand(time(NULL));
		flag = 1;
	}
	tlista* lista = inicializar();
	int j,k;
	for (j = 0; j < n; j++){
		k = rand()%3;
		if (k == 0){
			int* numero = (int*)malloc(sizeof(int));
			*numero = rand()%10;
			lista = insertar_pos_actual(lista,numero, 'i');
		}
		else if (k == 1){
			int* numero = (int*)malloc(sizeof(int));
			*numero = rand()%2;
			lista = insertar_pos_actual(lista,numero, 'b');
		}
		else if (k == 2){
			int* numero = (int*)malloc(sizeof(int));
			*numero = 65+rand()%6;
			insertar_pos_actual(lista,numero, 'c');
		}
	}
	return lista;
}

/*
Funcion que copia una lista, creando una nueva, identica pero independiente.
Recibe la lista que se desea copiar.
Retorna la copia de la lista.
*/
void* copiar(void* Lista){
	tlista* ayuda = (tlista*) Lista;
	tlista* copia = inicializar();
	int i;
	moveTostart(ayuda);
	for (i = 0; i < length(ayuda); i++){
		int* numero = (int*)malloc(sizeof(int));
		*numero = *((int*)ayuda->curr->data);
		copia = insertar_pos_actual(copia,numero,ayuda->curr->type);
		next(ayuda);
		next(copia);
	}
	return copia;
}

/*
Funcion que libera la memoria de una lista.
Recibe la lista que desea ser liberada
*/
void borrar(void* lista){
	int i,size;
	tlista* lista1 = (tlista*)lista;
	moveTostart(lista1);
	tnodo* nodo1;
	size = lista1->listsize;
	for (i = 0; i < size-1; i++){
		nodo1 = lista1->curr;
		next(lista1);
		free(nodo1->data);
		free(nodo1);
		lista1->listsize -= 1;
	}
	free(lista1->curr->data);
	free(lista1->curr);
	free(lista1);
}

/*
Funcion que imprime en pantalla todos los datos de una lista.
Recibe la lista que se desea imprimir.
*/
void imprimirSolucion(void* Lista){
	int i;
	tlista* aux = (tlista*)Lista;
	moveTostart(aux);
	for (i = 0; i < length(aux); i++){
		if (aux->curr->type == 'i') printf("(%d,%c)",*((int*)aux->curr->data),aux->curr->type);
		else if (aux->curr->type == 'b') printf("(%d,%c)",*((int*)aux->curr->data),aux->curr->type);
		else if(aux->curr->type == 'c') printf("(%c,%c)",*((char*)aux->curr->data),aux->curr->type);
		next(aux);
	}
	printf("\n");
}

/*
Funcion que intercambia la primera mitad entre dos listas.
Recibe las listas que se desean intercambiar
*/
void cruceMedio(void* Lista1, void* Lista2)
{
	tnodo* temp;
	int n;
	tlista* l1 = (tlista*)Lista1;
	tlista* l2 = (tlista*)Lista2;
	n = length(l1)/2 - 1;
	moveTopos(l1, n);
	moveTopos(l2, n);
	temp = l1->curr->next;
	l1->curr->next = l2->curr->next;
	l2->curr->next = temp;
	moveTostart(l1);
	moveTostart(l2);
	temp = l1->head;
	l1->head = l2->head;
	l2->head = temp;
	moveTostart(l1);
	moveTostart(l2);
}

/*
Funcion que intercambia los nodos en posicion par de ambas listas.
Recibe las listas a las cuales se lese desea intercambiar los nodos.
*/
void cruceIntercalado(void* Lista1, void* Lista2)
{
	tlista* l1 = (tlista*)Lista1;
	tlista* l2 = (tlista*)Lista2;
	tnodo* temp;
	moveTostart(l1);
	moveTostart(l2);

	temp = l1->head;
	l1->head = l2->head;
	l2->head = temp;
	while(l1->curr->next != NULL)
	{
		temp = l1->curr->next;
		l1->curr->next = l2->curr->next;
		l2->curr->next = temp;
		next(l1);
		next(l2);
	}
	if(length(l1)%2 != 0)
	{
		temp = l1->tail;
		l1->tail = l2->tail;
		l2->tail = temp;
	}
}

/*
Funcion que si o si cambia el valor de un nodo al azar, y tambien puede cambiarle el tipo
Recibe la lista a la cual se quiere mutar.
*/
void mutacionRand(void* Lista){
	tlista* mutar = (tlista*)Lista;
	int k = rand()%(length(mutar));
	int i = 1;
	int* boi;
	moveTopos(mutar,k);
	boi = (int*)mutar->curr->data;
	k = rand()%3;
	if (k == 0){
		if (mutar->curr->type == 'i'){
			while(i){
				k = rand()%10;
				if (k != *boi){
					*boi = k;
					i = 0;
				}
			}
		}
		else{
			k = rand()%10;
			*boi = k;
			mutar->curr->type = 'i';
		}
	}
	else if (k == 1){
		if (mutar->curr->type == 'c'){
			while(i){
				k = 65 + rand()%6;
				if (k != *boi){
					*boi = k;
					i = 0;
				}
			}
		}
		else{
			k = 65 + rand()%6;
			*boi = k;
			mutar->curr->type = 'c';
		}
	}
	else{
		if (mutar->curr->type == 'b'){
			while(i){
				k = rand()%2;
				if (k != *boi){
					*boi = k;
					i = 0;
				}
			}
		}
		else{
			k = rand()%2;
			*boi = k;
			mutar->curr->type = 'b';
		}
	}
}

/*
Funcion que toma un nodo al azar y cambia si valor al azar.
Recibe la lista a la cual se desea mutar.
*/
void mutacionTipo(void* Lista){
	tlista* mutar = (tlista*) Lista;
	int k = rand()%(length(mutar));
	int i = 1;
	int* boi;
	moveTopos(mutar,k);
	boi = (int*)mutar->curr->data;
	if(mutar->curr->type == 'i'){
		while(i){
			k = rand()%10;
			if(k != *boi){
				*boi = k;
				i = 0;
			}
		}
	}
	else if(mutar->curr->type == 'c'){
		while(i){
			k = 65+rand()%6;
			if(k != *boi){
				*boi = k;
				i = 0;
			}
		}
	}
	else if(mutar->curr->type == 'b'){
		if (*boi == 0){
			*boi = 1;
		}
		else{
			*boi = 0;
		}
	}
}

/*
Funcion que le asigna un valor a la lista.
Recibe la lista que se desea evaluar, y el puntero a la funcion con la cual se desean evaluar los nodos.
Retorna el puntaje de la lista.
*/
int evaluacionLista(int(*fun)(void*), void* Lista)
{
	int puntaje, e, i;
	tlista* l = (tlista*)Lista;
	moveTostart(l);
	puntaje = 0;
	for(i = 0; i < length(l); i++)
	{
		e = fun(((tlista*)Lista)->curr);
		puntaje = puntaje + e;
		next(l);
	}
	return(puntaje);
}

/*
Funcion que, genera dos listas, evalua su calidad, luego crea una lista hija
para cada lista inicial, cruza ambas listas hijas, revisa si las nuevas listas son mayores
, en caso de que ambas sean mayores, los padres son reemplazados, al final muta los hijos salidos del
cruzamiento, si algun hijo posee mayor puntaje que su padre, lo reemplaza. Luego repite los pasos
anteriores "iteracion" veces, al final imprime las listas obtenidas y su puntaje.
*/
void genetico(void(*muta)(void*),void(*cruce)(void*, void*), int n, int iteraciones){
	int vp1, vp2, vh1, vh2, i;
	tlista* padre1 = generarSolucion(n);
	tlista* padre2 = generarSolucion(n);
	for(i = 0; i < iteraciones;i++){
		tlista* hijo1 = copiar(padre1);
		tlista* hijo2 = copiar(padre2);
		cruce(hijo1,hijo2);
		vp1 = evaluacionLista(fun,padre1);
		vp2 = evaluacionLista(fun,padre2);
		vh1 = evaluacionLista(fun,hijo1);
		vh2 = evaluacionLista(fun,hijo2);
		if(vp1 < vh1 && vp2 < vh2){
			vp1 = vh1;
			vp2 = vh2;
			borrar(padre1);
			borrar(padre2);
			padre1 = copiar(hijo1);
			padre2 = copiar(hijo2);	
			muta(hijo1);
			muta(hijo2);		
			vh1 = evaluacionLista(fun,hijo1);
			vh2 = evaluacionLista(fun,hijo2);
			if(vp1 < vh1){
				borrar(padre1);
				padre1 = copiar(hijo1);
				borrar(hijo1);
			}
			else{
				borrar(hijo1);
			}
			if(vp2 < vh2){
				borrar(padre2);
				padre2 = copiar(hijo2);
				borrar(hijo2);
			}
			else{
				borrar(hijo2);
			}
		}
		else{		
			muta(hijo1);
			muta(hijo2);
			vh1 = evaluacionLista(fun,hijo1);
			vh2 = evaluacionLista(fun,hijo2);
			if(vp1 < vh1){
				borrar(padre1);
				padre1 = copiar(hijo1);
				borrar(hijo1);
			}
			else{
				borrar(hijo1);
			}
			if(vp2 < vh2){
				borrar(padre2);
				padre2 = copiar(hijo2);
				borrar(hijo2);
			}
			else{
				borrar(hijo2);
			}		
		}
	}
	printf("Puntaje: %d Lista: ",evaluacionLista(fun,padre1));
	imprimirSolucion(padre1);
	printf("Puntaje: %d Lista: ",evaluacionLista(fun,padre2));
	imprimirSolucion(padre2);
	borrar(padre1);
	borrar(padre2);
}