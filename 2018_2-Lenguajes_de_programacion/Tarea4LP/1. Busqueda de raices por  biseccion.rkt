#lang scheme

#|
Funcion: biseccion
Descripcion: Encuentra el valor (o el valor mas cercano) en la cual cierta funcion se vuelve 0, en funcion de un intervalo
Parametros:
funcion funcion
itervalo lista con dos enteros
iter entero
Retorno: retorna el valor mas cercano en el cual la funcion se vuelve 0, o null si el itervalo es incorrecto.
|#

(define ((biseccion funcion) intervalo iter)
  (let bis ((medio (/ (+ (car intervalo) (car(cdr intervalo))) 2)) (ite iter) (a (car intervalo)) (b (car(cdr intervalo))))
    (if ( < 0 (* (funcion (car intervalo)) (funcion (car(cdr intervalo)))))
        null
        (if (= ite 0)
            medio
            (if (= 0 (funcion medio))
                medio
                (if (< 0 (* (funcion medio) (funcion a)))
                    (bis(/ (+ medio b) 2) (- ite 1) medio b)
                    (bis(/ (+ medio a) 2) (- ite 1) a medio)
                )
            )
        )
    )
  )
)


