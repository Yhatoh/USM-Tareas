#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <time.h>
#include "TDALista.h"

/*
Funcion que inicializa un nodo.
Recibe un puntero a entero, y un char que representa el tipo de dato.
Retorna el nodo inicializado.
*/
tnodo *iniNodo(int* vali, char tipo){
	tnodo *nodo = (tnodo*)malloc(sizeof(tnodo));
	nodo->type = tipo;
	nodo->next = NULL;
	nodo->data = vali;
	return nodo;
}

/*
Funcion para inicializar una lista.
Crea una lista nueva (vacia) con memoria dinamica
Retorna el puntero de la lista
*/
tlista *inicializar(){
	tlista *nueva = (tlista*)malloc(sizeof(tlista));
	nueva->curr = nueva->tail = nueva->head = NULL;
	nueva->pos = 0;
	nueva->listsize = 0;
	return nueva;
}

/*
Funcion para insertar un nodo en la posicion actual.
Recibe la lista a la cual se le insertara el nodo, el valor, y el tipo del dato.
Retorna la lista con el nodo insertado.
*/
tlista *insertar_pos_actual(tlista *lista,int* vali,char tipo){
	if (lista->listsize == 0){
		lista->curr = lista->tail = lista->head = iniNodo(vali,tipo);
		lista->listsize++;
	}
	else{
		tnodo *aux = lista->curr->next;
		lista->curr->next = iniNodo(vali, tipo);
		lista->curr->next->next = aux;
		if(lista->curr == lista->tail){
			lista->tail = lista->curr->next;
		}
		lista->listsize++;
	}
	return lista;
}

/*
Funcion que cambia el puntero de curr de una lista, al principio de la lista.
Recibe la lista que se desea cambiar el puntero.
*/
void moveTostart(tlista *lista){
	lista->curr = lista->head;
	lista->pos = 0;
}

/*
Funcion que cambia el puntero de curr de una lista, al final de la lista.
Recibe la lista que se desea cambiar el puntero.
*/
void moveToend(tlista *lista){
	lista->curr = lista->tail;
	lista->pos = (lista->listsize)-1;
}


/*
Funcion que cambia el puntero de curr de una lista, al siguiente nodo.
Recibe la lista que se desea cambiar el puntero.
*/
void next(tlista *lista){
	if (lista->curr != lista->tail){
		lista->curr = lista->curr->next;
		lista->pos += 1;
	}
}

/*
Funcion que cambia el puntero de curr de una lista, a una posicion deseada.
Recibe la lista que se desea cambiar el puntero.
*/
void moveTopos(tlista *lista, int pos_n){
	int i;
	if ((pos_n >= 0) && (pos_n < lista->listsize)){
		lista->curr = lista->head;
		for (i = 0; i < pos_n; i++){
			lista->curr = lista->curr->next;
		}
		lista->pos = pos_n;
	}
}

/*
Funcion que retorna el largo de una lista.
Recibe la lista que se decea conocer el largo.
Retorna el valor del largo de la lista.
*/
int length(tlista *lista){
	return lista->listsize;
}
