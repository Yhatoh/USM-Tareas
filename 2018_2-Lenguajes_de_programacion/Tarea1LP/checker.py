import re

###
# Nombre: afterciclo
# Input lista, er, error
#   lista: Lista que posee todo el codigo a revisar.
#   er: Lista de los compilado de las expresiones regulares generales.
#   error: Conjunto con los errores encontrados hasta el momento.
# Descripcion: Revisa que todos los ciclos no se cierren en la siguiente linea.
# Return: Conjunto con los errores encontrados agregados.
###
def afterciclo(lista, er, error):
    l = []
    i = 0
    for linea in lista:
        if er[10].match(linea) != None:
            l.append((er[10].match(linea).group(2), i))
        i+=1
    for a in l:
        if er[11].match(lista[a[1]+1]) != None:
            if er[11].match(lista[a[1]+1]).group(2) == a[0]:
                error.add((lista[a[1]], a[1]))
                error.add((lista[a[1]+1], a[1]+1))
    return error

###
# Nombre: ciclos
# Input lista, er, error
#   lista: Lista que posee todo el codigo a revisar.
#   er: Lista de los compilado de las expresiones regulares generales.
#   error: Conjunto con los errores encontrados hasta el momento.
# Descripcion: Revisa que todos los ciclos se encuentren correctamente estructurados
# Return: Conjunto con los errores encontrados agregados.
###
def ciclos(lista, er, error):
    l = []
    pros = []
    errorpos = []
    i = 0
    for linea in lista:
        if er[10].match(linea) != None:
            l.append((er[10].match(linea).group(2), i, "ini"))
        if er[11].match(linea) != None:
            l.append((er[11].match(linea).group(2), i, "sto"))
        i += 1
    ledit = []
    for x in l:
        ledit.append(x)

    if len(l) > 0:
        for x in l:
            if x[2] == "ini":
                pros.append(x)
                ledit.remove(x)

            elif len(pros)> 0:
                if pros[len(pros)-1][0] == x[0]:
                    del pros[len(pros)-1]
                    ledit.remove(x)
                else:
                    flag = False
                    for z in pros:
                        if z[0] == x[0]:
                            flag = True
                            pre = z
                    if flag:
                        errorpos.append(pros[len(pros)-1][1])
                        ledit.remove(x)
                        pros.remove(pre)
                    else:
                        errorpos.append(x[1])
                        ledit.remove(x)

        if len(pros) > 0:
            for a in pros:
                error.add((lista[a[1]] , a[1]))
        if len(ledit) > 0:
            for a in ledit:
                error.add((lista[a[1]] , a[1]))
        if len(errorpos) > 0:
            for a in errorpos:
                error.add((lista[a], a))
        error = afterciclo(lista, er, error)
    return error

###
# Nombre: borly.
# Input string, er.
#   string: Linea que se encuentra antes de un "O RLY?".
#   er: Lista del compilado de las expresiones regulares generales.
# Descripcion: Retorna True, en caso de que el string haga match con alguna expresion,
#              o False en caso que no lo haga.
# Return: True o False.
###
def borly(string, er):
    if er[0].match(string) != None or er[9].match(string) != None or er[10].match(string) != None or er[11].match(string) != None or string == "\n":
        return False
    return True

###
# Nombre: afterya.
# Input string, er, er2.
#   string: Linea que se encuentra despues de un "YA RLY".
#   er: Lista del compilado de las expresiones regulares generales.
# Descripcion: Retorna True, en caso de que el string haga match con alguna expresion,
#              o False en caso que no lo haga.
# Return: True o False.
###
def afterya(string, er, er2):
    if er[0].match(string) != None or er2[2].match(string) != None:
        return False
    return True

###
# Nombre: afterno.
# Input string, er.
#   string: Linea que se encuentra despues de un "NO WAI".
#   er: Lista del compilado de las expresiones regulares generales.
# Descripcion: Retorna True, en caso de que el string haga match con alguna expresion,
#              o False en caso que no lo haga.
# Return: True o False,
###
def afterno(string, er, er2):
    if er[0].match(string) != None or er2[3].match(string) != None:
        return False
    return True

