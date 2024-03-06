package com.quickfixxMicroservice.electricianService.service;

import com.quickfixxMicroservice.electricianService.dto.ElectricianDto;
import com.quickfixxMicroservice.electricianService.model.Electrician;
import com.quickfixxMicroservice.electricianService.repository.ElectricainRepo;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

/**
 * Contains method to perform CRUD operations related to Electrician database
 * Todo- Fix the bug in the createElectricians method which inserts all the fields as null except the auto generated Id field.
 */
@Service
@RequiredArgsConstructor
public class ElectricianServiceImpl implements ElectricianService{

    private final ElectricainRepo electricainRepo;
    @Override
    public List<Electrician> getAllElectrician() {
        List<Electrician> electricianList = electricainRepo.findAll().stream().toList();
        for (Electrician ele:
             electricianList) {
            System.out.println(ele.getId() + " " + ele.getName());
        }
        return electricianList;
    }

    @Override
    public Optional<Electrician> getByID(Long id) {
        return electricainRepo.findById(id);
    }

    @Override
    public List<Electrician> getByName(String name) {
        List<Electrician> electricianList = electricainRepo.findByName(name).stream().toList();
        return electricianList;
    }

    @Override
    public List<Electrician> getByLocation(String location) {
        return electricainRepo.findByLocation(location);
    }

    @Override
    public void createElectrician(ElectricianDto electricianDto) {
        Electrician electrician = new Electrician();
//        electrician.setId(Long.valueOf(String.valueOf(UUID.randomUUID())));

       electrician.setName(electricianDto.getName());
       electrician.setLocation(electricianDto.getLocation());
       electrician.setAddress(electricianDto.getAddress());
       electrician.setContact(electricianDto.getContact());
       electrician.setExperience(electricianDto.getExperience());
       electrician.setQualification(electricianDto.getQualification());
        electricainRepo.save(electrician);

    }

    @Override
    public void removeElectrician(Long id) {
        electricainRepo.deleteById(id);
    }
}
