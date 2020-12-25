import socket as calcetin

direccionServidor = "localhost"
#49152-65535
puertoServidor = 50366

socketCliente = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_STREAM)

#se realiza el handshaking
socketCliente.connect((direccionServidor, puertoServidor))

while 1:
	aEnviar = input("url: ")
	#envia el url, si escribe terminate entonces se detiene.
	if aEnviar.upper() == "TERMINATE":
		print("No quiero preguntar mas >:c")
		socketCliente.send(aEnviar.encode())
		break
	socketCliente.send(aEnviar.encode())
	#--------------------------
	#Aqui recibo puerto
	respuesta = socketCliente.recv(4096).decode()
	#Abro upd cliente
	udpCliente = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_DGRAM)
	#Aqui mando OK para recibir la info, la recivo y lo escribo en URL.txt.
	ok = "OK"
	udpCliente.sendto(ok.encode(),(direccionServidor, int(respuesta)))
	head = udpCliente.recv(4096).decode()
	print("Recibi la info! :D")
	udpCliente.close()
	headerTxt = open(aEnviar+ ".txt", "w")
	headerTxt.write(head)
	headerTxt.close()
	#--------------------------
socketCliente.close()
