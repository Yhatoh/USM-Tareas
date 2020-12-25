InFileName: .asciz "input.txt"

.align
Myarray: .skip 1000*4
InFileHandle:.word 0
Nextline: .asciiz "\n"
mensajito: .asciiz "Encontre el numero en la posicion del numero es "
mensajito2: .asciiz "No lo encontre :c"


MAIN:
    ldr r0,=InFileName      @De la linea 12 a 14
    mov r1,#0               @Se abre el archivo input.txt
    swi 0x66                
    ldr r1,=InFileHandle    @De la linea 15 a 19
    str r0,[r1]             @Se lee el primer INT del archivo
    ldr r0,=InFileHandle    @El cual corresponde a la cantidad
    ldr r0,[r0]             @de elementos dentro del arreglo
    swi 0x6c                @y se almacena en r0

    mov r10, r0             @r10 va a ser igual al tamaño del arreglo

    mov r2, #0              @r2 nos servira de contador para leer los 
                            @n elemento y guardalos

    ldr r3, =Myarray        @r3 tendra la direccion base de nuestro arreglo

@Este bloque realizara la accion de un for
FOR:
    cmp r2, r10             @Aqui verifica que el contador no sea igual al
    beq MAIN2               @tamaño del arreglo si es asi lei todos los datos 
                            @y va a main2

    ldr r0,=InFileHandle    @Permite leer un entero por iteracion y alamenarlo
    ldr r0,[r0]             @en r0
    swi 0x6c            

    mov r5, r2, lsl #2      @Aqui r5 sera el indice dentro del arreglo basado en
                            @r2

    str r0, [r3, r5]        @Guarda el entero leido anteriomente en la posicion
                            @r2

    add r2, r2, #1          @suma 1 a r2, el contador

    b FOR                   @se devuelve al FOR

MAIN2:
    ldr r0,=InFileHandle    @aqui se lee el entero K a buscar y se alamcena en
    ldr r0,[r0]             @r0
    swi 0x6c

    sub r6, r10, #1         @r6 sera la pos del valor maximo de mi intervalo de busqueda
    mov r4, #0              @r4 sera la pos del valor minimo de mi intervalo de busqueda

    mov r5, r6, asr #1      @r5 correspondera a la mitad del intervalo redondeado
                            @para abajo, correspondera a la posicion a revisar

    mov r7, r0              @r7 sera el valor a buscar

    mov r9,#0               @r9 sera un contador que nos servira como condicion de
                            @ salida más adelante

BINARIOSII:
    mov r2, r5, lsl #2      @aqui defino cuanto debo sumarle a la direccion de memoria
                            @base para encontrar el elemento en la posicion r5

    ldr r1, [r3,r2]         @se alamcena en r1 el elemento en la posicion r5}
    cmp r7,r1               @se compara r7 (valor a buscar) y r1(valor actual)
    beq LOENCONTRE          @si lo encuentro lo manda a LOENCONTRE
    bgt ESMAYOR             @si r7 es mayor a r1 lo manda a ESMAYOR
    blt ESMENOR             @si r7 es menor a r1 lo manda a ESMENOR

ESMAYOR:
    mov r4, r5              @como r7 es mayor al elemento en posicion r5
                            @r5 pasa a ser r4, el nuevo minimo del intervalo 

    add r5, r5, r6          @estas dos lineas son para sacar el valor medio de la mitad
    mov r5, r5, asr #1      @del nuevo intervalo obtenido 

    add r8, r5, #1          @esto era para verificar si el elemento a verificar es el 
    cmp r8, r6              @ultimo elemento del arreglo, entonces esto arregla un errosilo
    moveq r5, r6            @por el redondeo inferior, entonces r5 pasa a ser el ultimo 
                            @elemento del arreglo

    addeq r9, r9,#1         @ahora las siguientes 3 lineas son para verificar si la cuestion
    cmp r9, #2              @se quedo en un loop debido a que el emento no esta en el arreglo
    beq NOENCONTRE          @y esta oscilando entre x1 elemento y x2 elemento, donde x2 = 1+x1
                            @aqui la utilidad del contador r9 donde si pasa 2 veces por el 
                            @mismo valor significa que sucedio el loop ergo el elemento no esta

    cmp r8, r10             @y esta condicion es por si me pase del largo del arreglo
    beq NOENCONTRE          @ergo el elemento no esta en el arreglo y lo mando a NOENCONTRE
    
    b BINARIOSII            @si sobrevive a todas las condiciones la busqueda continua

ESMENOR:
    mov r6, r5              @como el elemento r7 es menor al elemento en pos r5
                            @la pos r5 pasa a ser la pos del valor maximo del intervalo a 
                            @buscar

    add r5, r5, r4          @aqui saca el promedio entre la pos maxima y la pos minimo
    mov r5, r5, asr #1

    cmp r4, r6              @si pasa a la funcion menor y r4 y r6 son iguales, ergo
    beq NOENCONTRE          @el elemento no esta en el arreglo lo manda a NO ENCONTRE
    
    b BINARIOSII            @si no sigue la busqueda

LOENCONTRE:
    ldr r1, =mensajito      @esta seccion imprime un mensaje con la posicion donde
    mov r0, #1              @se encuentra el elemento K
    swi 0x69
    mov r1, r5
    swi 0x6b
    beq DONE

NOENCONTRE:
    ldr r1, =mensajito2     @aqui imprime que no se encontro el valor
    mov r0, #1
    swi 0x69
    beq DONE

DONE:
    ldr r0, =InFileHandle   @se cierra el archivo input.txt y termina el programa
    ldr r0,[r0]
    swi 0x68
    swi 0x11