###
# Nombre: miniifdetector
# Input currif, er, er2
#   currif: Lista que posee todo el codigo de un condicional simple, es decir que solo
#           posea un solo "O RLY?" y un solo "OIC", siendo la pociocion 0 el "O RLY?".
#   er: Lista de los compilado de las expresiones regulares generales.
#   er2: Lista de los compilados de las expresiones regulares del condicional.
#   error: Conjunto con los errores encontrados hasta el momento.
# Descripcion: Retorna True, en caso de que el string haga match con alguna expresion,
#              o False en caso que no lo haga.
# Return: Conjunto con los errores encontrados agregados.
###
def miniifdetector(currif, er, er2, error):
    y = 0
    n = 0
    z = 0
    for k in currif:
        if z > 1:
            if er2[1].match(k[0]) != None:
                y += 1
                if y > 1:
                    error.add(k)
            if er2[2].match(k[0]) != None:
                n += 1
                if n > 1:
                    error.add(k)
        z += 1

    if y == 0:
        error.add(currif[len(currif)-1])
        error.add(currif[1])
        if n == 1:
            for tup in currif:
                if er2[2].match(tup[0]) != None:
                    error.add(tup)
    if n == 0:
        error.add(currif[1])
        error.add(currif[len(currif)-1])
        if y == 1:
            for tup in currif:
                if er2[1].match(tup[0]) != None:
                    error.add(tup)
    if not(borly(currif[0][0],er)):
        error.add(currif[1])
    i = 2
    if er2[1].match(currif[2][0]) == None:
        if y > 0:
            pala = "thrash"
            while er2[1].match(pala) == None:
                error.add(currif[i])
                i += 1
                pala = currif[i][0]
    for k in currif:
        if er2[1].match(k[0]) != None:
            if not(afterya(currif[currif.index(k)][0],er,er2)):
                error.add(k)
        elif er2[2].match(k[0]) != None:
            if not(afterno(currif[currif.index(k)][0],er,er2)):
                error.add(k)
    return error

###
# Nombre: incode
# Input lista, l, er, er2
#   lista: Lista que posee todo el codigo a revisar.
#   l: Lista que posee tuplas con todos los condicionales ("O RLY", "YA RLY", "NO WAY", "OIC")
#      que se encuentran en el codigo y su respectiva posicion
#   er: Lista de los compilado de las expresiones regulares generales.
#   er2: Lista de los compilados de las expresiones regulares del condicional.
# Descripcion: Crea en base a la variable lista, una lista que posea todo el codigo entre
#              el primer "O RLY?" hasta el ultimo "OIC" que se encuentren en la variable
#              lista.
# Return: Lista creada a base de la variable lista.
###
def incode(lista, l, er, er2):

    incode = []
    frly = l[0][1]
    loic = l[len(l)-1][1]
    while  loic >= frly:
        incode.append((lista[frly],frly))
        frly += 1
    return incode

