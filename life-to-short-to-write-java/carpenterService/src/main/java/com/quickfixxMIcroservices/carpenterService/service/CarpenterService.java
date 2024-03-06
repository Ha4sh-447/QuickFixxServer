package com.quickfixxMIcroservices.carpenterService.service;


import com.quickfixxMIcroservices.carpenterService.dto.CarpenterDto;
import com.quickfixxMIcroservices.carpenterService.model.Carpenter;
import com.quickfixxMIcroservices.carpenterService.repository.CarpenterRepo;
import lombok.RequiredArgsConstructor;
import org.hibernate.sql.model.PreparableMutationOperation;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class CarpenterService implements ServiceImpl {
    private final CarpenterRepo carpenterRepo;


    @Override
    public List<Carpenter> getAllCarepenter() {
        List<Carpenter> carpenterList = carpenterRepo.findAll().stream().toList();
        return carpenterList;
    }

    @Override
    public Carpenter getById(long id) {
        Carpenter carpenter = carpenterRepo.getById(id);
        return carpenter;
    }

    @Override
    public Carpenter createCarpenter(CarpenterDto carpenterDto) {

        Carpenter carpenter = new Carpenter();
        carpenter.setName(carpenterDto.getName());
        carpenter.setLocation(carpenterDto.getLocation());
        carpenter.setAddress(carpenterDto.getAddress());
        carpenter.setContactinfo(carpenterDto.getContactinfo());
        carpenter.setExperience(carpenterDto.getExperience());
        carpenter.setQualification(carpenterDto.getQualification());

        carpenterRepo.save(carpenter);
        return carpenter;
    }

    @Override
    public Carpenter delete(long id) {
        Carpenter carpenter = carpenterRepo.getReferenceById(id);
        carpenterRepo.delete(carpenter);
        return carpenter;
    }

    @Override
    public List<Carpenter> getByName(String name) {
        List<Carpenter> carpenterList = carpenterRepo.findAllByName(name).stream().toList();
        return carpenterList;
    }
}
