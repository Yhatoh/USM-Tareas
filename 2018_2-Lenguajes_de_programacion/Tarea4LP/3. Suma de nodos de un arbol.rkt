#lang scheme

#|
Funcion: suma_arbol
Descripcion: busca un elemento dentro de un arbol binario y suma todos los numeros por los que paso
             para encontrar ese elemento
Parametros: arbol numero
arbol arbol el cual se le buscara el numero
numero entero a buscar
Retorno: la suma de los elementos por los que paso para encontrar el numero (contandose) si este no se
         encuentra retorna null.
|#
(define (suma_arbol arbol numero)
  (let suma ((tree arbol) (sumat 0))
    (cond
      ((null? tree) null)
      ((= numero (car tree)) (+ sumat (car tree)))
      ((< numero (car tree)) (suma (car (cdr tree)) (+ sumat (car tree))))
      ((> numero (car tree)) (suma (car (cdr (cdr tree))) (+ sumat (car tree))))        
    )
  )
)


