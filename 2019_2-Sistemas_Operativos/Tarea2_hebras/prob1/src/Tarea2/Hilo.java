package tarea2;

import java.io.*;
import java.util.*;

class Hilo extends Thread {

	public Arbol arbolito;
	public int x;
	public Map<String,Funcion> mapita;
	public int valorsillo;

	public Hilo(Funcion funcion, Map<String,Funcion> mapita) {
		this.arbolito = funcion.arbolito;
		this.x = funcion.x;
		this.mapita = mapita;

	}

	public void run() {
		try {
			this.valorsillo = this.calculando(this.x,this.arbolito);
		} catch (Exception e) {
			System.out.println("Me cai aqui " + e);
		}
	}


	/*
		Nombre  : calcular

		Objetivo: dado un arbol de precedencia y un valor x me calcula el valor de una funcion 
				  evaluado en x. 
				  El valor lo calula recorreciendio el arbol en post orden hasta que llegue a un numero o funcion y 
				  empieza a retonar y wiiiii resultado.
	
		Retorna : el resultado :D
	*/
	public int calculando(int x, Arbol ab) {
		if(ab.raiz.arbol_izq == null && ab.raiz.arbol_der == null && ab.raiz.getDato().charAt(0) != 'x' && this.sereUnNumero(ab.raiz.getDato())) return Integer.parseInt(ab.raiz.getDato());
		else if(ab.raiz.arbol_izq == null && ab.raiz.arbol_der == null && ab.raiz.getDato().charAt(0) == 'x' && ab.raiz.getDato().length() == 1) return x;
		else if(ab.raiz.arbol_izq == null && ab.raiz.arbol_der == null && ab.raiz.getDato().length() >= 1) {
			Funcion funcion = mapita.get(ab.raiz.getDato());
			funcion.x = x;

			Hilo hilitoLindoPrecioso = new Hilo(funcion,mapita);
			try{
				hilitoLindoPrecioso.start();
				hilitoLindoPrecioso.join();
				return hilitoLindoPrecioso.valorsillo;
			} catch (Exception e) {
				System.out.println("No aqui " + e);
				return -1;
			}
		}

		int valIzq, valDer;
		valIzq = valDer = 0;
		if(ab.raiz.arbol_izq != null) {
			valIzq = this.calculando(x, ab.raiz.arbol_izq);
		}
		if(ab.raiz.arbol_der != null) {
			valDer = this.calculando(x, ab.raiz.arbol_der);
		}

		if(ab.raiz.getDato().charAt(0) == '*') return valIzq * valDer;
		else if(ab.raiz.getDato().charAt(0) == '+') return valIzq + valDer;
		else if(ab.raiz.getDato().charAt(0) == '-') return valIzq - valDer;
		else return valIzq / valDer;
	}

	public boolean sereUnNumero(String palabra){
		int i;
		for(i = 0; i < palabra.length(); i++){
			if(palabra.charAt(i) != '0' && palabra.charAt(i) != '1' && palabra.charAt(i) != '2' && palabra.charAt(i) != '3' && palabra.charAt(i) != '4' && palabra.charAt(i) != '5' && palabra.charAt(i) != '6' && palabra.charAt(i) != '7' && palabra.charAt(i) != '8' && palabra.charAt(i) != '9'){
				return false;
			}
		}
		return true;
	}
}