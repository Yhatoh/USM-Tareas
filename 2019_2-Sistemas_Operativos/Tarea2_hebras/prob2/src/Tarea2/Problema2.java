package tarea2;
import java.util.concurrent.TimeUnit;
import java.util.Scanner;

public class Problema2 	
{

	public static void main(String[] args) {
	
	Scanner input = new Scanner(System.in);
	int a = input.nextInt();
	while(a > 0){
			
			int n = input.nextInt();
			int k = 0;
			int [] arr;

			arr = new int[n];
			while (k < n)
			{
				arr[k] = input.nextInt();
				k++;	
			}

			printa pri = new printa();
			quick qs = new quick(arr, 0, n-1);
			qs.start();
			try{
				qs.join();
			}
			catch(InterruptedException ie)
			{
				Thread.currentThread().interrupt();
			}

			pri.printas(arr);
			a--;

		}
	}
}


