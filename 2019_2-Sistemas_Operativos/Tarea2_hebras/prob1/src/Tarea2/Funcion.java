package tarea2;

import java.util.LinkedList;

public class Funcion extends Thread {

	public Arbol arbolito;
	public int x;

	public Funcion(String funcion) {
		this.arbolito = new Arbol();
		funcion = reemplazando_f(funcion);
		this.arbolito = this.ayuda_esta_funcion_la_hice_a_las_4am_y_entro_a_trabajar_a_las_9_30(funcion);
	}

	//--- Conjuntos de funciones para parsear la funcion entregada


	/*
		Nombre  : ayuda_esta_funcion_la_hice_a_las_4am_y_entro_a_trabajar_a_las_9_30 (AEFLHAL4YEATAL930)

		Objetivo: Esta funcion es la gran funcion recursiva que permite tomar el string que
				  representa a la funcion y devolverlo en un arbol de precendecia, para asi
				  tener un manejo completo de que operacion se realiza primero y cual no.

				  Para esto primero se pregunta si hay paratensis buscando la parentesis de mas adentro
				  si es que encuentra se llama recursivamente si no sigue adelante con su vida.

				  Luego verifica que el string funcion que recibe tiene * o /, si es que tiene
				  toma lo que hay a la izquierda y a la derecha de la operacion reemplazando aquello por un # y 
				  guardando eso en un arbol que sera puesto en un arreglo dependiendo de cual #_i es de izquierda
				  a derecha. Hace esto en un while hasta que no haya mas * o /.

				  Aplica la misma logica para + o -

				  Ahora si durante cualquiera de esos dos loops anteriores se encuentra con un # o dos #,
				  los va sacando de la lista de # y los va uniendo, para ir creando el arbol correctamente.

				  por ejemplo:
				  	2*4+3*2+1*(2-2)
				  	1 Iteracion (aqui creo un arbol de 2 - 2 = #_1)
				  		Listas de # [#_1]
						2*4+3*2+1*#_1

					2 Iteracion (aqui creo otro arbol de 2 * 4 = #_2)
						Listas de # [#_2,#_1]
						#_2+3*2+1*#_1

					3 Iteracion (aqui creo otro arbol de 3 * 2 = #_3)
						Listas de # [#_2,#_3,#_1]
						#_2+#_3+1*#_1

					4 Iteracion (aqui creo otro arbol 1 * #_1 = #_4)
						Listas de # [#_2,#_3,#_4]
						#_2+#_3+#_4

					5 Iteracion (aqui crea otro arbol #_2 + #_3 = #_5)
						Listas de # [#_5,#_4]
						#_5+#_4

					6 Iteracion (crea el arbol final uniendo #_5 + #_4 = #_6)
						Listas de # [#_6]
						#_6

					Donde #_6 sería
					          +
					      /      \
					     +        *
					   /   \     / \
                      *     *   1   -
                     / \   / \     / \
                    2   4 3   2   2   2
		Retorna : Arbol que corresponde al arbol de precendecia correspondiente
	*/
	public Arbol ayuda_esta_funcion_la_hice_a_las_4am_y_entro_a_trabajar_a_las_9_30(String funcion) {
		Arbol returne = new Arbol();
		String sub;
		int pos, pos_izq, pos_der;
		LinkedList<Arbol> bosque = new LinkedList<Arbol>();
		int[] arr = this.masAdentro(funcion);

		while(arr[0] != -1 && arr[1] != -1) {
			sub = funcion.substring(arr[0]+1,arr[1]);

			bosque.add(ayuda_esta_funcion_la_hice_a_las_4am_y_entro_a_trabajar_a_las_9_30(sub));
			funcion = funcion.substring(0,arr[0])+"#"+funcion.substring(arr[1]+1,funcion.length());
			arr = this.masAdentro(funcion);
		}

		pos = this.hayMultODiv(funcion);
		while(pos != -1) {
			Nodo nodo = new Nodo();
			nodo.setDato(""+funcion.charAt(pos));
			if(funcion.charAt(pos-1) == '#' && funcion.charAt(pos+1) == '#') {
				pos_izq = pos-1; pos_der = pos+2;
				nodo.arbol_izq = bosque.get(quienSoy(funcion, pos-1));
				nodo.arbol_der = bosque.get(quienSoy(funcion, pos));
				bosque.remove(quienSoy(funcion, pos-1));
				bosque.remove(quienSoy(funcion, pos+1)-1);
			}
			else if(funcion.charAt(pos-1) != '#' && funcion.charAt(pos+1) == '#') {
				pos_der = pos+2; pos_izq = cuantosCharIzq(funcion,pos)+1;
				nodo.arbol_izq = new Arbol();
				nodo.arbol_izq.raiz.setDato(funcion.substring(pos_izq,pos));
				nodo.arbol_der = bosque.get(quienSoy(funcion, pos));
				bosque.remove(quienSoy(funcion, pos+1));
			}
			else if(funcion.charAt(pos-1) == '#' && funcion.charAt(pos+1) != '#') {
				pos_izq = pos-1; pos_der = cuantosCharDer(funcion,pos);
				nodo.arbol_der = new Arbol();
				nodo.arbol_der.raiz.setDato(funcion.substring(pos+1,pos_der));
				nodo.arbol_izq = bosque.get(quienSoy(funcion, pos-1));
				bosque.remove(quienSoy(funcion, pos-1));
			}
			else {
				pos_izq = cuantosCharIzq(funcion,pos)+1; pos_der = cuantosCharDer(funcion,pos);
				nodo.arbol_izq = new Arbol();
				nodo.arbol_izq.raiz.setDato(funcion.substring(pos_izq,pos));
				nodo.arbol_der = new Arbol();
				nodo.arbol_der.raiz.setDato(funcion.substring(pos+1,pos_der));
			}
			Arbol aux = new Arbol();
			aux.raiz = nodo;

			funcion = funcion.substring(0,pos_izq)+"#"+funcion.substring(pos_der,funcion.length());

			if(bosque.size() == 0){bosque.add(aux);}
			else if(quienSoy(funcion,pos_izq) == bosque.size()){bosque.add(aux);}
			else{bosque.add(quienSoy(funcion,pos_izq),aux);}
			pos = this.hayMultODiv(funcion);

		}

		pos = this.hayAddOSus(funcion);
		while(pos != -1) {
			Nodo nodo = new Nodo();
			nodo.setDato(""+funcion.charAt(pos));
			if(funcion.charAt(pos-1) == '#' && funcion.charAt(pos+1) == '#') {
				pos_izq = pos-1; pos_der = pos+2;
				nodo.arbol_izq = bosque.get(quienSoy(funcion, pos-1));
				nodo.arbol_der = bosque.get(quienSoy(funcion, pos));
				bosque.remove(quienSoy(funcion, pos-1));
				bosque.remove(quienSoy(funcion, pos+1)-1);
			}
			else if(funcion.charAt(pos-1) != '#' && funcion.charAt(pos+1) == '#') {
				pos_der = pos+2; pos_izq = cuantosCharIzq(funcion,pos)+1;
				nodo.arbol_izq = new Arbol();
				nodo.arbol_izq.raiz.setDato(funcion.substring(pos_izq,pos));
				nodo.arbol_der = bosque.get(quienSoy(funcion, pos));
				bosque.remove(quienSoy(funcion, pos+1));
			}
			else if(funcion.charAt(pos-1) == '#' && funcion.charAt(pos+1) != '#') {
				pos_izq = pos-1; pos_der = cuantosCharDer(funcion,pos);
				nodo.arbol_der = new Arbol();
				nodo.arbol_der.raiz.setDato(funcion.substring(pos+1,pos_der));
				nodo.arbol_izq = bosque.get(quienSoy(funcion, pos-1));
				bosque.remove(quienSoy(funcion, pos-1));
			}
			else {
				pos_izq = cuantosCharIzq(funcion,pos)+1; pos_der = cuantosCharDer(funcion,pos);

				nodo.arbol_izq = new Arbol();
				nodo.arbol_izq.raiz.setDato(funcion.substring(pos_izq,pos));
				nodo.arbol_der = new Arbol();
				nodo.arbol_der.raiz.setDato(funcion.substring(pos+1,pos_der));
			}
			Arbol aux = new Arbol();
			aux.raiz = nodo;
			funcion = funcion.substring(0,pos_izq)+"#"+funcion.substring(pos_der,funcion.length());

			if(bosque.size() == 0){bosque.add(aux);}
			else if(quienSoy(funcion,pos_izq) == bosque.size()){bosque.add(aux);}
			else{bosque.add(quienSoy(funcion,pos_izq),aux);}
			pos = this.hayAddOSus(funcion);
		}
		returne = bosque.get(0);
		return returne;
	}

