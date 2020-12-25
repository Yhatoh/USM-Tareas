package tarea3;

public class Elfo extends Raza{
	private int dex;
	private int con;
	private int str;

//setter y getter dex
	public int getDex(){return dex;}

	public void setDex(){dex = 2;}

//setter y getter con
	public int getCon(){return con;}

	public void setCon(){con = 1;}

//setter y getter str
	public int getStr(){return str;}
	
	public void setStr(){str = 0;}

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
* Aumenta en 2 el lanzamiento del dado de 20  para evadir
* Retorna un el entero del lanzamiento con el aumento realizado
*/
	public int habilidad(int n,int dado){
		System.out.println("Un elfo rolleo d20 para evadir");	
		System.out.println((n+2) +" es el nuevo valor de evacion o resistencia");
		return n + 2;
	}
}