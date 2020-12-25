%Para la consulta se debe realizar de la siguiente forma:
%diccionario(L)
%L es diccionario a utilizar, lo ideal seria usar [], pero si quieres empezar con un diccionario ya creado.
%el formato de diccionario utilizado fue [[llave,[valor]]
%
%cuando haga la consulta se le preguntara que quieres hacer y puedes escribir las siguientes 5:
%   -agregar: para agregar una llave con su valor correspondiente al diccionario
%   -largo: te entrega el largo del diccionario (cantidad de llaves en el)
%   -valor: te entrega el valor de una llave pedida
%   -llaves: te entrega todas las llaves del diccionario
%   -valores: te entrega todos los valores de cada llave del diccionario
%   -stop: para terminar de hacer consultas
%   
%Nota: las llaves no pueden contener caracteres reservados de prolog
diccionario(L):- 
    write("Que quieres hacer? "),
    read(X),
    accion(X,L),!.

accion(stop,L):- write("Diccionario final "), nl, write(L),!.
accion(agregar,L):- 
    write("Valor a agregar: "),
    read(X),
    write("Llave correspondiente: "),
    read(Y),
    agregar(L,X,Y).
accion(largo,L):-
    largo(L).
accion(llaves,L):-
    llaves(L).
accion(valor,L):-
    write("Llave al cual quieres conocer su valor: "),
    read(X),
    valor(L,X).
accion(valores,L):-
    write("Los valores de tu diccionarios son los siguientes "),
    valores(L,[]).
        
%agregar a dicc
agregar(L,X,Y):- 
    not(member([Y|_],L)),
    append(L,[[Y|[[X]]]],R),
    write("Diccionario actual: "), nl, write(R), nl,
    diccionario(R);
    buscar(L,Y,R1),
    append(R1,[[Y|[[X]]]],R),
    write("Diccionario actual: "), nl, write(R), nl,
    diccionario(R).
buscar(L,Y,R1):-
    L = [X|R1],
    X = [K|_],
    Y = K;
    L = [A|R2],
    append(R2,[A],NL),
    buscar(NL,Y,R1).

%obtener largo del dicc
largo(L):-
    largo_aux(L,0),
    diccionario(L).
largo_aux(L,0):-
    L = [],
    write("El largo del diccionario es 0"),nl;
    L = [_|L1],
    largo_aux(L1,1).
largo_aux([],N):- 
    write("El largo del diccionario es "), write(N), nl.
largo_aux(L,N):-
    L = [_|L1],
    N1 is N + 1,
    largo_aux(L1,N1).

%obtener llaves del dicc
llaves(L):- 
    llaves_aux(L,[]),
    diccionario(L).
llaves_aux(L,[]):-
    L = [],
    write("EL diccionario no tiene llaves"),nl;
    L = [X|L1],
    X = [Y|_],
    append([],[Y],K),
    llaves_aux(L1,K).
llaves_aux([],K):-
    write("El diccionario tiene las siguientes llaves: "), write(K), nl.
llaves_aux(L,K):-
    L = [X|L1],
    X = [Y|_],
    append(K,[Y],K2),
    llaves_aux(L1,K2).

%obtener valor
valor(L,X):-
    existe(L,X,[L1]),
    write("El valor de la llave "), write(X), write(" es "), write(L1),nl,
    diccionario(L);
    write("La llave pedida no existe"),nl,
    diccionario(L).
existe([],_,_):- false.
existe(L,X,L2):-
    L = [Y|_],
    Y = [A|[L2]],
    A = X;
    L = [_|L1],
    existe(L1,X,L2).

%obtener todos los valores.
valores(L,_):-
    L = [X|L1],
    X = [_|[[K]]],
    write(" "),write(K),write(" "),
    valores(L1,L).
valores([],L):- 
    nl,
    diccionario(L).
    

