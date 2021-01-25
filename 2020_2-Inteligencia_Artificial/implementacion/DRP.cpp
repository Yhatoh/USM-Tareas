#include<bits/stdc++.h>
using namespace std;

typedef pair< int, int > pii;
typedef pair< double, double > pdd;
typedef vector< int > vi;
typedef vector< bool > vb;
typedef vector< pdd > vpdd;
typedef vector< vi > vvi;
typedef vector< set< int > > vsi;

//variables globales utiles para guardar la informacion
vi maxWhereInstall;
vi maxCovered;
vpdd coords, coordsNotInstalled, coordsInstalled;
int events, maxX;
double budget, maxCost;

//copia los valores de un vector a otro
void copy(vi &s, vi &d){
	int i;
	for(i = 0; i < (int) s.size(); i++){
		d[i] = s[i];
	}
}

//calcula distancia entre dos puntos
double distance(pdd a, pdd b){
	return sqrt((a.first - b.first) * (a.first - b.first) + (a.second - b.second) * (a.second - b.second));
}

//sirve para printear las coordenadas donde estan los AED de la solucion
//si el AED estaba isntaldo previamente se ve si es que se movio o no y se printea al respecto
//si no se dice que es nuevo y se muestra donde se instalo
void printCoords(vi &whereInstall, vpdd &coords, vpdd &coordsInstalled){
	int i;
	int nInstalled = (int) coordsInstalled.size();
	for(i = 0; i < (int) whereInstall.size(); i++){
		if(whereInstall[i] != -1){
			if(i < nInstalled){
				if(whereInstall[i] >= nInstalled){
					cout << "REPOSICIONADO (" << coordsInstalled[i].first << "," << coordsInstalled[i].second << ") => ";
					cout << "(" << coords[whereInstall[i]].first << "," << coords[whereInstall[i]].second << ")\n";
				} else {
					cout << "NO SE MOVIO " << "(" << coords[whereInstall[i]].first << "," << coords[whereInstall[i]].second << ")\n";
				}
			} else {
				cout << "NUEVO " << "(" << coords[whereInstall[i]].first << "," << coords[whereInstall[i]].second << ")\n";
			}
		}
	}
}

int DRP2(vi &maxWhereInstall, vi &maxCovered, vi whereInstall, vi checker, vi covered, vi used, vvi &canCover, vvi possiblePos, vsi &CBJ, double budget, double* maxCost, double cost, int* maxX, int x, int AED, int instanced){
	int i, j;

	//si es que es la primera isntancia no se revisa puesto que no hay todavia nada instanciado
	if(instanced != -1){
		//se guarda que la posicion en que se instancia haya sido usada
		used[whereInstall[instanced]] = 1;
		//cuenta cuantos puntos nuevos han sido cubiertos debido a esta instanciacion
		for(i = 0; i < (int) canCover[whereInstall[instanced]].size(); i++){
			x += 1 & (covered[canCover[whereInstall[instanced]][i]] + 1);
			covered[canCover[whereInstall[instanced]][i]] = 1;
		}
		//realiza proceso de forward checking donde elimina del dominio de la variables relacionadas a instanciar
		//toda coordenada ya usada, puesto que no puede instanciarse dos AED en la misma coordenada
		for(i = instanced + 1; i < (int) possiblePos.size(); i++){
			for(j = 0; j < (int) possiblePos[i].size(); j++){
				if(used[possiblePos[i][j]]){
					possiblePos[i][j] = possiblePos[i][(int) possiblePos[i].size() - 1];
					possiblePos[i].pop_back();
					j--;
				}
			}
		}
	}

	//restriccion del presupuesto, si el presupuesto es menor quiere decir que no se puede instanciar mas, entonces 
	//se detiene y agrega toda variable que haya sido instancia a conjunto de restricciones puesto que toda variables instanciada
	//tiene relacion con esta nueva debido a que el presupuesto es una restriccion global
	if(budget < cost){
		for(i = 0; i < instanced; i++){
			if(checker[i]){
				CBJ[instanced].insert(i);
			}
		}
		return instanced;
	}

	//si es que la cantidad de puntos cubiertos en el actual es mayor que el registrado
	//entonces se reemplaza la solucion actual con la nueva
	if(*maxX < x){
		*maxX = x;
		*maxCost = cost;
		copy(covered, maxCovered);
		copy(whereInstall, maxWhereInstall);
	}

	//en el caso que haya llegado al limite de AED posibles
	if(AED == (int) whereInstall.size()){
		return -1;
	}

	
	//si es un AED instlado previamente se prueba el caso donde se deja en la misma posicion
	int wasMoved = 0;
	if(whereInstall[AED] != -1){
		wasMoved = 1;
		DRP2(maxWhereInstall, maxCovered, whereInstall, checker, covered, used, canCover, possiblePos, CBJ, budget, maxCost, cost, maxX, x, AED + 1, AED);
	}
	//luego se revisan todo el resto de posiciones posibles del AED a instanciar
	int auch;
	for(i = 0; i < (int) possiblePos[AED].size(); i++){
		double aux;
		//se guarda la posicion del AED y se dice que se instancio (esto es para CBJ mas tarde)
		whereInstall[AED] = possiblePos[AED][i];
		checker[AED] = 1;
		//si era un AED instalado previamente al coste se le suma 0.2, si no se le suma 1
		if(wasMoved) aux = cost + 0.2;
		else aux = cost + 1;
		//se almacena la variable que no tenia mas valores a instancias (para aplicar el restorno guiado)
		auch = DRP2(maxWhereInstall, maxCovered, whereInstall, checker, covered, used, canCover, possiblePos, CBJ, budget, maxCost, aux, maxX, x, AED + 1, AED);
		if(auch != -1 && auch != AED){
			//se revisa si el AED actual esta en el conjunto, si lo esta quiere deci que esta es la variable mas reciente instanciada
			//con restricciones con la variable que no tuvo mas opciones
			//si no esta entonces se sigue retornando debido a que no es la variable que causo el dilema (CBJ SI)
			if(CBJ[auch].find(AED) == CBJ[auch].end()){
				return auch;
			} else {
				CBJ[auch].clear();
			}
		}
	}
	return instanced;
}