###
# Nombre: ifdetector.
# Input lista, er, er2, error.
#   lista: Lista que posee todo el codigo a revisar.
#   er: Lista de los compilado de las expresiones regulares generales.
#   er2: Lista de los compilados de las expresiones regulares del condicional.
#   error: Conjunto con los errores encontrados hasta el momento.
# Descripcion: Revisa que todas las expresiones condicionales dentro del codigo se encuentren
#              estructuralmente correctas, guardando los errores que encuentre
#              lista.
# Return: Conjunto con los errores encontrados agregados o False, en caso de no haber ningun
#         condicional en el codigo.
###
def ifdetector(lista, er, er2, error):
    borrar = []
    orl = 0
    yarl = 0
    nowa = 0
    oi = 0
    l = []
    i = 0
    for linea in lista:
        if er2[0].match(linea) != None:
            l.append(("O RLY?", i))
            orl += 1
        if er2[1].match(linea) != None and orl != 0:
            l.append(("YA RLY", i))
            yarl += 1
        elif er2[1].match(linea) != None and orl == 0:
            error.add(("YA RLY", i))

        if er2[2].match(linea) != None and orl != 0:
            l.append(("NO WAI", i))
            nowa += 1
        elif er2[2].match(linea) != None and orl == 0:
            error.add(("NO WAI", i))

        if er2[3].match(linea) != None and orl != 0:
            l.append(("OIC", i))
            oi += 1
        elif er2[3].match(linea) != None and orl == 0:
            error.add(("OIC", i))
        i += 1
    if orl > 0 and oi > 0:
        incod = incode(lista, l, er, er2)
        incodecount = 0
        while 0 < len(l):
            currif = []
            if 0 < orl and 0 < oi:
                lastr =-float("inf")
                for c,pos in l:
                    pal = "wololo"
                    if c == "O RLY?":
                        lastr = pos
                i = 0
                while pal != "OIC":

                    pal = l[i][0]
                    firto = l[i][1]
                    i += 1

                currif.append((lista[lastr-1],lastr-1))
                zapato = incod.index((lista[lastr],lastr))
                enif = zapato
                while not(er2[3].match(incod[enif][0])):
                    enif += 1
                #enif = incod.index(("OIC\n",firto))
                while zapato <= enif:
                    currif.append(incod[zapato])
                    if er2[0].match(incod[zapato][0]) != None:
                         l.remove(("O RLY?",incod[zapato][1]))
                         orl -= 1
                    elif er2[1].match(incod[zapato][0]) != None:
                         l.remove(("YA RLY",incod[zapato][1]))
                         yarl -= 1
                    elif er2[2].match(incod[zapato][0]) != None:
                        l.remove(("NO WAI",incod[zapato][1]))
                        nowa -= 1
                    elif er2[3].match(incod[zapato][0]) != None:
                        l.remove(("OIC",incod[zapato][1]))
                        oi -= 1
                    zapato += 1
                i = 0
                while i < len(incod):
                    if currif[1] == incod[i]:
                        incod[i] = ("4ar\n",incod[i][1], currif[len(currif)-1][1])
                    i += 1
                i = 0
                for word in currif:
                    if 1 < i:
                        incod.remove(word)
                    i += 1
                error = miniifdetector(currif, er, er2, error)
            else:
                for w in l:
                    error.add(w)
                    l.remove(w)

    elif len(l) == 0:
        return False
    else:
        for w in l:
            error.add(w)
    anadir = []
    for t in error:
        if "4ar\n" == t[0]:
            borrar.append((t[1], t[2]))
            g = t[1]
            pol = " "
            i = t[2]
            while g <= i :
                anadir.append((lista[g], g))
                g += 1
                pol = lista[g]

    if len(anadir) != 0:
        for palabra,posicion in anadir:
            error.add((palabra, posicion))
    if len(borrar) != 0:
        for posi,posi2 in borrar:
            error.remove(("4ar\n", posi, posi2))
    return error

