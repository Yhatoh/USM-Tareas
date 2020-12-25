package tarea3;

public class Jugador implements Personaje {
	private String name,raza,cla;
	private int vida;
	private Raza race;
	private Clase clase;

	Jugador(String nombre,int n,int n2){
		this.asignarNombre(nombre);
		this.asignarRaza(n);
		this.asignarVida(0);
		this.asignarClase(n2);
	}

//getter y setter name
	public String getName(){return name;}

	public void asignarNombre(String nombre){name = nombre;}

//getter y setter raza
	public void setRaza(String g){raza = g;}

	public String getRaza(){return raza;}

//getter y setter cla
	public void setCla(String g){cla = g;};

	public String getCla(){return cla;}

//getter y setter vida
	public int getVida(){ return vida;}

	public void asignarVida(int a){
		if (vida != 0){
			vida = a;
		}
		else{
			Raza b = this.getRace();
			vida = 10 + b.getCon();
		}
	}

//getter y setter race
	public Raza getRace(){return race;}

	public void asignarRaza(int a){
		if (a == 1){
			race = new Humano();
			race.crearRaza();
			this.setRaza("Humano");
		}
		else if(a == 2){
			race = new Elfo();
			race.crearRaza();
			this.setRaza("Elfo");
		}
		else if (a == 3){
			race = new Enano();
			race.crearRaza();
			this.setRaza("Enano");
		}
		else if (a == 4){
			race = new Orco();
			race.crearRaza();
			this.setRaza("Orco");
		}
		else{
			System.out.println("APRENDE A SELECCIONAR UNA OPCION");
		}
	}

//getter y setter clase
	public Clase getClase(){return clase;}

	public void asignarClase(int a){
		if (a == 1){
			clase = new Barbaro();
			clase.crearClase();
			this.setCla("Barbaro");
		}
		else if(a == 2){
			clase = new Picaro();
			clase.crearClase();
			this.setCla("Picaro");
		}
		else if (a == 3){
			clase = new Mago();
			clase.crearClase();
			this.setCla("Mago");
		}
		else if (a == 4){
			clase = new Clerigo();
			clase.crearClase();
			this.setCla("Clerigo");
		}
		else{
			System.out.println("APRENDE A SELECCIONAR UNA OPCION");
		}			
	}

//metodos para tener y manejar la informacion de las variables clase y race

	/*
	void setModopelea(int n)

	Recibe un int n tal que modifica de la variable Clase clase el modo de pelea. 
	*/
	public void setModopelea(int n){
		Clase c = this.getClase();
		if(n == 0) c.setModopelea(n);
		else c.defender();
	}

	/*
	int getModopelea(int n)

	retorna el objeto Clase C.
	*/
	public int getModopelea(){
		Clase c = this.getClase();
		return c.getModopelea();
	}

	/*
	void setDefensa()

	Llama al metodo defender de la variable c que es objeto tipo clase (Picaro, Clerigo, Barabaro o Mago).
	*/
	public void setDefensa(){
		Clase c = this.getClase();
		c.defender();
	}

	/*
	int getStr()

	Llama al metodo getStr() de la variable b la cual es tipo raza (Humano, Enano, Orco o Elfo).
	Retorno el valor de la fuerza de este.
	*/
	public int getStr(){
		Raza b = this.getRace();
		return b.getStr();
	}

	/*
	int getDex()

	Llama al metodo getDex() de la variable b la cual es tipo raza (Humano, Enano, Orco o Elfo).
	Retorno el valor de la destreza de este.
	*/
	public int getDex(){
		Raza b = this.getRace();
		return b.getDex();
	}

	/*
	int getCon()

	Llama al metodo getCon() de la variable b la cual es tipo raza (Humano, Enano, Orco o Elfo).
	Retorno el valor de la constitucion de este.
	*/
	public int getCon(){
		Raza b = this.getRace();
		return b.getCon();
	}

	/*
	int CartaTrampa(int n, int dado)

	Recibe un entero n y dado, tal que con estos se utilicen como parametros para llamar al metodo habilidad 
	de la variable b la cual es tipo raza (Humano, Enano, Orco o Elfo).
	Retorna un entero el cual es el resultado de la activacion de la habilidad de la raza.
	*/
	public int cartaTrampa(int n, int dado){
		Raza b = this.getRace();
		return b.habilidad(n, dado);
	}
	
	/*
	void ataque(Enemigo enem, Jugador player)

	Recibe como parametro un enemigo y un jugador, para asi utilizarlos de parametro para realizar un ataque del
	jugador hacia el enemigo de la respectiva clase del jugador.
	*/
	public void ataque(Enemigo enem, Jugador player) {
		Clase c = player.getClase();
		c.ataque(enem, player, 0);
	}
	
	/*
	int getArmor()

	Llama al metodo getArmadura() de la variable Clase clase para asi de este forma obtner la armadura del personaje.
	Retorna el valor de la armadura.
	*/
	public int getArmor() {
		Clase c = this.getClase();
		return c.getArmadura();
	}
}
	