//esta funcion muestra el resultado encontrado en el momento que se hace ctrl+C
void ctrlC(int s){
	cout << "Quantity OHCA covered: " << maxX << "\n";
	cout << "Percentage: " << ((double) maxX / (double) events) * 100.0 << "%\n";
	cout << "Budget left: " << budget - maxCost << "\n";
	printCoords(maxWhereInstall, coords, coordsInstalled);
	exit(1); 
}

int main(){
	//handler 
	struct sigaction sigIntHandler;

	sigIntHandler.sa_handler = ctrlC;
	sigemptyset(&sigIntHandler.sa_mask);
	sigIntHandler.sa_flags = 0;

	sigaction(SIGINT, &sigIntHandler, NULL);
	
	int i, j;

	double r;
	cin >> events >> budget >> r;

	//data
	vi used(events, 0);
	vvi canCover(events);

	//Aqui se lee las coordenadas y si esta instalado o no
	for(i = 0; i < events; i++){
		//read coordinates
		int x, y, exist;
		cin >> x >> y >> exist;
		if(exist){
			coordsInstalled.push_back(pdd(x, y));
		} else {
			coordsNotInstalled.push_back(pdd(x, y));
		}
	}
	int nInstalled = (int) coordsInstalled.size();
	int nNotInstalled = (int) coordsNotInstalled.size();

	//Aqui se crea el arreglo de la representacion
	coords.insert(coords.end(), coordsInstalled.begin(), coordsInstalled.end());
	coords.insert(coords.end(), coordsNotInstalled.begin(), coordsNotInstalled.end());

	//useful for answer
	vi covered(events, 0);
	vi whereInstall(nInstalled + (int) budget, -1);
	vvi possiblePos(nInstalled + (int) budget);
	vsi CBJ(nInstalled + (int) budget);
	vi checker(nInstalled + (int) budget, 0);
	maxWhereInstall.assign(nInstalled + (int) budget, 0);
	maxCovered.assign(events, 0);


	//Se define el valor inicial de los AED ya instalados, previamente
	for(i = 0; i < nInstalled; i++){
		whereInstall[i] = i;
	}
	//se definen todas las posibles posiciones de cada AED
	for(i = 0; i < nNotInstalled; i++){
		for(j = 0; j < nInstalled + (int) budget; j++){
			possiblePos[j].push_back(i + nInstalled);
		}
	}
	
	//precalcula cuales coordenadas pueden ser cubiertas por otra coordenada 
	for(i = 0; i < events; i++){
		canCover[i].push_back(i);
		for(j = i + 1; j < events; j++){
			if(distance(coords[i], coords[j]) <= r){
				canCover[i].push_back(j);
				canCover[j].push_back(i);
			}
		}
	}

	maxX = 0;
	maxCost = 0.0;
	DRP2(maxWhereInstall, maxCovered, whereInstall, checker, covered, used, canCover, possiblePos, CBJ, budget, &maxCost, 0, &maxX, 0, 0, -1);
	cout << "Quantity OHCA covered: " << maxX << "\n";
	cout << "Percentage: " << ((double) maxX / (double) events) * 100.0 << "%\n";
	cout << "Budget left: " << budget - maxCost << "\n";
	printCoords(maxWhereInstall, coords, coordsInstalled);
	return 0;
}