###
# Nombre: obh (operacion bien hecha).
# Input lista, er, er2.
#   linea: Linea que contiene operacion binaria.
#   er: Lista de los compilado de las expresiones regulares generales.
#   er2: Lista de los compilados de las bases de todas las sentencias.
# Descripcion: Revisa la valides de una linea que sea o tenga una operacion binaria.
# Return: True o False, si la sentencia esta bien o mal escrita respectivamente.
###
def obh(linea,er,er2):
    matches = []
    #print(linea)
    i = 0
    for exp in er2:
        matches.append(exp.findall(linea))
    for find in matches:
        for exp2 in find:
            if (i == 0):
                linea = re.sub(exp2," #o ", linea)
            elif (i == 1):
                linea = re.sub("".join(exp2)," #a ",linea)
            elif (i == 2):
                linea = re.sub("".join(exp2)," #d ", linea)
            elif (i == 3):
                linea = re.sub("".join(exp2)," #b ",linea)
            elif (i == 4):
                linea = re.sub("".join(exp2)," #s ", linea)
            elif (i == 5):
                linea = re.sub("".join(exp2)," #e ", linea)
            elif (i == 6):
                linea = re.sub("".join(exp2)," #n ", linea)
        i += 1
    entra = r'([ ]*#e[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*)'
    salefin = r'([ ]*#s[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*|[ ]*\".+\"[ ]*|[ ]*-?[0-9]+\.[0-9]+[ ]*|[ ]*-?[0-9]+[ ]*)[ ]*'
    operfinsi = r'([ ]*#o[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*|[ ]*\".+\"[ ]*|[ ]*-?[0-9]+\.[0-9]+[ ]*|[ ]*-?[0-9]+[ ]*)([ ]+AN[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*|[ ]*\".+\"[ ]*|[ ]*-?[0-9]+\.[0-9]+[ ]*|[ ]*-?[0-9]+[ ]*)'
    unafinsi = r'([ ]*#n[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*|[ ]*\".+\"[ ]*|[ ]*-?[0-9]+\.[0-9]+[ ]*|[ ]*-?[0-9]+[ ]*)'
    decla2 = r'([ ]*#d[ ]+[A-Za-z][A-Za-z0-9_]*[ ]+)(#b[ ]+)([A-Za-z][A-Za-z0-9_]+[ ]*|\".+\"[ ]*|-?[0-9]+\.[0-9]+[ ]*|-?[0-9]+[ ]*)'
    decla3 = r'([ ]*#d[ ]+[A-Za-z][A-Za-z0-9_]*[ ]+)(#b[ ]*)?'
    asigna2 = r'([ ]*#a[ ]+)([A-Za-z][A-Za-z0-9_]*[ ]*|\".+\"[ ]*|-?[0-9]+\.[0-9]+[ ]*|-?[0-9]+[ ]*)'

    operfinsi = re.compile(operfinsi)
    unafinsi = re.compile(unafinsi)
    decla2 = re.compile(decla2)
    asigna2 = re.compile(asigna2)
    entra = re.compile(entra)
    salefin = re.compile(salefin)
    decla3 = re.compile(decla3)

    er3 = [entra,salefin,operfinsi,unafinsi,decla2,decla3,asigna2]
    j = 0
    lineacomp = ""
    condicion = True
    matches = []
    while(condicion):
        i = 0
        for exp in er3:
            matches.append(exp.findall(linea))
        for find in matches:
            for fi in find:
                k = 0
                if(i == 5):
                    casos = re.compile(r'([ ]*#d[ ]+[A-Za-z][A-Za-z0-9_]*[ ]+)(#b[ ]*)')
                    casos = casos.findall(linea)
                    for caso in casos:
                        if ("".join(caso) == "".join(fi)):
                            k = -1
                if (k != -1):
                    linea = re.sub("".join(fi)," 1 ",linea)
            i += 1
        lineacomp = linea
        if(lineacomp == linea):
            j += 1
        if(j > 12):
            condicion = False
    final = re.match(r'[ ]*1[ ]*', linea)
    if (final != None and final.group(0) == linea):
        return True
    return False

