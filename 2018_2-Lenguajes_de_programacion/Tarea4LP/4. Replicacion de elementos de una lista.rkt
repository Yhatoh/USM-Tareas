#lang scheme

#|
Funcion: replicar_aux
Descripcion: a la lista proveniente de la funcionar replicar le agrega el valor n veces a la lista entregada
Parametros: lista valor iter
lista con elementos
valor elemento a agregar
iter la cantidad de veces que se debe agregar el elemento
Retorno: la lista con los elementos agregados n veces.
|#
(define (replicar_aux lista valor iter)
  (let multi ((lista1 lista) (n iter))
    (if (= n 0)
        lista1
        (multi (append lista1 (list valor)) (- n 1))
    )
  )
)

#|
Funcion: replicar
Descripcion: recibe dos listas de mismo tamaño a la cual a la primera lista se le replicaran los elementos
             n veces donde n corresponde a los elementos de la segunda lista
Parametros: lista lista_replicacion
lista: lista con elementos a replicar
lista_replicacion: lista con la cantidad de veces que se debe replicar cada elementos
Retorno: entrega una lista con los elementos replicados.
|#
(define (replicar lista lista_replicacion)
  (let add ((lista1 lista) (lista2 lista_replicacion) (resultado '()))
    (if (null? lista1)
        resultado
        (add (cdr lista1) (cdr lista2) (replicar_aux resultado (car lista1) (car lista2)))
    )
  )
)

