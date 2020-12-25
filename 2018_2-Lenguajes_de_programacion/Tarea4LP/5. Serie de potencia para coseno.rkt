#lang scheme

#|
Funcion: factorial
Descripcion: calcula el factorial de un numero
Parametros: n
n entero
Retorno: el factorial del elemento n.
|#
(define (factorial n)
  (let fac ((val 1) (f n))
    (if (= f 0)
        val
        (fac (* val f) (- f 1))
        )
    ) 
  )

#|
Funcion: operacion
Descripcion: a partir de un elemento n y valor x realiza una operacion ((-1)^n)*(x^(2n))/((2n)!)
Parametros: n x
n entero
x numero real
Retorno: el valor de la operacion dicha en la descripcion.
|#
(define (operacion n x)
  (/ (* (expt -1 n) (expt x (* 2 n))) (factorial (* 2.0 n)))
  )

#|
Funcion: rec_cos
Descripcion: realiza el calculo aproximado utilizando el teorema de Maclaurin de cos(x), utilizando recursion de cola
Parametros: valor iter
valor numero real
iter entero
Retorno: retorna el valor de la aplicacion del teorema de maclaurin de la sumatoria de 0 a (n-1).
|#
(define (rec_cos valor iter)
  (let rec ((val 0) (ite iter))
    (if (= ite 0)
        val
        (rec (+ val (operacion (- ite 1) valor)) (- ite 1))
        )
    )
  )

#|
Funcion: rec_iter
Descripcion: realiza el calculo aproximado utilizando el teorema de Maclaurin de cos(x), a traves de la iteracion
Parametros: valor iter
valor numero real
iter entero
Retorno: retorna el valor de la aplicacion del teorema de maclaurin de la sumatoria de 0 a (n-1).
|#
(define (iter_cos valor iter)
  (do ((x (- iter 1) (- x 1)) (y 0 (+ y (operacion x valor)))
    )((= x -1) y))
  )
