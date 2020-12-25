from django.shortcuts import render, redirect
from django.contrib import auth
from django.contrib.auth.models import User
from .forms import CartaForm
from .models import Mazo, Carta, Equipo, Fotito
from django.contrib.auth.decorators import login_required
from datetime import datetime

import random

def inicio(request):
	if request.method == 'POST':
		if ("username" in request.POST.keys()) and ("password" in request.POST.keys()):
			user = auth.authenticate(username= request.POST['username'], password= request.POST['password'])
			if user is not None and user.is_active:
				auth.login(request, user)
				return redirect("/home/")
			else:
				contexto = {"error":"error"}
				return render(request, "inicio.html", contexto)
		elif ("nusername" in request.POST.keys()) and ("npassword" in request.POST.keys()):
			if(request.POST['npassword'] == request.POST['npassword2']):
				try:
					User.objects.create_superuser(username=request.POST['nusername'],password=request.POST['npassword'],email='',first_name="/media/maguito1.png" )
					return render(request, 'inicio.html')
				except:
					return render(request,"inicio.html",{'registro':'registro','yacreado':'yacreado'})
			return render(request,"inicio.html",{'registro':'registro','contraequiv':'contraequiv'})
		elif "regi" in request.POST.keys():
			if request.POST['regi'] == 'reg':
				registro = {'registro':'registro'}
				return render(request,"inicio.html", registro)
		elif "conectt" in request.POST.keys():
			if request.POST['conectt'] == 'conectt':
				return render(request,"inicio.html")
		else:
			contexto = {"error":"error"}
			return render(request, "inicio.html", contexto)
	return render(request, 'inicio.html')

@login_required(login_url='/')
def home(request):
	now = datetime.now()
	context={'usuario':request.user,'time':now}
	if request.method == 'POST':
		if ("cerrar" in request.POST.keys()):
			auth.logout(request)
			return redirect("/")
		elif("cartas" in request.POST.keys()):
			return redirect("/mazo/")
		elif("sobres" in request.POST.keys()):
			return redirect("/sobre/")
		elif("equipos" in request.POST.keys()):
			return redirect("/equipo/")
		elif("perfil" in request.POST.keys()):
			return redirect("/perfil/")
		elif("juegito" in request.POST.keys()):
			return redirect("/jugar/")
	return render(request, 'home.html',context)

@login_required(login_url='/')
def mazo(request):
	userr = request.user
	lista = list(Mazo.objects.filter(jugador=userr.username))
	i = 0

	le = len(lista)
	lista2 = []
	lista3 = []
	while(i < le/3):
		lista2.append(lista[0])
		del lista[0]
		i+=1
	i = 0
	while(i < le/3):
		lista3.append(lista[0])
		del lista[0]
		i+=1
	
	context = {'lista':lista,'lista2':lista2,'lista3':lista3}
	
	
	return render(request, 'mazo.html',context)

@login_required(login_url='/')
def perfil(request):
	userr=request.user
	if(request.method=='POST'):
		user =list(User.objects.filter(username=request.user.username))
		userr = user[0]
		userr.first_name=request.POST['aaa']
		userr.save()
	return render(request, 'perfil.html',{'user':userr})

@login_required(login_url='/')
def info(request):
	return render(request, 'info.html')

