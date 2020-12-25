from django import forms
from .models import *

class CartaForm(forms.ModelForm):
	class Meta:
		model = Equipo
		fields = ('carta1','carta2','carta3','carta4','carta5','jugador')
	def __init__(self, *args, **kwargs):
		super(CartaForm, self).__init__(*args, **kwargs)
		self.fields['carta1'].widget.attrs.update({'class':'form-control','placeholder':'Ingrese Cantidad de Cartas'})
		self.fields['carta2'].widget.attrs.update({'class':'form-control','placeholder':'Ingrese Cantidad de Cartas'})
		self.fields['carta3'].widget.attrs.update({'class':'form-control','placeholder':'Ingrese Cantidad de Cartas'})
		self.fields['carta4'].widget.attrs.update({'class':'form-control','placeholder':'Ingrese Cantidad de Cartas'})
		self.fields['carta5'].widget.attrs.update({'class':'form-control','placeholder':'Ingrese Cantidad de Cartas'})
		self.fields['jugador'].widget.attrs.update({'class':'form-control','placeholder':'Ingrese Cantidad de Cartas'})