###
# Nombre: llew (la linea esta wena).
# Input lista, er, erdawn.
#   linea: Expresion a  verificar.
#   er: Lista de los compilado de las expresiones regulares generales.
#   erdawn: Lista de los compilados de las bases de todas las sentencias.
# Descripcion: Revisa si linea se encuentra  sintacticamente correctamente escrita
#              excepto el loop y el condicional.
# Return: True o False, si la sentencia esta bien o mal escrita respectivamente.
###
def llew(linea,er,erdawn):
    matches = []
    mal = re.compile(r'[ ]*[0-9][A-Za-z0-9_]+[ ]*')
    if (mal.match(str(linea)) != None):
        return False
    for expr in er:
        match = expr.match(linea)
        matches.append(match)
    if (matches[0] != None):
        return True
    elif(matches[1] != None and matches[2] != None and matches[1].group(0) == matches[2].group(0)):
        return True
    elif(matches[5] != None and matches[6] != None and matches[5].group(0) == matches[6].group(0)):
        return True
    elif(matches[3] != None and matches[4] != None and matches[3].group(0) == matches[4].group(0)):
        return True
    elif(matches[7] != None and matches[8] != None and matches[7].group(0) == matches[8].group(0)):
        return True
    elif(matches[13] != None and matches[14] != None and matches[13].group(0) == matches[14].group(0)):
        return True
    elif(matches[12] != None):
        return True
    elif(matches[10] != None):
        return llew(matches[10].group(7),er,erdawn)
    elif(matches[11] != None):
        return True
    elif(matches[9] != None):
        return True
    elif(matches[13] != None):
        return llew(matches[13].group(2),er,erdawn)
    elif(matches[1] != None):
        return obh(matches[1].group(0),er,erdawn)
    elif(matches[3] != None):
        return llew(matches[3].group(2),er,erdawn)
    elif(matches[5] != None):
        return llew(matches[5].group(5),er,erdawn)
    elif(matches[7] != None):
        return llew(matches[7].group(3),er,erdawn)
    else:
        return False

###
# Nombre: iolb (if o loop bueno).
# Input lista, er, erdawn.
#   linea: Expresion a  verificar.
#   er: Lista de los compilado de las expresiones regulares generales.
#   erdawn: Lista de los compilados de las bases de todas las sentencias.
# Descripcion: Revisa si linea, loop o condicional  se encuentra  sintacticamente correcta.
# Return: True o False, si la sentencia esta bien o mal escrita sintacticamente.
###
def iolb(line,er,erdawn):
    if(er[10].match(line) != None):
        return llew(er[10].match(line).group(7),er,erdawn)
    elif(er[11].match(line) != None and er[11].match(line).group(0) == line):
        return True
    elif(er[9].match(line) != None and er[9].match(line).group(0) == line):
        return True
    return False

###
# Nombre: iolb (la linea tiene hai o kthxbye como variable).
# Input lista.
#   linea: Expresion a  verificar.
# Descripcion: Revisa que todas la linea no tengan como variable "HAI"
#              o "KTHXBYE".
# Return: True o False, si la sentencia esta bien o mal escrita sintacticamente.
###
def llthokcv(line):
    openclose2 = re.compile(r'([ ]+HAI[ ]+|[ ]+KTHXBYE[ ]+)')
    openclose2 = openclose2.findall(line)
    for exp in openclose2:
        return False
    return True

###
# Nombre: dc (donde comienza).
# Input codigo.
#   codigo: Lista que contiene todas las sentencias.
# Descripcion: Encuentra donde inicia o termina el programa.
# Return: Retorna la posicion de inicio y termino del programa, en caso de no.
#         haber "HAI" o "KTHXBYE" retorna [-1.-1].
###
def dc(codigo):
    i = 0
    j = 0
    k = 0
    open = re.compile(r'[ ]*HAI[ ]*')
    close = re.compile(r'[ ]*KTHXBYE[ ]*')
    for line in codigo:
        if (open.match(line) != None):
            j = i
            k = 1
        elif (close.match(line) != None and k == 1):
            return [j,i]
        i += 1
    return [-1,-1]

