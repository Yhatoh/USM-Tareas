package com.TOMC.demo.repository;

import java.io.Serializable;
import java.util.List;

import com.TOMC.demo.entity.PersonalSillon;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("repositorypersonalsillon")
public interface PersonalSillonRepository extends JpaRepository<PersonalSillon, Serializable>{
    public abstract PersonalSillon findById(long id);
    public abstract List<PersonalSillon> findAll();

}