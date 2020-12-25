%Para la consulta se debe realizar de la siguiente forma:
%ciudades(C,G).
%G es la cantidad de gasolina disponible
%C es la ciudad a la cual se quiere saber a cuales ciudades se pueden llegar con una gasolina G

%Adyacentes a pallet town
adyacente(pallet_town,viridian_city,6).
%Adyacentes a viridian city
adyacente(viridian_city,pewter_city,28).
adyacente(viridian_city,new_bark_town,7).
%Adyacentes a pewter city
adyacente(pewter_city,viridian_city,26).
adyacente(pewter_city,cerulean_city,29).
%Adyacentes a cerulean city
adyacente(cerulean_city,lavender_town,9).
%Adyacentes a vermilion city
adyacente(vermilion_city,fuchsia_city,42).
adyacente(vermilion_city,lavender_town,28).
%Adyacentes a saffron city
adyacente(saffron_city,cerulean_city,34).
adyacente(saffron_city,celadon_city,17).
adyacente(saffron_city,vermilion_city,16).
%Adyacentes a lavender town
adyacente(lavender_town,saffron_city,33).
adyacente(lavender_town,fuchsia_city,58).
%Adyacentes a celadon city
adyacente(celadon_city,saffron_city,20).
%Adyacentes a cinnabar island
adyacente(cinnabar_island,pallet_town,47).
adyacente(cinnabar_island,fuchsia_city,35).
%Adyacentes a fuchsia city
adyacente(fuchsia_city,cinnabar_island,35).
adyacente(fuchsia_city,celadon_city,34).
%Adyacentes a new_bark_town
adyacente(new_bark_town,pallet_town,4).
adyacente(new_bark_town,cherrygrove_city,13).
%Adyacentes a cherrygrove city
adyacente(cherrygrove_city,new_bark_town,10).
adyacente(cherrygrove_city,violet_city,29).
%Adyacentes a blackthorn city
adyacente(blackthorn_city,new_bark_town,35).
adyacente(blackthorn_city,cherrygrove_city,25).
%Adyacentes a violet city
adyacente(violet_city,mahogany_town,22).
%Adyacentes a mahogany town
adyacente(mahogany_town,blackthorn_city,44).

%Regla que permite determinar a que ciudades puedo llegar desde otra ciudad con una cantidad de gasolina G
viaje(C1,C2,G):- 
    adyacente(C1,C2,X), 
    X =< G;
    adyacente(C1,Y,Z), 
    Z =< G, 
    viaje(Y,C2,G-Z).
%Toma una lista y elimina los elementos repetidos dentro de ella
eliminar(_,[],L,L):-!.
eliminar(C1,L,L1,L4):- 
    L = [X|L2],
    not(member(C1,[X])),
    not(member(X,L1)),
    append(L1,[X],L3),
    eliminar(C1,L2,L3,L4);
    L = [_|L2],
    eliminar(C1,L2,L1,L4).
%Imprime los elementos de una lista en este caso imprime las ciudades
imprimir([]).
imprimir(L):-
    L = [X|L1],
    write(X),nl,
    imprimir(L1).
%Crea una lista con todas las ciudades alcanzables por todos los caminos posibles
%Luego elimina todas las ciudades repetidas dentro de esa lista y luego las imprime
ciudades(C1,G):- 
    findall(X,viaje(C1,X,G),L1),
    eliminar(C1,L1,[],L),
    write("Las ciudades visitables son:"),nl,
    imprimir(L),!.