###
# Nombre: visible.
# Input linea, erver
#   linea: Sentencia bien escrita.
#   erver: Palabras claves del programa.
# Descripcion: Muestra en pantalla la sentencia coloriada.
# Return: No poesee return.
###
def visible(linea,erver):
	matches = []
	for exp in erver:
		matches.append(exp.findall(linea))
	i = 0
	matches.append(re.findall(r'[ ]+AN',linea))
	for find in matches:
		for exp in find:
			if(i == 0):
				linea = re.sub(erver[0],"\u001b[36m"+"".join(exp)+"\u001b[0m",linea)
			elif(i == 1):
				linea = re.sub("".join(exp),"\u001b[35m"+exp[0]+"\u001b[37m"+exp[1]+"\u001b[35m"+exp[2]+exp[3]+"\u001b[37m"+exp[4]+"\u001b[35m"+exp[5]+"\u001b[0m"+exp[6],linea)
			elif(i == 2):
				linea = re.sub("".join(exp),"\u001b[35m"+exp[0]+"\u001b[0m"+exp[1],linea)
			elif(i == 3):
				linea = re.sub("".join(exp),"\u001b[31m"+exp[0]+"\u001b[0m"+exp[1],linea)
			elif(i == 4):
				linea = re.sub("".join(exp),"\u001b[34m"+exp[0]+"\u001b[0m"+exp[1],linea)
			elif(i == 5):
				linea = re.sub(exp[0]+exp[1]+exp[3]+exp[4],"\u001b[33m"+exp[0]+"\u001b[37m"+exp[1]+"\u001b[33m"+exp[3]+"\u001b[0m"+exp[4],linea)
			elif(i == 6):
				linea = re.sub("".join(exp),exp[0]+"\u001b[31m"+exp[1]+"\u001b[0m"+exp[2]+exp[3],linea)
			elif(i == 7):
				linea = re.sub(exp[0],"\u001b[34m"+exp[0]+"\u001b[0m",linea)
				linea = re.sub(exp[1]+exp[2]+exp[3], exp[1]+"\u001b[34m"+exp[2]+"\u001b[0m"+exp[3],linea)
			elif(i == 8):
				linea = re.sub(exp[0]+exp[1],"\u001b[31m"+exp[0]+"\u001b[0m"+exp[1],linea)
			elif(i == 9):
				linea = re.sub("".join(exp),"\u001b[32m"+exp+"\u001b[0m",linea)
			elif(i == 10):
				linea = re.sub(exp,"\u001b[34m"+exp+"\u001b[0m",linea)
			elif(i == 11):
				linea = re.sub("".join(exp),exp[0]+"\u001b[31m"+exp[1]+"\u001b[0m",linea)
			elif(i == 12):
				linea = re.sub("".join(exp),"\u001b[33m"+exp+"\u001b[0m",linea)
			elif(i == 13):
				linea = re.sub("".join(exp),"\u001b[31m"+exp+"\u001b[0m",linea)
			elif(i == 14):
				linea = re.sub("".join(exp),"\u001b[31m"+exp+"\u001b[0m",linea)
			elif(i == 15):
				linea = re.sub("".join(exp),"\u001b[34m"+exp+"\u001b[0m",linea)
			elif(i == 16):
				linea = re.sub("".join(exp),"\u001b[33m"+exp+"\u001b[0m",linea)
			elif(i == 17):
				linea = re.sub("".join(exp),"\u001b[34m"+exp+"\u001b[0m",linea)
		i += 1
	print(linea.strip())



###
# Nombre: main.
# Descripcion: Se declara todas las expresiones regulares y sus respectivas compilaciones,
#              tambien se llama a las diferentes funciones para el analisis  de la sintaxis
#              del codigo de "LOLCODE", y por ultimo se llama a la funcion que imprime
#              en pantalla los resultados del analisis,
###

