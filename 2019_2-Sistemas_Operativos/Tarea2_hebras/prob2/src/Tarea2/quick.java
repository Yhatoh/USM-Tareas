package tarea2;

public class quick extends Thread
{

	int [] arr;
	int lower;
	int upper;
	
/*
	Funcion quck(int [] arr, int lower, int upper)

	Input: Arreglo de enteros, lower: inicio del sub arreglo, upper: final del subarreglo
	Descripcion: Constructor de la clase quick
	No tiene output
*/
	public quick(int [] arr, int lower, int upper)
	{
		this.arr = arr;
		this.lower = lower;
		this.upper = upper;
	}

/*
	Funcion run()

	No tiene input
	Descripcion: Inicia la threat, toma el dato que se encuentra al centro del arreglo y lo utiliza como pivote
	dejando asi, todos los datos mayores que el pivote a la derecha, y los menores a la izquierda. Una vez realiza esto
	crea dos threats los cual vuelven a llamar a esta funcion, pero con el sub arreglo izquierdo y con el sub arreglo derecho.
	Esto se realiza hasta que se encuentre un arreglo de largo 1.
	No tiene output
*/
	public void run()
	{

		int c = lower;
		int k = upper;
		int pivote = arr[(lower+upper)/2];
		while(c<k)
		{

			while(arr[k] < pivote) k--;
			while(arr[c] > pivote) c++;
		
			if (c <= k)
			{
				int temp = arr[c];
				arr[c] = arr[k];
				arr[k]= temp;
				k--;
				c++;
			}
			quick a = new quick(arr, lower, k);
			if (lower < k) 
			{
				a.start();		
			}
			quick b = new quick(arr, c, upper);		
			if (c < upper) 
			{
				b.start();
			}	


			try{
				a.join();
			}
			catch(InterruptedException ie)
			{
				Thread.currentThread().interrupt();
			}

			try{
				b.join();
			}
			catch(InterruptedException ie)
			{
				Thread.currentThread().interrupt();
			}
		}
	}
}