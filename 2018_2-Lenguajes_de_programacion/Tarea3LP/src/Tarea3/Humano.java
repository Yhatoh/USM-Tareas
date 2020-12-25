package tarea3;

public class Humano extends Raza{
	private int str;
	private int dex;
	private int con;

//setter y getter str
	public void setStr(){str = 1;}

	public int getStr(){return str;}

//setter y getter dex
	public void setDex(){dex = 1;}

	public int getDex(){return dex;}

//setter y getter con
	public int getCon(){return con;}

	public void setCon(){con = 1;}

/* 
* void crearRaza()
* Inicializacion de la raza
*/

	public void crearRaza(){
		this.setStr();
		this.setDex();
		this.setCon();
	}
/*
 * int habilidad(int n, int dado)
 * Recibe un entero n y un entero dado
 * Revisa que el dato n no sea 1
 * Retorna un entero distinto de 1
 */
		
	public int habilidad(int n, int dado){
		while(n == 1){
			System.out.println("El dado lanzado fue 1, se vuelve a lanzar");
			n = Juego.lanzarDados(dado);		
			System.out.println("El nuevo valor obtenido es " + n);
		}
		return n;
	}
}