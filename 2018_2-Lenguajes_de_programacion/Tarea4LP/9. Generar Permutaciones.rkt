#lang scheme


#|
Funcion: insert
Descripcion: Añade un objeto (lista, entero, etc) en una pocicion determinada de la lista
Parametros:
lista lista
num obj
pos entero
Retorno: Retorna la lista on el objeto insertado.
|#

(define (insert lista num pos)
  (let in ((new '()) (iter 0) (l1 lista))
    (if (not (= iter pos))
        (in (append new (list (car l1))) (+ iter 1) (cdr l1))
        (append new (append (list num) l1))
    )
  )
)


#|
Funcion: allin
Descripcion: Coloca un obj en todas las posiciones de una lista, utilizando la funcion insert
Parametros:
lista lista
numero obj
Retorno: Lista que posee todas las listas, cada una con el obj añadido en una posicion diferente
|#

(define (allin lista num )
  (let all ((new '()) (iter 0) (l1 lista))
    (if (not (= iter (+ (length lista) 1)))
       (all (append new (list(insert l1 num iter))) (+ iter 1) lista)
        new
        )
    )
)


#|
Funcion: alllist
Descripcion: Aplica la funcion allin, a todas las listas de una lista
Parametros:
lista lista
num obj
Retorno: La lista con las listas a las cuales se le aplico la funcion allin.
|#

(define (alllist lista num)
  (let all ((new '()) (l1 lista) )
    (cond 
        ((null? l1) new)
        ((null? num) l1) 
        (else (all (append new (allin (car l1) num)) (cdr l1)))
    )
  )
 ) 


#|
Funcion: permutaciones
Descripcion: Encuentra todas las permutaciones posibles de una lista, utilizando la funcion allist
Parametros:
lista lista
Retorno: lista que posee todas las permutaciones posibles de la lista original.
|#
(define (permutaciones lista)
  (let per ((next lista) (prev '()))
    (if (= (length next) 1)
        (allin next prev)
        (alllist (per (cdr next) (car next)) prev)
        )
    )
)
