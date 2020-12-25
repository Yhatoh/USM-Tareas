
void* generarSolucion(int largo);
void* copiar(void* Lista);
void borrar(void* Lista);
void imprimirSolucion(void* Lista);
void cruceMedio(void* Lista1, void* Lista2);
void cruceIntercalado(void* Lista1, void* Lista2);
void mutacionRand(void* Lista);
void mutacionTipo(void* Lista);
int evaluacionLista(int (*fun)(void*), void* Lista);
void genetico(void (*muta)(void*), void (*cruce)(void*,void*), int n, int iteraciones);