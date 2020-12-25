package com.TOMC.demo.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.GenericGenerator;

import jdk.nashorn.internal.objects.annotations.Getter;

@Table(name="Personal")
@Entity
public class Personal{
	@Id
    @Column(name="rut")
    String rut;

    @Column(name="nombre")
    String nombre;

    @Column(name="apellido")
    String apellido;

    @Column(name="numero")
    String numero;

    @Column(name="mail")
    String mail;

	@Column(name="profesion")
    String profesion;

	@Column(name="especializacion")
    String especializacion;

    @Column(name="tipo_personal")
    int tipoPersonal;

    @Column(name="disponibilidad")
    Boolean disponibilidad;

	public Personal(){

	}

	public Personal(String rut, String nombre, String apellido, String numero, String mail, String profesion, String especializacion, int tipoPersonal, Boolean disponibilidad){
		
	}

	public String getRut(){
		return this.rut;
	}

	public void setRut(String rut){
		this.rut = rut;
	}

	public String getNombre(){
		return this.nombre;
	}

	public void setNombre(String nombre){
		this.nombre = nombre;
	}

	public String getApellido(){
		return this.apellido;
	}

	public void setApellido(String apellido){
		this.apellido = apellido;
	}

	public String getNumero(){
		return this.numero;
	}

	public void setNumero(String numero){
		this.numero = numero;
	}
	
	public String getMail(){
		return this.mail;
	}

	public void setMail(String mail){
		this.mail = mail;
	}

	public int getTipoPersonal(){
		return this.tipoPersonal;
	}

	public void setTipoPersonal(int tipoPersonal){
		this.tipoPersonal = tipoPersonal;
	}

	public Boolean getDisponibilidad(){
		return this.disponibilidad;
	}

	public void setDisponibilidad(Boolean disponibilidad){
		this.disponibilidad = disponibilidad;
	}

	public String getEspecializacion() {
		return especializacion;
	}

	public String getProfesion() {
		return profesion;
	}

	public void setEspecializacion(String especializacion) {
		this.especializacion = especializacion;
	}

	public void setProfesion(String profesion) {
		this.profesion = profesion;
	}
}
