import socket as calcetin
from _thread import *
import threading

#49152-65535
puertoServidor = 50366

#urlsCache y headers son dos listas que nos servirán para manejar el cache
urlsCache = []
headers = []
#este candado nos sirve para cuando pongan terminate no haya problemas en
#escribir en cache.txt
candado = threading.Lock()

def threadSiThread(c, num):
    while 1:
        global urlsCache
        global headers
        #Aqui esperamo que el cliente nos de consulta
        target = c.recv(2048).decode()
        print("Cliente numero",num,"me pide consultar por",target)
        if target.upper() == "TERMINATE":
            print("Gracias por preferirnos cliente numero", num)
            break
        if target not in urlsCache:
            #Peticion a la pagina web  
            sockete = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_STREAM)
            sockete.connect((target.lower(), 80))
            request = "GET / HTTP/1.1\r\nHost:"+ target +"\r\n\r\n"
            sockete.send(request.encode())
            respuesta = sockete.recv(4096)
            sockete.close()
            respuestaDecodificada = respuesta.decode(encoding="cp437")
            #--------------------------
            #Cortamos el header
            i = 0
            head = ""
            for x in respuestaDecodificada:
                head = head+x
                if x == '\r' and i == 0:
                    i += 1
                elif x == '\n' and i == 1:
                    i += 1
                elif x == '\r' and i == 2:
                    i += 1
                elif x == '\n' and i == 3:
                    break
                else:
                    i = 0
            #--------------------------
            #Si es que hay 5 elementos en urlsCache sacamos el menos usado
            #porque así funciona LRU
            if len(urlsCache) > 4:
                urlsCache = urlsCache[1:]
                headers = headers[1:]
        else:
            #Esto nos servira para poder mover un link de la cache a la 
            #primera posicion ya que asi es como funciona LRU
            head = headers[urlsCache.index(target)] 
            headers.remove(headers[urlsCache.index(target)])
            urlsCache.remove(target)
        #Mandamo puerto a cliente esperamos respuesta
        z = "49152"
        c.send(z.encode())
        socketTransfer = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_DGRAM)
        socketTransfer.bind(("", int(z)))
        mensaje, direccionCliente = socketTransfer.recvfrom(4096)
        #--------------------------
        #Agregamos la ultima url usada para la lista
        urlsCache.append(target)
        headers.append(head)
        #--------------------------
        #recibimos oki y mandamos la info
        oki = mensaje.decode()
        print(oki)
        if oki.upper() == "OK":
            socketTransfer.sendto(head.encode(), direccionCliente)
        socketTransfer.close()
        #--------------------------
    c.close()
    #Si es que el cliente pone terminate entonces la thread que llega
    #primero toma el candado y guarda el cache en un txt para luego
    #soltar el candado 
    candado.acquire()
    cache = open("cache.txt","w")
    i = len(urlsCache) - 1
    while i >= 0:
        cache.write(urlsCache[i] + "\n")
        cache.write("---------------------\n")
        cache.write(headers[i])
        cache.write("---------------------\n")
        i -= 1
    cache.close()  
    candado.release() 
    #--------------------------

socketServidor = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_STREAM)
socketServidor.bind(("", puertoServidor))

#revisamos la cache guardad en txt y guardamos en las listas
try:
    cache = open("cache.txt","r")
    cache = cache.read().split("---------------------\n")
    i = 0
    for x in cache:
        if x == "":
            break
        if i % 2 == 0:
            x = x.strip()
            urlsCache.append(x)
        else:
            headers.append(x)
        i += 1
    cache.close()
    urlsCache.reverse()
    headers.reverse()
except Exception as e:
    pass
#--------------------------
count = 0
while 1:
    #Esperamos a un cliente y le asignamos un numero para poder saber
    #que consulta es de quien.
    socketServidor.listen(1)
    print("Esperando cliente.")
    socketCliente, dnte = socketServidor.accept()
    count += 1
    print("Cliente numero",count,"se ha conectado al servidor.")
    start_new_thread(threadSiThread, (socketCliente, count))
    #--------------------------
socketServidor.close()