@login_required(login_url='/')
def jugar(request):
	if(request.method == 'POST'):
		if("mazo" in request.POST.keys()):
			return redirect("/equipo/")
		elif("juego" in request.POST.keys()):
			l_user = list(User.objects.all())
			l_user.remove(request.user)
			
			a = random.randint(0,len(l_user)-1)
			ene = l_user[a]
			jug = request.user

			equipo_ene = Equipo.objects.filter(jugador=ene.username)
			while(list(equipo_ene) == []):
				l_user.remove(ene);
				a = random.randint(0,len(l_user)-1)
				ene = l_user[a]
				equipo_ene = Equipo.objects.filter(jugador=ene.username)
			equipo_jug = Equipo.objects.filter(jugador=jug.username)
			if(list(equipo_jug) == []):
				return render(request, 'jugar.html', {'nomazo': 'nomazo'})
			cartas_ene,cartas_ene2=obtenernombreseimagente([Carta.objects.filter(id=equipo_ene[0].carta1)[0],Carta.objects.filter(id=equipo_ene[0].carta2)[0],Carta.objects.filter(id=equipo_ene[0].carta3)[0],Carta.objects.filter(id=equipo_ene[0].carta4)[0],Carta.objects.filter(id=equipo_ene[0].carta5)[0]])
			cartas_jug,cartas_jug2=obtenernombreseimagente([Carta.objects.filter(id=equipo_jug[0].carta1)[0],Carta.objects.filter(id=equipo_jug[0].carta2)[0],Carta.objects.filter(id=equipo_jug[0].carta3)[0],Carta.objects.filter(id=equipo_jug[0].carta4)[0],Carta.objects.filter(id=equipo_jug[0].carta5)[0]])

			cartas_jug =",".join(cartas_jug)
			cartas_ene =",".join(cartas_ene)
			
			dicc = {'jug':jug.username,'ene':ene.username,'cartajug0':cartas_jug2[0],'cartajug1':cartas_jug2[1],'cartajug2':cartas_jug2[2],'cartajug3':cartas_jug2[3],'cartajug4':cartas_jug2[4],'cartaene0':cartas_ene2[0],'cartaene1':cartas_ene2[1],'cartaene2':cartas_ene2[2],'cartaene3':cartas_ene2[3],'cartaene4':cartas_ene2[4],'cartajug':cartas_jug,'cartaene':cartas_ene,'pv':0,'pc':0}
			
			return render(request, 'jugar.html', dicc)

		elif("carta" in request.POST.keys()):
			jug = request.user
			carta = request.POST["carta"]
			pc = int(request.POST["pc"])
			pv = int(request.POST["pv"])
			ene = request.POST["ene"]
			cartas_jug = request.POST["cartajug"].split(",")
			cartas_ene = request.POST["cartaene"].split(",")

			cartas_jug = obtenercartas(cartas_jug)
			cartas_ene = obtenercartas(cartas_ene)

			carta_jug = cartas_jug[int(carta)]

			i = 0
			lista_pos = []
			for carta1 in cartas_ene:
				if(carta1 != "-1" and carta1 != "-2"):
					lista_pos.append(i)
				i+=1
			pos = lista_pos[random.randint(0,len(lista_pos)-1)]
			carta_ene = cartas_ene[pos]

			if(carta_jug.ataque-carta_ene.defensa >= 0 and carta_ene.ataque-carta_jug.defensa >= 0):
				if(random.randint(0,1) == 1):
					cartas_ene[pos] = "-1"
					cartas_jug[int(carta)] = "-2"
					pc += 1
				else:
					cartas_ene[pos] = "-2"
					cartas_jug[int(carta)] = "-1"
					pv += 1
			elif(carta_jug.ataque-carta_ene.defensa >= 0):
				cartas_ene[pos] = "-2"
				cartas_jug[int(carta)] = "-1"
				pv += 1
			elif(carta_ene.ataque-carta_jug.defensa >= 0):
				cartas_ene[pos] = "-1"
				cartas_jug[int(carta)] = "-2"
				pc += 1
			else:
				turnos_j = int(carta_ene.defensa/carta_jug.ataque)+1
				turnos_e = int(carta_jug.defensa/carta_ene.ataque)+1

				p = random.randint(0,1)

				#parte jugador
				if(p == 0):
					if(turnos_j >= turnos_e):
						cartas_ene[pos] = "-2"
						cartas_jug[int(carta)] = "-1"
						pv += 1
					else:
						cartas_ene[pos] = "-1"
						cartas_jug[int(carta)] = "-2"
						pc += 1
				else:
					if(turnos_e >= turnos_j):
						cartas_ene[pos] = "-1"
						cartas_jug[int(carta)] = "-2"
						pc += 1
					else:
						cartas_ene[pos] = "-2"
						cartas_jug[int(carta)] = "-1"
						pv += 1

			cartas_jug,cartas_jug2 = obtenernombreseimagente(cartas_jug)
			cartas_ene,cartas_ene2 = obtenernombreseimagente(cartas_ene)

			cartas_jug =",".join(cartas_jug)
			cartas_ene =",".join(cartas_ene)
			if(pc == 3):
				dicc = {'jug': jug, 'ene': ene, 'pc': pc, 'pv': pv, 'victoriac':'victoriac'}
			elif(pv == 3):
				dicc = {'jug': jug, 'ene': ene, 'pc': pc, 'pv': pv, 'victoriap':'victoriap'}
			else:
				dicc = {'jug':jug,'ene':ene,'nombrevic':carta_jug.nombre,'nombreper':carta_ene.nombre,'cartajug0':cartas_jug2[0],'cartajug1':cartas_jug2[1],'cartajug2':cartas_jug2[2],'cartajug3':cartas_jug2[3],'cartajug4':cartas_jug2[4],'cartaene0':cartas_ene2[0],'cartaene1':cartas_ene2[1],'cartaene2':cartas_ene2[2],'cartaene3':cartas_ene2[3],'cartaene4':cartas_ene2[4],'cartajug':cartas_jug,'cartaene':cartas_ene,'pv':pv,'pc':pc}
			return render(request, 'jugar.html', dicc)

	return render(request, 'jugar.html')

