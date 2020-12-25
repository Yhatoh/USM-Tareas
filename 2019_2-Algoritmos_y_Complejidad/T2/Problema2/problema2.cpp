#include<bits/stdc++.h>

using namespace std;

int main(){
	ios_base::sync_with_stdio(false); 
	cin.tie(NULL);
	cout.setf(ios::fixed);
	cout.precision(4);
	string arbol;
	int n, i, nodo, pos, padre, pos_padre, p, q, v;

	while(cin >> n){
		cin >> arbol;

		vector<vector<int>> edd_rara(n);
		padre = -1;
		pos_padre = -1;
		pos = 0;
		i = 0;
		//Armamiento de la mejor estructura de datos
		while(i < n){
			if(arbol[pos] == '('){
				cin >> nodo;

				edd_rara[i].push_back(nodo);
				edd_rara[i].push_back(padre);
				edd_rara[i].push_back(pos_padre);

				padre = nodo;
				pos_padre = i;
				pos++;
				i++;
			}
			else{
				padre = edd_rara[pos_padre][1];
				pos_padre = edd_rara[pos_padre][2];
				pos++;

				while(arbol[pos] != '('){
					pos++;
					padre = edd_rara[pos_padre][1];
					pos_padre = edd_rara[pos_padre][2];
				}
			}
		}
		//Escribimiento de la soluciones de los consultamientos
		cin >> q;
		cout << q << "\n";
		while(q--){
			cin >> v >> p;
			while(edd_rara[v][1] >= p) v = edd_rara[v][2];
			cout << v << "\n";
		}
	}
	return 0;
}