%Para la consulta se debe realizar de la siguiente forma:
%accept(P)
%P es la palabra a escribir
%Ejemplos de formato de palabra:
%aaaaaaaaaaaaaaaa
%"ababababab"
%"abcaskhasdj"
%Puede realmente escribir cualquier caracter entre comillas, sin comillas puede escribir cualquier caracter
%que pertenesca al abecedario y sin mayusculas al principio

%alfabeto
sim(a).
sim(b).
sim(c).
%estado 0
ady(0,1,a).
ady(0,7,b).
ady(0,4,c).
%estado 1
ady(1,2,b).
%estado 2
ady(2,3,b).
%estado 3
ady(3,1,a).
ady(3,4,c).
%estado 4
ady(4,5,a).
%estado 5
ady(5,6,b).
%estado 6
ady(6,0,a).
%estado 7
ady(7,8,a).
%estado 8
ady(8,9,c).
%estado 9
ady(9,0,a).

%Comprueba que existe un camino de estados tal que se acabe la palabra y llegue a aceptacion
accept(P,X,X,"No es parte del lenguaje"):- P = [Y|_], not(sim(Y)).
accept([],0,0,"Es parte del lenguaje").
accept([],3,3,"Es parte del lenguaje").
accept([],X,X,"No es parte del lenguaje"):- X =\= 3, X =\= 0.
accept(P,X,F,S):- 
    P = [Y|L],
    sim(Y),
    ady(X,A,Y),
    accept(L,A,F,S),!.
accept(_,X,X,"No es parte del lenguaje").
accept(L):-
    string_chars(L,P),
    accept(P,0,F,S),
    write("Paro en este estado "),write(F),nl,
    write(S),!.
