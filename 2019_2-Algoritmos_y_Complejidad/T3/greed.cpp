#include <bits/stdc++.h>
using namespace std;


typedef pair <int,int> ii;
typedef vector< ii > pares;


int main()
{
    int a,b,c,d,f;
    int pos;
    unsigned int i;
    int resta;
    cin >> a;
    pares par(a);
    pares aceptados;
    for ( i = 0; (int)i < a; i++)
    {
        cin >> b >> c;
        par[i].first = b;
        par[i].second = c;
    }
    
    while (par.size() != 0)
    {
        for (i = 0; i < par.size(); i++)
        {
            resta = par[i].second- par[i].first;
            if(i == 0 || resta < d)
            {
                pos = i;
                d = resta;
            }
        }
        f=1;
        for (i = 0; i < aceptados.size(); i++)
        {
            if(!((par[pos].first > aceptados[i].second && par[pos].second > aceptados[i].second) || (par[pos].second < aceptados[i].first && par[pos].first < aceptados[i].second)))
            {
            	f = 0;
            	break;
            }
        }
        if(f)
        {
            aceptados.push_back(par[pos]);
        }
        par.erase(par.begin()+pos);
    }
    sort(aceptados.begin(),aceptados.end());

    cout << aceptados.size() << endl;
    for ( i = 0; i < aceptados.size(); i++)
    {
        cout << aceptados[i].first << " " << aceptados[i].second << endl;
    }
    
    


    return 0;
}
