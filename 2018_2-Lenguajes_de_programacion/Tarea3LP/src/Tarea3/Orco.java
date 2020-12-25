package tarea3;

public class Orco extends Raza{
	private int str;
	private int con;
	private int dex;

//setter y getter str
	public void setStr(){str = 2;}

	public int getStr(){return str;}

//setter y getter con
	public void setCon(){con = 1;}

	public int getCon(){return con;}

//setter y getter dex	
	public void setDex(){dex = 0;}
	
	public int getDex(){return dex;}

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
		System.out.println("Un orco ataca, gana +2 a la fuerza");
		return n + 2;
	}
}