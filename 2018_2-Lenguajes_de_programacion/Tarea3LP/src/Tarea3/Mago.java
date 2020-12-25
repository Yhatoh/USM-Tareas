package tarea3;


public class Mago extends Clase{	
	private int clase;
	private int tipoAtaque; // 1 fuerza, 2 destreza, 3 constitucion
	private int modopelea; //0 ataque 1 defensa
	private int armadura;
	
//getter y setter clase
	public int getClase(){return clase;}

	public void setClase(int clase) {this.clase = clase;}

//getter y setter tipoataque
	public int getTipoAtaque(){return tipoAtaque;}

	public void setTipoAtaque(int tipoAtaque){this.tipoAtaque = tipoAtaque;}

//getter y setter modopelea
	public int getModopelea(){return modopelea;}

	public void setModopelea(int modopelea){this.modopelea = modopelea;}

//getter y setter armadura
	public int getArmadura(){return armadura;}

	public void setArmadura(int armadura){this.armadura = armadura;}


/* 
* void defender()
* Cambio de modo a defensa 
*/
	public void defender() {
		setModopelea(1);
	}
/* 
 * void crearClase()
 * Inicializacion de la clase
 */
	public void crearClase() {
		setClase(3);
		setTipoAtaque(2);
		setArmadura(10);
		setModopelea(0);
	}
	

/*
 * void ataque(Enemigo enem, Jugador player, int n).
 * Recibe el objeto Enemigo, el objeto Jugador y el entero n.
 * Si n es 0, el Jugador esta atacando al Enemigo, si n es 1 ocurre lo opuesto.
 * Realiza la accion de atacar de un mago, usando dados de 6 para atacar.
 * El atacado utiliza el dado de 20 para intentar esquivar o resistir el ataque
 * Al final, cambia el valor de la vida del atacado, en caso que el ataque logre ser efectuado.
 */
	
	public void ataque(Enemigo enem, Jugador player, int n) {
		int seis;
		int veinte=0;
		int damage;
		
		if(n == 0) {
			if(enem.getModopelea() == 1){
				System.out.println( enem.getName() +" esta en modo defensa " + enem.getName()+ " lanza el dado 2 veces");
				int try1 = Juego.lanzarDados(3);
				int try2 = Juego.lanzarDados(3);
				if (enem.getRaza() == "Humano"){
					try1 = enem.cartaTrampa(try1, 3);
					try2 = enem.cartaTrampa(try2, 3);
				}
				System.out.println("Lanzamiento 1: "+try1);
				System.out.println("Lanzamiento 2: "+try2);
				if(try1 > try2){
					veinte = try1;
					System.out.println(try1+ " es mayor");
				}
				else{
					System.out.println(try2+ " es mayor");
					veinte = try2;
				}
			}
			else{ 
				veinte = Juego.lanzarDados(3);
				
				if(enem.getRaza()== "Humano"){
					veinte = enem.cartaTrampa(veinte,3);
				}
			}
			
			System.out.println("El valor del lanzamiento de " + enem.getName() +" del dado de 20 para esquivar es: "+ veinte);
			
			if(enem.getRaza() == "Elfo"){
				veinte = enem.cartaTrampa(veinte, 1);
			}

			if (veinte+enem.getDex() < 13) {
				System.out.println(enem.getName() + " No logro esquivar el ataque");
				seis = Juego.lanzarDados(1);

				if(player.getRaza()== "Humano"){
					seis = player.cartaTrampa(seis, 1);
				}
				System.out.println("El valor del lanzamiento de " + player.getName() + " del d6 es: "+seis);
				
				damage = seis;
				enem.asignarVida(enem.getVida()-damage);
				System.out.println(enem.getName() +" recibe "+ damage + " de danio");
			}
			else {
				if(veinte == 20) {
					System.out.println(enem.getName() + " esquivo el ataque");
				}
				else {
					System.out.println(enem.getName() + " logro resistir el ataque");
					seis = Juego.lanzarDados(1);
					if(player.getRaza()== "Humano"){
						seis = player.cartaTrampa(seis, 1);
					}
					System.out.println("El valor del lanzamiento de " + player.getName() + " del d6 es: "+seis);
					damage = (seis)/2;
					enem.asignarVida(enem.getVida()-damage);
					System.out.println(enem.getName() +" recibe "+ damage + " de danio");
				}
			}
		}
		else {
			if(player.getModopelea() == 1){
				System.out.println( player.getName() +" esta en modo defensa, "+ player.getName()+ " lanza el dado 2 veces");
				int try1 = Juego.lanzarDados(3);
				int try2 = Juego.lanzarDados(3);
				if (player.getRaza() == "Humano"){
					try1 = player.cartaTrampa(try1, 3);
					try2 = player.cartaTrampa(try2, 3);
				}
				System.out.println("Lanzamiento 1: "+try1);
				System.out.println("Lanzamiento 2: "+try2);
				if(try1 > try2){
					veinte = try1;
					System.out.println(try1+ " es mayor");
				}
				else{
					System.out.println(try2+ " es mayor");
					veinte = try2;
				}
			}

			else{ 
				veinte = Juego.lanzarDados(3);
				if(player.getRaza()== "Humano"){
					veinte = player.cartaTrampa(veinte,3);
				}
			}
			System.out.println("El valor del lanzamiento de " + player.getName() +" del dado de 20 para esquivar es: "+ veinte);
			if(player.getRaza() == "Elfo"){
				veinte = player.cartaTrampa(veinte, 1);
			}

			if (veinte + player.getCon() < 13) {
				
				System.out.println(player.getName() + " No logro esquivar el ataque");
				seis = Juego.lanzarDados(1);

				if(enem.getRaza() == "Humano"){
					seis = enem.cartaTrampa(seis, 1);
				}
				System.out.println("El valor del lanzamiento de " + enem.getName() + " del d6 es: "+seis);
				damage = seis;
				player.asignarVida(player.getVida()-damage);
				System.out.println(player.getName() +" recibe "+ damage + " de danio");
			}
			else {
				if(veinte == 20) {
					System.out.println(player.getName() + " esquivo el ataque");
				}
				else {
					System.out.println(player.getName() + " logro resistir el ataque");
					seis = Juego.lanzarDados(1);

					if(enem.getRaza()== "Humano"){
						seis = enem.cartaTrampa(seis, 1);
					}
					
					System.out.println("El valor del lanzamiento de " + enem.getName() + " del d6 es: "+seis);
					
					damage = seis/2;
					player.asignarVida(player.getVida()-damage);
					System.out.println(player.getName() +" recibe "+ damage + " de danio");
				}
			}
		}
	}
}
