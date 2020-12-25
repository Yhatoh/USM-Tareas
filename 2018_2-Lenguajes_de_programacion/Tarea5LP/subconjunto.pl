/**
 * La consulta se debe realizar de la siguiente forma:
 * acum(l,M,F).
 * L es una lista que posee numeros enteros positivos no repetidos
 * un ejemplo de L seria [1,2,3].
 * M es el entero positivo al cual se le quiere encontrar los subconjuntos
 * un ejemplo de M seria 4.
 * F es la variable que obtendra el valor de la lista de subconjuntos que suman M.
 */

%Entrega todos los subconjuntos de una lista
subset([], []).
subset([X|Y], [X|Z]):-subset(Y, Z).
subset([_|Y], Z):-subset(Y, Z).

%Revisa cuales listas que se encuentran en una lista, suman M, si lo hacen, se guarda en otra lista.
suma([X|[]],Z,H,Q):- sum_list(X,T), T=Z, append(H,[X],A), Q=A,!;
    Q=H.
suma([X|L],Z,H,Q):-sum_list(X,T), T=Z, append(H,[X],A),suma(L,Z,A,Q),!;
    suma(L,Z,H,Q),!.

%Consigue todos los valores de subset dentro de una lista, y luego se la entrega a suma.
acum(W,Z,Q):- findall(X,subset(W,X),L), suma(L,Z,[],Q).




