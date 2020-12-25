Para compilar, en una terminal ubicada en la carpeta /prob2 correr el comando:

make run
 

Para correr el codigo utilizar 

java -jar Problema2.jar


Para la entrada se necesita un numero n, el cual equivale a la cantidad de arreglos que se quiere
ordenar, seguido de n lineas.
Para esto, puede usar el siguiente formato:

2
4 3 6 3 2
3 2 1 7

En caso de querer utilizar un archivo .txt correr el codigo con el siguiente comando:

java -jar Problema2.jar < prueba.txt



Para este problema se utilizo el algoritmo de QuickSort, al cual se le añadieron threats, las cuales se generan
cada vez que el algorimto genera una particion en el proceso de dividir. Esto beneficia a la eficiencia del algoritmo
ya que cada particion puede trabajar apenas es llamada, en vez de tener que esperar que su "hermana" (la que se genera
con la otra mitad) termine.
