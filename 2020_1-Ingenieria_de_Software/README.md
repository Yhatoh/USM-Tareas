# Tomc's

Proyecto realizado primer semestre 2020 para la Universidad Técnica Federico Santa María. Corresponde a una API realizada en springboot.

# API

### Personal

A través de esta API se manejan datos de personal que siguen la siguiente estructura:

```
{
  rut: string,
  nombre: string,
  apellido: string,
  numero: string,
  mail: string,
  profesion: string,
  especializacion: string,
  tipoPersonal: int,
  disponibilidad: bool
}
```

Donde:

* rut es todo con guión y dígito verificador
* nombre, apellido, profesion, especializacion son string con información del personal
* numero es su numero de telefono
* mail es su correo
* tipoPersonal es un entero donde 1: Pabellon, 2: Quimioterapia, 3: Recuperación y 4: Medico
* disponibilidad es un booleano de si esta disponible o no

### POST

* Link: https://tomcs.herokuapp.com/v1/personal/
* VerboHTTP: POST
* Acción: Almacenar un personal
* Cuerpo: personal
* Respuesta: string, que dice si se guardo o no

### PUT

* Link: https://tomcs.herokuapp.com/v1/personal/
* VerboHTTP: PUT
* Acción: Actualizar un personal
* Cuerpo: personal
* Respuesta: string, que dice si se actualizo o no

### GET

#### Por rut

* Link: https://tomcs.herokuapp.com/v1/personal/{rut}
* VerboHTTP: GET
* Acción: Obtener un personal por un rut
* Argumento: rut
* Respuesta: personal

#### Por tipoPersonal

* Link: https://tomcs.herokuapp.com/v1/personal/{tipoPersonal}
* VerboHTTP: GET
* Acción: Obtener todo personal de un tipo
* Argumento: tipoPersonal
* Respuesta: lista de personal del mismo tipo

#### Por Disponibilidad

* Link: https://tomcs.herokuapp.com/v1/personal/{Disponibilidad}
* VerboHTTP: GET
* Acción: Obtener todo personal disponible o no disponible
* Argumento: Disponibilidad
* Respuesta: lista de personal disponible o no disponible

#### Todo personal

* Link: https://tomcs.herokuapp.com/v1/personal/getAll
* VerboHTTP: GET
* Acción: Obtener todo el personal
* Respuesta: lista de personal

### DELETE

* Link: https://tomcs.herokuapp.com/v1/personal/{rut}
* VerboHTTP: DELETE
* Acción: Borrar un personal por rut
* Argumento: rut
* Respuesta: string, que dice sise borro o no

# Proyecto Ingeniería de Software

### Prerequisitos

Para poder utilizar la aplicación necesitan tener instalado lo siguiente:

* [Maven](https://maven.apache.org/)  
* [Java](https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) 
* [Postgres](https://www.postgresql.org/download/)

### Preparativos
Antes de montar el proyecto se debe crear una base de datos en postgres con el nombre personal. Una vez creada la base de datos se debe ejecutar el script dentro de la carpeta Api el cual se llama create.sql. Pare ejecutarlo se debe copiar su contenido, y pegarlo en la opción 'Create Script' de PgAdmin4. Finalmente, en la carpeta resources dentro de la carpeta API, se debe modificar el valor de la contraseña de la base de datos en el archivo 'applicaction.properties'.

### Montar
Para poder montar el proyecto primero deben realizar lo siguiente:
* Abrir la consola de comandos y dirigirse a la carpeta API, luego ejecutar:
```
mvn install
mvn spring-boot:run
```
Nota: De esta forma se montará la api en el puerto 8000 por defecto, en caso de tenerlo ocupado puede modificar su valor en 'application.properties', donde modificó la contraseña de la base de datos.

## Autores

* **Sebastián Campos** 
* **Gabriel Carmona**
* **Jorge Ludueña**



