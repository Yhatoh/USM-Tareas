package tarea3;

public class Enano extends Raza{
	private int str;
	private int con;
	private int dex;

//setter y getter str
	public int getStr(){return str;}

	public void setStr(){str = 1;}

//setter y getter dex
	public int getDex(){return dex;}

	public void setDex(){dex = 0;}

//setter y getter con
	public int getCon(){return con;}

	public void setCon(){con = 2;}

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
* Aumenta en 1 el dato n que representa la vida
* Retorna la nueva vida
*/
	public int habilidad(int n, int dado){
		System.out.println("Te has curado 1 de vida");
		return n + 1;
	}
}