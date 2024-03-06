package com.quickfixxMicroservice.electricianService.service;

import com.quickfixxMicroservice.electricianService.dto.ElectricianDto;
import com.quickfixxMicroservice.electricianService.model.Electrician;

import java.util.List;
import java.util.Optional;

public interface ElectricianService {
    public List<Electrician> getAllElectrician();
    public Optional<Electrician> getByID(Long id);
    public List<Electrician> getByName(String LastName);
    public List<Electrician> getByLocation(String location);
    public void createElectrician(ElectricianDto electrician);
    public void removeElectrician(Long id);
}
