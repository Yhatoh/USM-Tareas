package tarea3;
import java.util.List;
import java.util.Random;
import java.util.Scanner;
import java.util.ArrayList;
public class Juego {

	/*
	int lanzarDados(int a)

	Recibe un numero entre 1, 2 y 3 y dependiendo de esos valores lanza un d6, d8, d10 respectivamente.
	Retornando un valor int que es el resultado del lanzamiento.
	*/
	public static int lanzarDados(int a){
		//1 de dado de 6, 2 dado de 8, 3 dado de 20
		Random rand = new Random();
		int num;
		if (a == 1){
			num = rand.nextInt(6) + 1;
			return num;
		}
		else if (a == 2){
			num = rand.nextInt(8) + 1;
			return num;
		}
		else if (a == 3){
			num = rand.nextInt(20) + 1;
			return num; 
		}
		else{
			return 0;
		}
	}

	public static void main(String[] args) {
		List<String> enemigos = new ArrayList<String>();
		Random randw = new Random();
		enemigos.add("Klrak, el enano barbaro");
		enemigos.add("Adran, el elfo picaro");
		enemigos.add("Isaac, el humano clerigo");
		enemigos.add("Elysium, el elfo mago");
		enemigos.add("Krrogh, el orco barbaro");
		enemigos.add("Jenkins, el humano mago");
		
		int d,n,n2,i = 0,ii,newlife;
		List<Enemigo> enemigosP = new ArrayList<Enemigo>();
		while (i < 3){
			i++;
			d = lanzarDados(1);
			Enemigo enem = new Enemigo(enemigos.get(d-1));
			enemigosP.add(enem);
		}
		Scanner reader = new Scanner(System.in);

		System.out.println("Te advertimos que esta aventura sera tan epica que si tiendes a tener problemas del corazon por favor no juegues");
		System.out.println("Dinos tu nombre");
		String nombresito = reader.nextLine();
		System.out.println("Elige raza:\n1-Humano\n2-Elfo\n3-Enano\n4-Orco");
		n = reader.nextInt();
		System.out.println("Elige clase:\n1-Barbaro\n2-Picaro\n3-Mago\n4-Clerigo");
		n2 = reader.nextInt();

		Enemigo e;
		Jugador j1 = new Jugador(nombresito,n,n2);
		System.out.println("Ya wena " + j1.getName() + " eri un " + j1.getRaza() + " " + j1.getCla());
		System.out.println("Tus caracteristicas son: \nFuerza: " + j1.getStr() + "\nDestreza: " + j1.getDex() +"\nConstitucion: " + j1.getCon());
		System.out.println("-------COMIENZA TU SUPER AVENTURA SUPER PREPARATE-------");
		i = j1.getVida();
		while(j1.getVida() > 0 && !(enemigosP.isEmpty())){
			e = enemigosP.get(0);
			enemigosP.remove(0);
			ii = e.getVida();
			System.out.println(j1.getName() + " te enfrenteras contra " + e.getName() + "  (suerte la necesitaras)");
			while(j1.getVida() > 0 && e.getVida() > 0){
				System.out.println(j1.getName() + " HP: " + j1.getVida());
				System.out.println(e.getName() + " HP: " + e.getVida());
				System.out.println("-------Tu turno-------");

				if(j1.getRaza().equals("Enano") && i != j1.getVida()){
					newlife = j1.cartaTrampa(j1.getVida(),2);
					j1.asignarVida(newlife);
				}
				System.out.println("Que quieres hacer:\n1-atacar furiosamente \n2-defender duramente");
				n = reader.nextInt();
				j1.setModopelea(n-1);
				if(n == 1){
					j1.ataque(e,j1);
				}
				if(e.getVida() <= 0 && !(enemigosP.isEmpty())){
					System.out.println("Bien hecho debe de haber sido dificil matar al enemigo debe haber sido suerte\nVEREMOS QUE TAL TE VA CONTRA EL SIGUIENTE.");	
					break;
				}
				else if(e.getVida() <= 0 && enemigosP.isEmpty()) break;

				System.out.println(j1.getName() + " HP: " + j1.getVida());
				System.out.println(e.getName() + " HP: " + e.getVida());
				System.out.println("-------Turno del enemigo-------");
				if(e.getRaza().equals("Enano") && ii != e.getVida()){
					newlife = e.cartaTrampa(j1.getVida(),2);
					e.asignarVida(newlife);
				}
				n = randw.nextInt(999999999)%2;
				e.setModopelea(n);
				if(n == 0){
					System.out.println("El oponente decidio atacarte furiosamente");
					e.ataque(e,j1);
				}
				else System.out.println("Se puso a defender ya que con su sharingan intenta predecir tu ataque");
				
				if(j1.getVida() <= 0){
					System.out.println("Oye como que no estas respirando.");	
				}
			}
		}
		if(j1.getVida() <= 0){
			System.out.println("RIP " + j1.getName() + " 2018-2018");
		}
		else if(enemigosP.isEmpty()){
			System.out.println("Felicidades " + j1.getName() + " eres el mejor jugador que ha tenido este juego (kappa).");
		}
		reader.close();
	}
}


