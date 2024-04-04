package com.quickfixxMicroservice.electricianService.repository;

import com.quickfixxMicroservice.electricianService.model.ElectricianSP;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import java.util.List;


public interface ElectricainRepo extends JpaRepository<ElectricianSP, Long> {
//    List<ElectricianSP> findByLocation(String location);

    @Query("SELECT e, u FROM ElectricianSP e JOIN Users u ON e.uId = u.id")
    List<Object[]> getAllElectriciansWithUsers();

    @Query("SELECT e, u FROM ElectricianSP e JOIN Users u ON e.uId = u.id WHERE e.eID = ?1")
    List<Object[]> findElectricianWithUserById(Long eId);
//    Electrician findDataById(Long id);
//    List<ElectricianSP> findByName(String name);
    @Query("SELECT e, u FROM ElectricianSP e JOIN Users u ON e.uId = u.id WHERE e.specz = ?1")
    List<Object[]> findBySpecz(String specialization);
}
