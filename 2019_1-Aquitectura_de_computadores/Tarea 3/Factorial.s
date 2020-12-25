
InFileName: .asciz "input.txt"        ;Nombre del archivo donde viene el input   
.align

InFileHandle:.word 0                  ;Asignacion de memoria para la lectura

LDR R0,=InFileName
                    ;Se carga el nombre del archivo en R0
MOV R1,#0
    
SWI 0x66                              ;Se abre el archivo

LDR R1,=InFileHandle
                  ;Se carga la asignacion de memoria en R1
STR R0,[R1]                           ;Se carga en R0 la direccion de R1

LDR R0,=InFileHandle
                  ;Se carga la asignacion de memoria en R0
LDR R0,[R0]
                           ;Se guarda el numero leido del archivo en R0  
SWI 0x6c                              ;Se lee el entero en R0

MOV R1 , #1                           ;Se copia a R1 el primer valor de los factoriales permitidos (1)
CMP R1, R0                            ;Se compara el input con el 1!
BNE DOS_FACT                          ;Si no es igual se salta a corroborar con 2!
MOV R1, #1                            ;Si es igual guarda el factorial en R1
B FINAL_SIFACT                        ;Se va al final para printear el factorial 

DOS_FACT:
	MOV R1 , #2                   ;Se copia a R1 el segundo valor de los factoriales permitidos (2)
	CMP R1, R0                    ;Se compara el input con el 2!
	BNE TRES_FACT                 ;Si no es igual se salta a corroborar con 3!
	MOV R1, #2		      ;Si es igual guarda el factorial en R1
	B FINAL_SIFACT		      ;Se va al final para printear el factorial 

TRES_FACT:
	MOV R1 , #6		      ;Se copia a R1 el tercer valor de los factoriales permitidos (6)
	CMP R1, R0 		      ;Se compara el input con el 3!		
	BNE CUATRO_FACT 	      ;Si no es igual se salta a corroborar con 4!
	MOV R1, #3		      ;Si es igual guarda el factorial en R1
	B FINAL_SIFACT		      ;Se va al final para printear el factorial 

CUATRO_FACT:
	MOV R1 , #24		      ;Se copia a R1 el cuarto valor de los factoriales permitidos (24)
	CMP R1, R0 		      ;Se compara el input con el 4!
	BNE CINCO_FACT 		      ;Si no es igual se salta a corroborar con 5!
	MOV R1, #4		      ;Si es igual guarda el factorial en R1
	B FINAL_SIFACT		      ;Se va al final para printear el factorial 

CINCO_FACT:
	MOV R1 , #120 		      ;Se copia a R1 el quinto valor de los factoriales permitidos (120)
	CMP R1, R0 		      ;Se compara el input con el 5!
	BNE NO_FACT 		      ;Si no es igual se salta a printear 0, ya que el numero ingresado no es ningun factorial
	MOV R1, #5		      ;Si es igual guarda el factorial en R1
	B FINAL_SIFACT		      ;Se va al final para printear el factorial 

FINAL_SIFACT:  
	MOV R0, #1
    		      
	SWI 0x6b		      ;Printea el numero guaradado en R1 asociado al valor correspondiente del factorial
	LDR R0, =InFileHandle
    
	LDR R0, [R0]
    
	SWI 0x68
	SWI 0x11		      ;Cierra el archivo

NO_FACT:
	MOV R1,#-1  		      ;Asigna 0 a R1 para printearlo como indicador de que el numero no es factorial de otro
	MOV R0, #1
    
	SWI 0x6b		      ;Se printea 0 como resultado 
	LDR R0, =InFileHandle
    
	LDR R0, [R0]
    
	SWI 0x68
	SWI 0x11		      ;Se cierra el archivo
