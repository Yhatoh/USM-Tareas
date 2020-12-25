#include <bits/stdc++.h>
using namespace std;

typedef vector< int > vi;
typedef vector< vi > grafo;

int finde;
int conte;

/*
Esta función lo que hace es encontrar los puentes aplicando un dfs modificado

padre    es un arreglo el cual indica en la posición i, el nodo del cuál visitamos el nodo i
iter     es un arreglo el cual indica en la posición i, cuantas veces aplique dfs en el nodo i
mini     es un arreglo el cual indica en la posición i, el menor valor de iter que puede alcanzar 
	     el nodo i, sin contar su antecesor directo. Al principio mini[i] = iter[i]
grafo    es el grafo correspondiente
nodoreal es un map el cual tiene pares (nodoreal, posición en el grafo), esto es porque se puede 
      	 dar el caso donde tenga 4 nodos y los numeros sean 4000, 2, 7347 y 1
a        es el nodo en el arreglo
real     es el valor real del nodo a
*/

void damepuente(int a,int real, vi &padre, vi &iter, vi &mini, grafo &g, map<int, int> &nodoreal)
{
	int z, i;
	mini[a] = iter[a] = conte++;
	z = g[a].size();
	for (i = 0; i < z; ++i)
	{
		int v = g[a][i];
		
		std::map<int,int>::iterator it;
		it = nodoreal.find(v);
		if (iter[it->second] == -1)
		{
			padre[it->second] = a;
			damepuente(it->second,it->first, padre, iter, mini, g,nodoreal);
			if (mini[it->second] > iter[a] && a != it->second)
			{	 
				cout << real << " " << it->first << "\n";
				finde++;
			}
			mini[a] = min(mini[a], mini[it->second]);
		}
		else if (it->second != padre[a]){
			mini[a] = min(mini[a], iter[it->second]);
		}
	}
}


int main() {
    int n,a,i;
	std::map<int,int>::iterator itr;
    while(cin >> n)
    {	
		map<int,int> nodoreal;
    	vi padre(n, -1);
    	vi iter(n, -1);
    	vi mini(n, 0);
    	grafo ga(n);

    	string line;
    	getline( cin, line );
    	istringstream is( line );

    	for (i = 0; i < n; ++i)
    	{
    		getline( cin, line );
    		istringstream is( line );
			is >> a;
			nodoreal.insert(pair<int,int>(a,i));
    		while(is >> a){
				ga[i].push_back(a);
			}
    	}
    	conte = 0;
    	finde = 0;
    	for (itr = nodoreal.begin(); itr != nodoreal.end(); ++itr) {
			if(iter[itr->second] == -1) damepuente(itr->second,itr->first, padre, iter, mini, ga, nodoreal);
		}	
    	if (finde == 0)cout << "No Existe Corte.\n";
    	cout << "\n";
    }
    return 0;
}