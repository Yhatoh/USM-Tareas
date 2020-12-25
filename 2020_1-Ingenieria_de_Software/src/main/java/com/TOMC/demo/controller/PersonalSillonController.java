package com.TOMC.demo.controller;

import java.util.List;

import javax.validation.Valid;

import com.TOMC.demo.entity.PersonalSillon;
import com.TOMC.demo.service.PersonalSillonService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/v1/personalsillon")
public class PersonalSillonController{
    @Autowired
    @Qualifier("servicepersonalsillon")
    PersonalSillonService servicio;
    
    @GetMapping("/{id}")
    public PersonalSillon obtenerPersonalSillon(@PathVariable("id") long id){
        return servicio.getById(id);
    }

    @PostMapping("")
    public boolean agregarPersonalSillon(@RequestBody @Valid PersonalSillon personalsillon){
        return servicio.addPersonalSillon(personalsillon);
    }

    @PutMapping("")
    public boolean actualizarPersonalSillon(@RequestBody @Valid PersonalSillon personalsillon){
        return servicio.updatePersonalSillon(personalsillon);
    }

    @DeleteMapping("/{id}")
    public boolean borrarPersonalSillon(@PathVariable("id") long id){
        return servicio.deletePersonalSillon(id);
    }

    @GetMapping("/getAll")
    public List<PersonalSillon> getAll(){
        return servicio.getAll();
    }
}
