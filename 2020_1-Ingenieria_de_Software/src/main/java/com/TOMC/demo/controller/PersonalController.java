package com.TOMC.demo.repository;

import java.util.List;

import javax.validation.Valid;

import com.TOMC.demo.entity.Personal;
import com.TOMC.demo.service.PersonalService;

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
@RequestMapping("/v1/personal")
public class PersonalController{
    @Autowired
    @Qualifier("servicepersonal")
    PersonalService service;

    @GetMapping("{rut}")
    public Personal getPersonal(@PathVariable("rut") String rut){
        return service.getByRut(rut);
    }

    @PostMapping("")
    public String addPersonal(@RequestBody @Valid Personal personal){
        if(faltanCampos(personal))
        {
            if(personal.getTipoPersonal() < 0 || personal.getTipoPersonal() > 4 || !verificarRut(personal.getRut()))
            {
                return "Personal incorrecto";
            }

            if(service.addPersonal(personal))
            {
                return "Personal guardado";
            }
            else
            {
                return "Personal incorrecto";
            }
        }
        else
        {
            return "Persona con campos faltantes";
        }
    }

    @PutMapping("")
    public String updatePersonal(@RequestBody @Valid Personal personal){
        if(faltanCampos(personal))
        {
            if(personal.getTipoPersonal() < 0 || personal.getTipoPersonal() > 4 || !verificarRut(personal.getRut()))
            {
                return "Personal incorrecto";
            }

            if(service.updatePersonal(personal))
            {
                return "Personal actualizado";
            }
            else
            {
                return "Personal incorrecto";
            }
        }
        else
        {
            return "Personal con campos vacios";
        }
    }

    @DeleteMapping("/{rut}")
    public String deletePersonal(@PathVariable("rut") String rut){
        if(service.deletePersonal(rut))
        {
            return "Personal eliminado";
        }
        else
        {
            return "Personal no existente";
        }
    }

    @GetMapping("/getAll")
    public List<Personal> getAllPersonal(){
        return service.getAll();
    }

    @GetMapping("/tipo/{tipoPersonal}")
    public List<Personal> getByTipoPersonal(@PathVariable("tipoPersonal") int tipoPersonal){
        return service.getByType(tipoPersonal);
    }
    
    @GetMapping("/disp/{disponibilidad}")
    public List<Personal> getByDisponibilidad(@PathVariable("disponibilidad") Boolean disponibilidad){
        return service.getByDisponibilidad(disponibilidad);
    }

    public boolean faltanCampos(Personal personal){
        return personal.getNombre() != null && personal.getRut() != null && personal.getApellido() != null && personal.getNumero() != null && personal.getMail() != null && personal.getTipoPersonal() != 0 && personal.getDisponibilidad() != null && personal.getEspecializacion() != null && personal.getProfesion() != null;
    }

    public boolean verificarRut(String rut){
        int len = rut.length();
        int num = 2;
        int suma = 0;
        for(int i = len-3; i >= 0; --i)
        {
            suma = suma + num * (rut.charAt(i) - '0');
            num++;
            if(num > 7) num = 2;
        }
        suma = suma % 11;
        suma = 11 - suma;
        if(suma == 11) suma = 0;
        if(suma == 10) suma = 'k' - '0';

        if(suma != (rut.charAt(len - 1) - '0'))
        {
            return false;
        }
        else
        {
            return true;
        }
    }
}
