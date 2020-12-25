package tarea2;

public class Arbol {
	public Nodo raiz;

	public Arbol() {
		this.raiz = new Nodo();
	}

	//Me daba laba printear el arbol a mano entonces hice esta funcion print
	public void print() {
		if(this.raiz.arbol_izq != null){this.raiz.arbol_izq.print();}
		if(this.raiz.arbol_der != null){this.raiz.arbol_der.print();}
		System.out.println("---------");
		System.out.println(this.raiz.getDato());
	}
	
}