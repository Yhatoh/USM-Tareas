package tarea3;

//1 fuerza, 2 destreza, 3 constitucion 
public abstract class Clase{	
	public abstract void crearClase();
	public abstract void ataque(Enemigo enem, Jugador player, int n);
	public abstract void defender();
	public abstract void setModopelea(int modopelea);
	public abstract int getClase();
	public abstract int getTipoAtaque();
	public abstract int getModopelea();
	public abstract int getArmadura();
}