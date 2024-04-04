package com.quickfixxMIcroservices.carpenterService.repository;

import com.quickfixxMIcroservices.carpenterService.model.Carpenter;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.List;

public interface CarpenterRepo extends JpaRepository<Carpenter, Long> {
    List<Carpenter> findAllByName(String name);
    List<Carpenter> findByField(String field);
}
