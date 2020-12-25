#include <bits/stdc++.h>
#include <time.h>
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


string tococomple(string a)
{
	int flag = 0;

	for (int i = (int)a.length()-1; i >= 0; --i)
	{
		if(flag == 0 and a[i] == '1')flag = 1;
		else if (flag == 1){
			if(a[i] == '0') a[i] = '1';
			else a[i] = '0';
		}
	}
	return a;
}


string bigresta(string p1, string p2, string p3)
{
	string c, mp1,mp2;
	
	if(p3.length() >= p2.length())
	{
		if (p3.length() >= p1.length())
		{
			while(p3.length() > p2.length()) p2 = "0"+p2;
			while(p3.length() > p1.length()) p1 = "0"+p1;
		}
		else
		{
			while(p1.length() > p2.length()) p2 = "0"+p2;
			while(p1.length() > p3.length()) p3 = "0"+p3;
		}
	}
	else if (p1.length() >= p2.length())
	{
		while(p1.length() > p2.length()) p2 = "0"+p2;
		while(p1.length() > p3.length()) p3 = "0"+p3;
	}
	else
	{
		while(p2.length() > p3.length()) p3 = "0"+p3;
		while(p2.length() > p1.length()) p1 = "0"+p1;
	}


	p3 = "0"+p3;
	p2 = "0"+p2;
	p1 = "0"+p1;

	mp1 = tococomple(p1);
	mp2 = tococomple(p2);

	c = "0" + sumita(p3,mp1).erase(0,1);
	while(c.length() < mp2.length())c = "0" + c;
	c = sumita(c,mp2);
	return c.erase(0,1);
}

string mul(string x, string y)
{
	int largo, down, up;
	string xl,xr,yl,yr,p1,p2,p3,p4,sumax,sumay;

	if (x.length() == 1 and y.length() == 1)
	{
		if(x[0]== '0' or y[0] == '0') return "0";
		else return "1";
	}
	else if(x.length() == 0 or y.length() == 0) return "0";

	while(x.length() < y.length()) x = "0" + x;
	while(y.length() < x.length()) y = "0" + y;

	largo = x.length();
	down = largo/2;
	up = largo - down;

	xl = x.substr(0,down);
	xr = x.substr(down,up);
	yl = y.substr(0,down);
	yr = y.substr(down,up);

	p1 = mul(xl,yl);
	p2 = mul(xr,yr);
	p3 = mul(sumita(xl,xr),sumita(yl,yr));

	p4 = bigresta(p1,p2,p3);

	for (int i = 0; i < up; ++i) p4 = p4 + "0"; 
	for (int i = 0; i < up*2; ++i) p1 = p1 + "0";

	return sumita(sumita(p1,p4), p2);
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