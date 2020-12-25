#lang racket

#|
Funcion: operar
Descripcion: Aplica la funcion func entre a dos elementos de una lista binaria
Parametros: func lista1 lista2
func funcion
lista1 lista
lista2 lista
Retorno: Lista con el resultado de func aplicado a cada elemento de la lista en orden.
|#
(define ((operar func) lista1 lista2)
  (let op ((new '()) (l1 lista1) (l2 lista2))
    (if (null? l1)
        new
        (op (append new (list(func (car l1) (car l2)))) (cdr l1) (cdr l2))
        )
    )
  )
