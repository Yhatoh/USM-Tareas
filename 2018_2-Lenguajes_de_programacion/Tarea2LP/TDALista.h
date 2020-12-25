/*
Estructura que representa un nodo de la lista.
La estructura posee:
void* data: Dato que puede ser un int o char
char type: Especifica si el dato es un int "i", un char "c" o bit "b" el cual esta representado como un int.
struct tnodo *next: Puntero a la siguiente posicion de la lista
*/
typedef struct tnodo{
	void* data;
	char type; //'i' 'c' 'b'
	struct tnodo *next;
}tnodo;

/*
Estructura que repesenta la lista.
La estructura posee:
int listsize: int que guarda el largo de la lista.
tnodo *head: Puntero al inicio
tnodo *tail: Puntero al final
tnodo *curr: Puntero al nodo actual
*/
typedef struct tlista{
	int listsize;
	tnodo *head;
	tnodo *tail;
	tnodo *curr;
	int pos;
}tlista;


tnodo *iniNodo(int* vali, char tipo);
tlista *inicializar();
tlista *insertar_pos_actual(tlista *lista,int* vali,char tipo);
void moveTostart(tlista *lista);
void moveToend(tlista *lista);
void next(tlista *lista);
void moveTopos(tlista *lista, int pos_n);
int length(tlista *lista);