package com.quickfixxMicroservice.electricianService.service;

import com.quickfixxMicroservice.electricianService.dto.ElectricanWithUserDto;
import com.quickfixxMicroservice.electricianService.model.Electrician;
import com.quickfixxMicroservice.electricianService.model.ElectricianSP;

import java.util.List;
import java.util.Optional;

public interface ElectricianService {
    public List<ElectricianSP> getAllElectrician();
    public List<Object[]> getAllElectriciansWithUsers();
    public Optional<ElectricanWithUserDto> getByID(Long id);
    public List<Electrician> getByName(String LastName);
    public List<Electrician> getByLocation(String location);
    public void createElectrician(ElectricianSP electrician);
    public void removeElectrician(Long id);
    public List<Object[]> getByspecialization(String specialization);
}
