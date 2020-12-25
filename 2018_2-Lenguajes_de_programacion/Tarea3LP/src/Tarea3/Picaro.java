package tarea3;

public class Picaro extends Clase{
	private int clase;
	private int tipoAtaque; // 1 fuerza, 2 destreza, 3 constitucion
	private int modopelea; //0 ataque 1 defensa
	private int armadura;
	
//getter y setter clase
	public int getClase(){return clase;}

	public void setClase(int clase){this.clase = clase;}

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
		setClase(2);
		setTipoAtaque(2);
		setArmadura(10);
		setModopelea(0);
	}

/*
 * void ataque(Enemigo enem, Jugador player, int n).
 * Recibe el objeto Enemigo, el objeto Jugador y el entero n.
 * Si n es 0, el Jugador esta atacando al Enemigo, si n es 1 ocurre lo opuesto.
 * Realiza la accion de atacar de un picaro, usando dados de 8 y de 20, y utilizando la destreza del atacante.
 * Al final, cambia el valor de la vida del atacado, en caso que el ataque logre ser efectuado.
 */
	
	public void ataque(Enemigo enem, Jugador player, int n) {
		int ocho;
		int veinte=0;
		int damage;

		if(n == 0) {	
			if(enem.getModopelea() == 1){
				System.out.println( enem.getName() +" esta en modo defensa " + player.getName()+ " lanza el dado 2 veces");
				int try1 = Juego.lanzarDados(3);
				int try2 = Juego.lanzarDados(3);
				if (player.getRaza() == "Humano"){
					try1 = player.cartaTrampa(try1, 3);
					try2 = player.cartaTrampa(try2, 3);
				}
				System.out.println("Lanzamiento 1: "+try1);
				System.out.println("Lanzamiento 2: "+try2);
				if(try1 < try2){
					veinte = try1;
					System.out.println(try1 + " es menor");
				}
				else{
					veinte = try2;
					System.out.println(try2 + " es menor");
				}
			}

			else{ 
				veinte = Juego.lanzarDados(3);
				if(player.getRaza()== "Humano"){
					veinte = player.cartaTrampa(veinte,3);
				}
			}
			System.out.println("El valor del lanzamiento de " + player.getName() +" del dado de 20 para atacar es: "+ veinte);
			if (veinte >= enem.getArmor()){
				System.out.println(veinte+" es mayor que la armadura de " +enem.getName()+", el ataque afecta");
				ocho = Juego.lanzarDados(2);
				if(player.getRaza() == "Humano"){
					ocho = player.cartaTrampa(ocho,2);
				}
				System.out.println("El valor del lanzamiento de " + player.getName() +" del dado d8 para calcular el danio del ataque es: "+ ocho);
				damage = ocho + player.getDex();
				
				if(veinte == 20) {
					System.out.println("WOAH el lanzamiento d20 fue 20, el danio se duplica");
					damage = damage*2;
				}
				enem.asignarVida(enem.getVida()-damage);	
				System.out.println(enem.getName() +" recibe "+ damage + " de danio");
			}
			else {
				System.out.println(veinte+" es menor que la armadura de " +enem.getName()+", el ataque no afecta");
			}
		}	
		else {
			if(player.getModopelea() == 1){
				System.out.println( player.getName() +" esta en modo defensa " + enem.getName()+ " lanza el dado 2 veces");
				int try1 = Juego.lanzarDados(3);
				int try2 = Juego.lanzarDados(3);
				if (enem.getRaza() == "Humano"){
					try1 = enem.cartaTrampa(try1, 3);
					try2 = enem.cartaTrampa(try2, 3);
				}
				System.out.println("Lanzamiento 1: "+try1);
				System.out.println("Lanzamiento 2: "+try2);
				if(try1 < try2){
					veinte = try1;
					System.out.println(try1 + " es menor");
				}
				else{
					veinte = try2;
					System.out.println(try2 + " es menor");
				}
			}
			System.out.println("El valor del lanzamiento de " + enem.getName() +" del dado de 20 para atacar es: "+ veinte);
			if (veinte >= player.getArmor()) {
				System.out.println(veinte+" es mayor que la armadura de " +player.getName()+", el ataque afecta");
				ocho = Juego.lanzarDados(2);
				
				if(enem.getRaza() == "Humano"){
					ocho = enem.cartaTrampa(ocho,2);
				}
				System.out.println("El valor del lanzamiento de " + enem.getName() +" del dado d8 para calcular el danio del ataque es:"+ ocho);
				damage = ocho + enem.getDex();
				System.out.println( enem.getName() +" ataca con "+ damage + " de danio" );
				if(veinte == 20) {
					System.out.println("WOAH el lanzamiento d20 fue 20, el danio se duplica");
					damage = damage*2;
				}
				player.asignarVida(player.getVida()-damage);
				System.out.println(player.getName() + " recibe " + damage + "de danio");
			}
			else {
				System.out.println(veinte+" es menor que la armadura de " +player.getName()+", el ataque no afecta");
			}
		}
	}
}