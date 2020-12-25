Nombres:    Gabriel Carmona // Jorge Ludueña
Roles:      201773509-0     // 201773507-4

ATENCION: VIENE CON ARCHIVOS DE main.c y fun.c de prueba recomiendo modificarlo

Instrucciones de compilacion:

-A la estructura del nodo dada se le añadio un dato más quedando de la siguiente forma:

	typedef struct tnodo{
		void* data;
		char type; //'i' 'c' 'b'
		struct tnodo *next;
	}tnodo;

-Para realizar a compilacion debe tener un carpeta que contenga los siguientes archivos:
	!- TDALista.h, TDALista.c, genetico.c y genetico.h, los cuales seran archivos que estaran en el archivo mandado.
	!- fun.c, que contiene la funcion para evaluar las listas, para el codigo genetico.c se asumio que la funcion evaluar se llama "fun" y recibe como parametro un void* nodo 
	   Notar que fun.c debe tener #include TDALista.h y no tener la estructura tnodo en fun.c.
	!- fun.h, el cual es el header de la funcion fun.c
	!- main.c, el cual contiene el main del programa y debe incluir #include genetico.h
	!- makefile, el cual servira para la compilacion del programa.

-Ahora para la compilación abrir la consola en la carpeta y poner el comando make all
 Para borrar los archivos .o y el ejecutable puede usar make clean


