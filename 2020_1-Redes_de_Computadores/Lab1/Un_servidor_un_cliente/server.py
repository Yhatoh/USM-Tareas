import socket as calcetin

#49152-65535
puertoServidor = 50366


    socketServidor = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_STREAM)
    socketServidor.bind(("", puertoServidor))
    urlsCache = []
    headers = []
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

    while 1:
        socketServidor.listen(1)
        print("Esperando que un cliente me agarre de la manito!")
        socketCliente, dnte = socketServidor.accept()
        
        while 1:
            
            
            target = socketCliente.recv(2048).decode()
            if target.upper() == "TERMINATE":
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
                if len(urlsCache) > 4:
                    urlsCache = urlsCache[1:]
                    headers = headers[1:]
            else:
                head = headers[urlsCache.index(target)] 
                headers.remove(headers[urlsCache.index(target)])
                urlsCache.remove(target)

            #print(urlsCache)
            z = "49152"
            socketCliente.send(z.encode())
            

            socketTransfer = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_DGRAM)
            socketTransfer.bind(("", int(z)))
            mensaje, direccionCliente = socketTransfer.recvfrom(4096)
            
            urlsCache.append(target)
            headers.append(head)
            


            oki = mensaje.decode()
            print(oki)
            if oki.upper() == "OK":
                socketTransfer.sendto(head.encode(), direccionCliente)
            socketTransfer.close()
        socketCliente.close()
        cache = open("cache.txt","w")
        i = len(urlsCache) - 1
        while i >= 0:
            cache.write(urlsCache[i] + "\n")
            cache.write("---------------------\n")
            cache.write(headers[i])
            cache.write("---------------------\n")
            i -= 1
        cache.close()    
    socketServidor.close()

