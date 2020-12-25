package tarea2;

public class Nodo {
	public Arbol arbol_izq;
	public Arbol arbol_der;
	private String dato;

	public Nodo() {
		this.arbol_der = null;
		this.arbol_izq = null;
		this.dato = " ";
	}

	public String getDato() {
		return this.dato;
	}

	public void setDato(String dato) {
		this.dato = dato;
	}
}