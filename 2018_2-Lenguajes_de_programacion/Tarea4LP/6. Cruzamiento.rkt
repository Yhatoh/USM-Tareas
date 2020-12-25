#lang scheme

#|
Funcion: cruzar
Descripcion: intercambia los n primeros elementos de dos listas
Parametros: lista1 lista2 n
lista1 lista de elementos
lista2 lista de elementos
n entero que servira para intercambiar los elementos
Retorno: una lista con las dos listas donde sus primeros n elementos.
|#
(define (cruzar lista1 lista2 n)
  (let cruce ((l1 lista1) (l2 lista2) (r1 '()) (r2 '()) (largo 0))
    (if (null? l1)
        (list r1 r2)
        (if (< largo n)
            (cruce (cdr l1) (cdr l2) (append r1 (list (car l2))) (append r2 (list (car l1))) (+ largo 1))
            (cruce (cdr l1) (cdr l2) (append r1 (list (car l1))) (append r2 (list (car l2))) (+ largo 1))
        )
    )
  )
)
