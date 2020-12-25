import socket as calcetin

direccionServidor = "localhost"
#49152-65535
puertoServidor = 50366

socketCliente = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_STREAM)

#si handshaking si
socketCliente.connect((direccionServidor, puertoServidor))

while 1:
	aEnviar = input("url: ")

	if aEnviar.upper() == "TERMINATE":
		print("No quiero preguntar mas >:c")
		socketCliente.send(aEnviar.encode())
		break
	socketCliente.send(aEnviar.encode())
	#Aqui recibo puerto
	respuesta = socketCliente.recv(4096).decode()
	#Abro upd cliente
	udpCliente = calcetin.socket(calcetin.AF_INET, calcetin.SOCK_DGRAM)

	ok = "OK"
	udpCliente.sendto(ok.encode(),(direccionServidor, int(respuesta)))
	head = udpCliente.recv(4096).decode()
	udpCliente.close()

	headerTxt = open(aEnviar+ ".txt", "w")
	headerTxt.write(head)
	headerTxt.close()
socketCliente.close()
