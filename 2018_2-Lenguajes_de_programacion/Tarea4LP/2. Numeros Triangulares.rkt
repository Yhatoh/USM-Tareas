#lang scheme

#|
Funcion: triangular_cola
Descripcion: Comprueba si un numero es triangular, con recurcion de cola
Parametros:
numero entero
Retorno: #t si el numero es triangular, #f si el numero no es triangular.
|#


(define (triangular_cola numero)
  (let tri ((sum 0) (cont 1))
    (cond
      ((= numero 0) #f)
      ((< numero sum) #f)
      ((= numero sum) #t)
      (else (tri (+ sum cont) (+ cont 1)))
      )
    )
  )


#|
Funcion: triangular_simple
Descripcion: Comprueba si un numero es triangular, con recursion simple
Parametros:
numero entero
Retorno: #t si el numero es triangular, #f si el numero no es triangular.
|#

(define (triangular_simple numero)
  (let tri ((sum 0) (cont 1))
    (cond
      ((= numero 0) #f)
      ((< numero sum) #f)
      ((= numero sum) #t)
      (else (and (tri (+ sum cont) (+ cont 1)) (= (/ pi pi) 1)))
    )
  )
)
