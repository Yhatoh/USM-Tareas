from django.conf.urls import url, include
from .views import *

urlpatterns = [
    url(r"^$",inicio),
    url(r"^home/$",home),
    url(r"^jugar/$",jugar),
    url(r"^mazo/$",mazo),
    url(r"^perfil/$",perfil),
    url(r"^info/$",info),
    url(r"^sobre/$",sobre),
    url(r"^equipo/$",equipo)
     
]