@login_required(login_url='/')
def sobre(request):
	if request.method == 'POST':
		if('sobre' in request.POST.keys()):
			lista = list(Carta.objects.all())
			i = 0
			while(i < len(lista)):
				i+=1
			i = 0
			l1 = []
			l2 = []
			while(i < int(request.POST['sobre'])):
				a = random.randint(0,len(lista)-1)
				ob = list(Mazo.objects.filter(jugador=request.user.username,carta=lista[a]))
				if(ob == []):
					mazos = Mazo()
					mazos.jugador=request.user.username
					mazos.carta = lista[a]
					mazos.cantidad=1
					mazos.save()
				else:
					Mazo.objects.filter(jugador=request.user.username,carta=lista[a]).update(cantidad= (ob[0].cantidad)+1)
				if(i%2 == 0):
					l1.append(lista[a])
				elif(i%2 == 1):
					l2.append(lista[a])
				
				i+=1

			return render(request, 'sobre.html',{'l1':l1,'l2':l2})

	return render(request, 'sobre.html',{'a':'a'})

@login_required(login_url='/')
def equipo(request):

	lista = list(Mazo.objects.filter(jugador=request.user.username))
	if(len(lista) >= 5):
		form=CartaForm()

		if(request.method == 'POST'):
			if(request.POST['carta1'] != request.POST['carta2'] and request.POST['carta1'] != request.POST['carta3'] and request.POST['carta1'] != request.POST['carta4'] and request.POST['carta1'] != request.POST['carta5'] and request.POST['carta2'] != request.POST['carta3'] and request.POST['carta2'] != request.POST['carta4'] and request.POST['carta2'] != request.POST['carta5'] and request.POST['carta3'] != request.POST['carta4'] and request.POST['carta3'] != request.POST['carta5'] and request.POST['carta4'] != request.POST['carta5']):
				form=CartaForm(request.POST or None)
				if form.is_valid():
					form.save()
					return render(request,'equipo.html',{'lista':lista,'form':form,'user':request.user})
			return render(request,'equipo.html',{'lista':lista,'form':form,'user':request.user,'error':'error'})
		return render(request,'equipo.html',{'lista':lista,'form':form,'user':request.user})
	if(request.method == 'POST'):
		if("sobres" in request.POST.keys()):
			return redirect("/sobre/")
	return render(request, 'equipo.html', {'nocartas':'nocartas'})

def obtenercartas(cartas):
	cartasr = []
	for carta in cartas:
		if(carta == "-1"):
			cartasr.append("-1")
		elif(carta == "-2"):
			cartasr.append("-2")
		else:
			cartasr.append(Carta.objects.filter(nombre=carta)[0])
	return cartasr

def obtenernombreseimagente(cartas):
	cartasr = [[],[]]
	for carta in cartas:
		if(carta == "-1"):
			cartasr[0].append("-1")
			cartasr[1].append("-1")
		elif(carta == "-2"):
			cartasr[0].append("-2")
			cartasr[1].append("-2")
		else:
			cartasr[0].append(carta.nombre)
			cartasr[1].append(carta.imagen)
	'''
	cartasr[0].append(cartas[0].nombre)
	cartasr[0].append(cartas[1].nombre)
	cartasr[0].append(cartas[2].nombre)
	cartasr[0].append(cartas[3].nombre)
	cartasr[0].append(cartas[4].nombre)

	cartasr[1].append(cartas[0].imagen)
	cartasr[1].append(cartas[1].imagen)
	cartasr[1].append(cartas[2].imagen)
	cartasr[1].append(cartas[3].imagen)
	cartasr[1].append(cartas[4].imagen)
	'''
	return cartasr
