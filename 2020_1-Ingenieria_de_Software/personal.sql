CREATE TABLE personal (
   id_personal SERIAL PRIMARY KEY,
   rut VARCHAR(9) NOT NULL,
   nombre VARCHAR(50) NOT NULL,
   apellido VARCHAR(50) NOT NULL,
   numero VARCHAR(50) NOT NULL,
   mail VARCHAR(100) NOT NULL,
   tipoPersonal INT NOT NULL,
   disponibilidad BOOLEAN NOT NULL
);