package com.TOMC.demo.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.GenericGenerator;

import jdk.nashorn.internal.objects.annotations.Getter;

@Table(name="PersonalSillon")
@Entity
public class PersonalSillon{
	@Id
	@GenericGenerator(name="incrementcliente",strategy="increment")
	@GeneratedValue(generator="incrementcliente")
	@Column(name="id")
	long id;

    @Column(name="rut")
    String rut;

    @Column(name="id_sillon")
    long idsillon;

	public PersonalSillon(){

	}

	public PersonalSillon(long id, String rut, long idsillon){
		this.id = id;
		this.rut = rut;
		this.idsillon = idsillon;
	}

	public long getId(){
		return this.id;
	}

	public void setId(long id){
		this.id = id;
	}

	public String getRut(){
		return this.rut;
	}

	public void setRut(String rut){
		this.rut = rut;
	}

	public long getIdsillon(){
		return this.idsillon;
	}

	public void setIdsillon(long idsillon){
		this.idsillon = idsillon;
	}
}
