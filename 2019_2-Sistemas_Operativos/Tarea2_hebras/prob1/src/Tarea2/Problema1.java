package tarea2;

import java.io.*;
import java.util.*;

public class Problema1 {

	public static void main(String[] args) {
		File file = new File("./funciones.txt");
		String[] line;
		try{
			Map<String,Funcion> mapita = new HashMap<String,Funcion>();
			Funcion funcion;
			Scanner input = new Scanner(file);
			int n = Integer.parseInt(input.nextLine());
			while(n > 0) {
				line = input.nextLine().split("=");
				funcion = new Funcion(line[1]);
				mapita.put((line[0].split("\\("))[0], funcion);
				n--;
			}
			System.out.println("Funciones ingresadas!");
			Scanner scanner = new Scanner(System.in);
			String x, funcionString, numberString;
			System.out.println("Ingrese operacion: ");
			x = scanner.next();

			while(!x.equals("STOP")){
				funcionString = x.split("\\(")[0];
				numberString = (x.split("\\(")[1]).split("\\)")[0];
				n = Integer.parseInt(numberString);
				funcion = mapita.get(funcionString);
				funcion.x = n;
				Hilo hilitoLindoPrecioso = new Hilo(funcion, mapita);
				try{
					hilitoLindoPrecioso.start();
					hilitoLindoPrecioso.join();
					System.out.println("El resultado es: " + hilitoLindoPrecioso.valorsillo);
				} catch (Exception e) {
					System.out.println("Ya que paso " + e);
				}
				System.out.println("Ingrese operacion: ");
				x = scanner.next();
			}
			
		} catch (FileNotFoundException e) {
			System.out.println(e);
		}
	}
}


