#lang scheme

#|
Funcion: calculo
Descripcion: a una expresion le calculo el costo total de esta por los elementos que contenga
Parametros: costs expr
costs costos de elementos
expr lista con elementos a calcular su costo
Retorno: la suma de los costos de cada elemento de la lista.
|#
(define (calculo costs expr)
  (let kha ((total 0) (frase expr))
    (cond
      ((null? frase) total)
      ((assq (car frase) costs) (kha (+ total (cdr (assq (car frase) costs))) (cdr frase)))
      (else (kha total (cdr frase)))
    )
  )
)


