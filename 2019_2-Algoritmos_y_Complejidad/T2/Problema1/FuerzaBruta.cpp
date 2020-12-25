#include <bits/stdc++.h>
using namespace std;

string sumita(string a, string b)
{
	string c;
	int carry = 0;
	
	while(a.length() < b.length()) a = "0" + a;
	while(b.length() < a.length()) b = "0" + b;

	for (int i = (int)a.length()-1; i >= 0 ; --i)
	{
		if (a[i] == '1' and b[i] == '1')
		{
			if (carry) c = "1" + c;
			else
			{
				c = "0" + c;
				carry = 1;
			}
		}
		else if (a[i] == '1' or b[i] == '1')
		{
			if(carry) c = "0" + c;
			else c = "1" + c;
		}
		else
		{
			if(carry)
			{
				carry = 0;
				c = "1" + c;
			}
			else c = "0"+ c;
		}
	}
	if(carry) c = "1" + c;
	return c;
}

string mul(string x, string y)
{
	int i;
	string c,d,e;

	for (i = (int)x.length()-1 ; i >= 0; --i)
	{
		if (x[i] == '1')
		{
			c = y + d;
			e = sumita(e,c);	
		}
		d = d + "0";
	}
	 return e;
}	


int main()
{
	int a;

	while(cin >> a)
	{
		string x,y,final;
		cin >> x >> y;
		final = mul(x,y);
		while(final[0] == '0' and final.length() > 1)final.erase(0,1);
		cout << final << endl;
		
	}
	return 0;
}	