	/*
		Nombre  : reemplazando_f

		Objetivo: recibe el string correspondiente a la funcion a parsear, y lo que hace es
		          reemplazar todos los string(x) por string, simplemente porque me molestaban
				  el (x) eso basicamente,
	
		Retorna : el string a parsear.
	*/
	public String reemplazando_f(String funcion) {
		int i = 0;
		int inicio, termino;
		while(true){
			if(i < funcion.length() && funcion.charAt(i) != '+' && funcion.charAt(i) != 'x' && funcion.charAt(i) != '-' && funcion.charAt(i) != '*' && funcion.charAt(i) != '/' && funcion.charAt(i) != '1' && funcion.charAt(i) != '2' && funcion.charAt(i) != '3' && funcion.charAt(i) != '4' && funcion.charAt(i) != '5' && funcion.charAt(i) != '6' && funcion.charAt(i) != '7' && funcion.charAt(i) != '8' && funcion.charAt(i) != '9' && funcion.charAt(i) != '0' && funcion.charAt(i) != '(' && funcion.charAt(i) != ')'){
				inicio = i;
				while(i <= funcion.length()-1 && funcion.charAt(i) != '+' && funcion.charAt(i) != '-' && funcion.charAt(i) != '*' && funcion.charAt(i) != '/' && funcion.charAt(i) != '1' && funcion.charAt(i) != '2' && funcion.charAt(i) != '3' && funcion.charAt(i) != '4' && funcion.charAt(i) != '6' && funcion.charAt(i) != '7' && funcion.charAt(i) != '8' && funcion.charAt(i) != '9' && funcion.charAt(i) != '0'){
					i++;
				}
				termino = i;
				funcion = funcion.substring(0,inicio)+funcion.substring(inicio,termino-3)+funcion.substring(termino,funcion.length());
				i = termino-3;
				if(i == funcion.length()) break;
			}
			i++;
			if(i == funcion.length()) break;

			
		}
		return funcion;
	}

