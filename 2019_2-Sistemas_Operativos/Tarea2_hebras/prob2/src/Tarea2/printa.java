package tarea2;



public class printa
{
/*
	Funcion Printa(int [] arr)

	Input: Arreglo de enteros
	Descripcion: Muestra por pantalla un arreglo, entre [] y separando los datos por comas.
	No tiene output
*/
	public void printas(int [] arr)
	{
		System.out.print("[");
		System.out.print(arr[0]);
		for (int i = 1; i < arr.length ; i++ ) 
		{
			System.out.print(",");
			System.out.print(arr[i]);
		}
		System.out.print("]\n");

	}
}