orly = r'([ ]*O[ ]+RLY\?[ ]*)'
yarly = r'([ ]*YA[ ]+RLY[ ]*)'
noWAI = r'([ ]*NO[ ]+WAI[ ]*)'
oic = r'([ ]*OIC[ ]*)'
openclose = r'([ ]*HAI[ ]*|[ ]*KTHXBYE[ ]*)'
entra = r'([ ]*GIMMEH[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*)'
salegen = r'([ ]*VISIBLE[ ]+)(.+)'
salefin = r'([ ]*VISIBLE[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*|[ ]*\".+\"[ ]*|[ ]*-?[0-9]+\.[0-9]+[ ]*|[ ]*-?[0-9]+[ ]*)'
oper = r'([ ]*SUM[ ]+OF[ ]+|[ ]*DIFF[ ]+OF[ ]+|[ ]*PRODUKT[ ]+OF[ ]+|[ ]*QUOSHUNT[ ]+OF[ ]+|[ ]*MOD[ ]+OF[ ]+|[ ]*BIGGR[ ]+OF[ ]+|[ ]*SMALLR[ ]+OF[ ]+|[ ]*BOTH[ ]+OF[ ]+|[ ]*EITHER[ ]+OF[ ]+|[ ]*BOTH[ ]+SAEM[ ]+|[ ]*DIFFRINT[ ]+)'
opergen = oper+r'((.+)([ ]+AN[ ]+)(.+)[ ]*)'
operfinsi = oper+r'([ ]*[A-Za-z][A-Za-z0-9_]*|[ ]*\".+\"|[ ]*-?[0-9]+\.[0-9]+|[ ]*-?[0-9]+)([ ]+AN[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*|[ ]*\".+\"|[ ]*-?[0-9]+\.[0-9]+|[ ]*-?[0-9]+)[ ]*'
una = r'([ ]*NOT[ ]+)'
unagen = una+r'(.+)'
unafinsi = una+r'([ ]*[A-Za-z][A-Za-z0-9_]*|[ ]*\".+\"|[ ]*-?[0-9]+\.[0-9]+|[ ]*-?[0-9]+)[ ]*'
declagen = r'([ ]*I[ ]+HAS[ ]+A[ ]+)([A-Za-z][A-Za-z0-9_]*)(([ ]+ITZ[ ]+)(.+))?'
decla2 = r'([ ]*I[ ]+HAS[ ]+A[ ]+)([A-Za-z][A-Za-z0-9_]*)(([ ]+ITZ[ ]+)([A-Za-z][A-Za-z0-9_]+|\".+\"|-?[0-9]+\.[0-9]+|-?[0-9]+)[ ]*)?'
asigna = r'([ ]*[A-Za-z][A-Za-z0-9_]*)([ ]+R[ ]+)(.+)'
asigna2 = r'([ ]*[A-Za-z][A-Za-z0-9_]*)([ ]+R)([ ]+)([A-Za-z][A-Za-z0-9_]*|\".+\"|-?[0-9]+\.[0-9]+|-?[0-9]+)[ ]*'
condi = r'([ ]*O[ ]+RLY\?[ ]*|[ ]*YA[ ]+RLY[ ]*|[ ]*NO[ ]+WAI[ ]*|[ ]*OIC[ ]*)'
loop = r'([ ]*IM[ ]+IN[ ]+YR[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*)([ ]+UPPIN[ ]+|[ ]+NERFIN[ ]+)(YR[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*)([ ]+TIL|[ ]+WILE)([ ]+.+)'
loopclose = r'([ ]*IM[ ]+OUTTA[ ]+YR[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*[ ]*)'
unabase = r'([ ]*NOT[ ]+)'
unafinsi2 = r'(NOT[ ]+)([ ]*[A-Za-z][A-Za-z0-9_]*|[ ]*\".+\"|[ ]*-?[0-9]+\.[0-9]+|[ ]*-?[0-9]+)[ ]*'
operbase = r'([ ]*SUM[ ]+OF[ ]+|[ ]*DIFF[ ]+OF[ ]+|[ ]*PRODUKT[ ]+OF[ ]+|[ ]*QUOSHUNT[ ]+OF[ ]+|[ ]*MOD[ ]+OF[ ]+|[ ]*BIGGR[ ]+OF[ ]+|[ ]*SMALLR[ ]+OF[ ]+|[ ]*BOTH[ ]+OF[ ]+|[ ]*EITHER[ ]+OF[ ]+|[ ]*BOTH[ ]+SAEM[ ]+|[ ]*DIFFRINT[ ]+)'
asignabase = r'([ ]*[A-Za-z][A-Za-z0-9_]*)([ ]+R[ ]+)'
asignabase2 = r'([ ]+[A-Za-z][A-Za-z0-9_]*)([ ]+R)'
declabase = r'([ ]*I[ ]+HAS[ ]+A[ ]+)'
declabase2 = r'(I[ ]+HAS[ ]+A[ ]+)'
ohboi = r'([ ]+ITZ[ ]+)'
ohboi2 = r'([ ]+ITZ)'
salebase = r'([ ]*VISIBLE[ ]+)'
entrabase = r'([ ]*GIMMEH[ ]+)'
salebase2 = r'(VISIBLE)'
entrabase2 = r'(GIMMEH)'
operbase2 = r'(SUM[ ]+OF|DIFF[ ]+OF|PRODUKT[ ]+OF|QUOSHUNT[ ]+OF|MOD[ ]+OF|BIGGR[ ]+OF|SMALLR[ ]+OF|BOTH[ ]+OF|EITHER[ ]+OF|BOTH[ ]+SAEM|DIFFRINT)'

