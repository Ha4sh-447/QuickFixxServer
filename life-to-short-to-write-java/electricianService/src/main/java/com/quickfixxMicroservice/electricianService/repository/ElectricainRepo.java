package com.quickfixxMicroservice.electricianService.repository;

import com.quickfixxMicroservice.electricianService.model.Electrician;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface ElectricainRepo extends JpaRepository<Electrician, Long> {
    List<Electrician> findByLocation(String location);

//    Electrician findDataById(Long id);
    List<Electrician> findByName(String name);
}
