from django.db import models


class Carta(models.Model):
	id=models.AutoField(primary_key=True)
	nombre=models.CharField(max_length=100, verbose_name="Nombre", blank=False, null=True)
	rareza=models.IntegerField(verbose_name="Rareza",blank=True,null=True)
	ataque=models.IntegerField(verbose_name="Ataque",blank=True,null=True)
	defensa=models.IntegerField(verbose_name="Defensa",blank=True,null=True)
	imagen=models.CharField(max_length=200, verbose_name="Imagen", blank=False, null=True)
	lore=models.TextField(verbose_name="Lore", blank=False, null=True)

	def __str__(self):
		return u'A: %d D: %d | %s' % (self.ataque,self.defensa,self.nombre,)
		
class Mazo(models.Model):
	carta=models.ForeignKey(Carta,on_delete=models.CASCADE,blank=True,null=True)
	jugador=models.CharField(max_length=20, verbose_name="Jugador", blank=False, null=True)
	cantidad=models.IntegerField(verbose_name="Cantidad",blank=True,null=True)
	def __str__(self):
		return u'%s | %s' % (self.jugador,self.carta.nombre)

class Equipo(models.Model):
	carta1=models.IntegerField(blank=True,null=True)
	carta2=models.IntegerField(blank=True,null=True)
	carta3=models.IntegerField(blank=True,null=True)
	carta4=models.IntegerField(blank=True,null=True)
	carta5=models.IntegerField(blank=True,null=True)
	jugador=models.CharField(max_length=20, verbose_name="Jugador", blank=True, null=True)
	def __str__(self):
		return u'%s' % (self.jugador)

class Fotito(models.Model):
	foto=models.TextField(verbose_name="Foto", blank=False, null=True)
	jugador=models.CharField(primary_key=True,max_length=20, verbose_name="Jugador")
	def __str__(self):
		return u'%s' % (self.jugador)