	/*
		Nombre  : cuantosCharIzq

		Objetivo: recibe el string a parsear y la posicion de una operacion +, -, * o /
		          lo que hace es saber cuantos caracteres hay hasta la siguiente operacion o parentesis,
		          puesto que dentro de la funcion AEFLHAL4YEATAL930 voy a reemplazar y tomar estos caracteres.
	
		Retorna : la cantidad de caracteres que hay a la izquierda de la operacion encontrada
	*/
	public int cuantosCharIzq(String funcion, int pos){
		int i;
		for(i = pos-1; i >= 0 && funcion.charAt(i) != '+' && funcion.charAt(i) != '-' && funcion.charAt(i) != '*' && funcion.charAt(i) != '/'; i--){}
		return i;
	}

	/*
		Nombre  : cuantosCharDer

		Objetivo: hace lo mismo que la funcion cuantosCharIzq, pero hacia la derecha de la operacion
	
		Retorna : la cantidad de caracteres que hay a la derecha de la operacion encontrada
	*/
	public int cuantosCharDer(String funcion, int pos){
		int i;
		for(i = pos+1; i < funcion.length() && funcion.charAt(i) != '+' && funcion.charAt(i) != '-' && funcion.charAt(i) != '*' && funcion.charAt(i) != '/'; i++){}
		return i;
	}

	/*
		Nombre  : quienSoy

		Objetivo: eventualmente un string funcion pasa a ser #+#+#+#, o cosas parecidas, esta funcion
		          lo que hace es saber en que posicion de una lista que almacena las funciones correspondientes
		          al #, esta un # que yo quiero.
		          Recorre de 0 hasta la posicion del # que quiero saber.
	
		Retorna : la posicion en la lista del # que quiero preguntar
	*/
	public int quienSoy(String funcion, int pos) {
		int i, cont;
		cont = 0;
		for(i = 0; i < pos; i++) {
			if(funcion.charAt(i) == '#') cont++;
		}
		return cont;
	}

	/*
		Nombre  : masAdentro

		Objetivo: teniendo el string a parsear de una funcion, esta funcion determinan el (cosasaqui) mas adentro 
				  del string.
	
		Retorna : las coordenadas x e y, del ( y ) más adentro.
	*/
	public int[] masAdentro(String funcion) {
		int[] xy = {-1,-1};
		int i;

		for(i = 0; i < funcion.length(); i++) {
			if(funcion.charAt(i) == '('){
				xy[0] = i; 
			}
			else if(funcion.charAt(i) == ')'){
				xy[1] = i;
				break;
			}
		}
		return xy;
	}

	/*
		Nombre  : hayMultODiv

		Objetivo: recibe un string y determina la primera posicion que haya un * o /
	
		Retorna : la posicion del primer * o / si no hay retorna -1
	*/
	public int hayMultODiv(String funcion) {
		int i;

		for(i = 0; i < funcion.length(); i++ ) {
			if(funcion.charAt(i) == '*' || funcion.charAt(i) == '/') {return i;}
		}
		return -1;
	}

	/*
		Nombre  : hayAddOSus

		Objetivo: recibe un string y determina la primera posicion que haya un - o +
	
		Retorna : la posicion del primer - o + si no hay retorna -1
	*/
	public int hayAddOSus(String funcion) {
		int i;

		for(i = 0; i < funcion.length(); i++ ) {
			if(funcion.charAt(i) == '+' || funcion.charAt(i) == '-') {return i;}
		}
		return -1;
	}

	//-- Funciones para imprimir porque me daba lata escribir System.out.println() --
	
	public void pline(String line) {
		System.out.println(line);
	}

	public void parray(String[] array) {
		for(int i = 0; i < array.length; i++) {
			this.pline(array[i]);
		}
	}
}