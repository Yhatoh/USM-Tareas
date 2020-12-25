package com.TOMC.demo.service;

import java.util.List;

import com.TOMC.demo.entity.PersonalSillon;
import com.TOMC.demo.repository.PersonalSillonRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

@Service("servicepersonalsillon")
public class PersonalSillonService{
    @Autowired
    @Qualifier("repositorypersonalsillon")
    private PersonalSillonRepository repository;

    public boolean addPersonalSillon(PersonalSillon personalsillon){
        try{
            repository.save(personalsillon);
            return true;
        } catch (Exception e){
            return false;
        }
    }

    public boolean updatePersonalSillon(PersonalSillon personalsillon){
        try{
            repository.save(personalsillon);
            return true;
        } catch (Exception e){
            return false;
        }
    }

    public boolean deletePersonalSillon(long id){
        try{
            PersonalSillon personalsillon = repository.findById(id);
            repository.delete(personalsillon);
            return true;
        } catch (Exception e){
            return false;
        }
    }

    public PersonalSillon getById(long id){
        return repository.findById(id);
    }

    public List<PersonalSillon> getAll(){
        return repository.findAll();
    }
}