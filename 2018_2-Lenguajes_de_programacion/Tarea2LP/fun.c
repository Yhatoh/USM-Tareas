#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "TDALista.h"

int fun(void* node)
{
	tnodo* nodo = (tnodo*) node;
	char type = nodo->type;
	int a = 0;
	if(type == 'i')
	{
		a += *(int*)nodo->data;
		
	}
	else if(type == 'c')
	{
		a = a/(*((int*)nodo->data)); //*(int*)nodo->data;
	}
	else
	{
		if(*(int*)nodo->data == 0)
		{
			a -= 1;
		}
		else
		{
			a += 1;
		}
	}
	return a;
}