openclose = re.compile(openclose)
opergen = re.compile(opergen)
operfinsi = re.compile(operfinsi)
unagen = re.compile(unagen)
loop = re.compile(loop)
unafinsi = re.compile(unafinsi)
unafinsi2 = re.compile(unafinsi2)
declagen = re.compile(declagen)
decla2 = re.compile(decla2)
asigna = re.compile(asigna)
asigna2 = re.compile(asigna2)
condi = re.compile(condi)
loopclose = re.compile(loopclose)
entra = re.compile(entra)
salegen = re.compile(salegen)
salefin = re.compile(salefin)
unabase = re.compile(unabase)
operbase = re.compile(operbase)
operbase2 = re.compile(operbase2)
asignabase = re.compile(asignabase)
asignabase2 = re.compile(asignabase2)
declabase = re.compile(declabase)
ohboi = re.compile(ohboi)
ohboi2 = re.compile(ohboi2)
salebase = re.compile(salebase)
entrabase = re.compile(entrabase)
declabase2 = re.compile(declabase2)
salebase2 = re.compile(salebase2)
entrabase2 = re.compile(entrabase2)
orly = re.compile(orly)
yarly = re.compile(yarly)
noWAI = re.compile(noWAI)
oic = re.compile(oic)

er = [openclose,opergen,operfinsi,unagen,unafinsi,declagen,decla2,asigna,asigna2,condi,loop,loopclose,entra,salegen,salefin]
erdawn = [operbase,asignabase,declabase,ohboi,salebase,entrabase,unabase]
erver = [condi,loop,loopclose,entra,unafinsi2,decla2,asigna2,operfinsi,salefin,openclose,operbase2,asignabase2,declabase2,salebase2,entrabase2,unabase,ohboi]
er2 = [orly,yarly,noWAI,oic]

file_name = input("nombre del archivo: ")
file = open(file_name)
codigo = []
for line in file:
    codigo.append(line)
file.close()
file = open(file_name)
codigo_new = dc(codigo)
codigofin = []
codigoifloop = []
i = 0
if (codigo_new[0] == -1):
    for line in file:
        print(u'\u001b[41m'+line.strip()+u'\u001b[0m')
else:
    j = 0
    for line in file:
        if(codigo_new[0] <= j and codigo_new[1] >= j):
            codigoifloop.append(line)
        j += 1
    for line in codigo:
        if(codigo_new[0] > i or codigo_new[1] < i):
            codigofin.append([line,-1])
        elif(codigo_new[0] == i or codigo_new[1] == i):
            codigofin.append([line,1])
        else:
            valides1 = llew(line,er,erdawn)
            valides2 = llthokcv(line)
            valides3 = iolb(line,er,erdawn)
            if (valides3 == True and valides2 == True):
                codigofin.append([line,1])
            elif(valides1 == True and valides2 == True):
                codigofin.append([line,1])
            else:
                codigofin.append([line,-1])
        i += 1
    error = set()
    error = ifdetector(codigoifloop, er, er2, error)
    if (error != False):
        for erro in error:
            codigofin[erro[1]+codigo_new[0]][1] = -1
    error = set()
    error = ciclos(codigoifloop,er,error)
    if (len(error) != 0):
        for erro in error:
            codigofin[erro[1]+codigo_new[0]][1] = -1
for line in codigofin:
    if (line[1] == -1):
        print("\u001b[41m"+line[0].strip()+"\u001b[0m")
    else:
        visible(line[0],erver)
file.close()
