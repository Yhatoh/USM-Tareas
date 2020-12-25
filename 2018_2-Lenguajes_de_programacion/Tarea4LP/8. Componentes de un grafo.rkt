#lang scheme

#|
Funcion: componentes
Descripcion: determina cuantas componentes conexas tiene un grafo
Parametros: grafo
grafo es una lista de adyacencia que representa un grafo
Retorno: la cantidad de componentes conexas del grafo dado.
|#
(define (componentes grafo)
  (let compo ((g grafo) (super '())) 
    (if (null? g)
        (length (fusion_ha super))
        (compo (cdr g) (append super (list (reescribir '() (car g)))))
    )
  )                       
)

#|
Funcion: fusion_ha
Descripcion: a partir de una lista se crea una segunda lista donde esta contendra la fusion de las sublistas de la lista
             si es que estos tiene algun elemento en comun
Parametros: l
l una lista que representa un grafo escrito de otra forma
Retorno: retorna una lista la cual contiene sublistas fusionadas.
|#
(define (fusion_ha l)
  (let haaaa ((grafo l) (fusion '()))
    (cond
      ((null? grafo) fusion)
      ((null? fusion) (haaaa (cdr grafo) (append fusion (list (car grafo)))))
      ((pert_1_lista (car grafo) fusion) (haaaa (cdr grafo) (añadir2 (car grafo) fusion)))
      (else (haaaa (cdr grafo) (append fusion (list (car grafo)))))
    )    
  ) 
)

#|
Funcion: pert_1_lista
Descripcion: recibe una lista y una lista de sublistas, verifica si la lista tiene algun elemento en comun
             con alguna sublista
Parametros: g f
g una lista
f una lista con sublistas
Retorno: retorna verdadero si es que hay una sublista en f que tenga un elemento en comun con la lista g si no retorna falso.
|#
(define (pert_1_lista g f)
  (let pert ((gra g) (fu f))
    (cond
      ((null? fu) #f)
      ((null? gra) (pert g (cdr fu)))
      ((member (car gra) (car fu)) #t)
      (else (pert (cdr gra) fu))
    )
  )
)

#|
Funcion: añadir2
Descripcion: se recibe una lista y una lista con sublistas, a esta la lista con sublistas se busca cual sublista
             tiene un elemento en comun con la lista entregada, si los tiene se añade a una nueva lista elemento las
             dos listas juntadas si no la sublista se agrega sin efectuar cambio a la nueva lista
Parametros: l1 l2
l1 lista
l2 lista con sublistas
Retorno: retorna una lista de sublistas donde una de las sublistas anteriores contiene la lista entregada (l1).
|#
(define (añadir2 l1 l2)
  (let add2 ((l l1) (f l2) (nueva '()) (logrado 0))
    (cond
      ((and (= logrado 1) (null? f)) nueva)
      ((= logrado 1) (add2 l1 (cdr f) (append nueva (list (car f))) 1))
      ((and (= logrado 0) (null? f)) (add2 (cdr l) l2 '() 0))
      ((null? l) nueva)
      ((member (car l) (car f)) (add2 l1 (cdr f) (append nueva (list (append (car f) l1))) 1))
      (else (add2 l (cdr f) (append nueva (list (car f))) 0))                        
    )
  )
)

#|
Funcion: reescribir
Descripcion: toma un grafo del problema y lo reescribe quitandole las parentesis que tiene
Parametros: a g
a '()
g lista de adyancencia de un nodo
Retorno: agrega a la lista a los elementos de la lista de adyacencia del nodo.
         Ejemplo: g = (1 (2 3)) -> retorno a = (1 2 3)
|#
(define (reescribir a g)
  (let add ((lista a) (nodo g) (iter 0))
    (cond
      ((null? nodo) lista)
      ((= iter 0) (add (append lista (list (car nodo))) (car (cdr nodo)) 1))
      ((= iter 1) (add (append lista (list (car nodo))) (cdr nodo) 1))
    )
  ) 
)
               

