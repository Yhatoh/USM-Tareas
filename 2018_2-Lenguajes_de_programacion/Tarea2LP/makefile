all: genetic

genetic: main.o genetico.o fun.o TDALista.o
	gcc -o genetic TDALista.o main.o genetico.o fun.o

main.o: main.c
	gcc -c main.c -Wall

TDALista.o: TDALista.c TDALista.h
	gcc -c TDALista.c TDALista.h -Wall

genetico.o: genetico.c genetico.h
	gcc -c genetico.c genetico.h -Wall

fun.o: fun.c fun.h
	gcc -c fun.c fun.h -Wall

clean:
	rm main.o genetico.o fun.o TDALista.o genetic