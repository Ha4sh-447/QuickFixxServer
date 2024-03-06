package com.quickfixxMicroservices.PlumberService.repository;

import com.quickfixxMicroservices.PlumberService.model.Plumber;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface PlumberRepo extends JpaRepository<Plumber, Long> {
    List<Plumber> findAllByName(String name);
}
