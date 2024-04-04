package com.quickfixxMicroservices.PlumberService.service;

import com.quickfixxMicroservices.PlumberService.dto.PlumberDto;
import com.quickfixxMicroservices.PlumberService.model.Plumber;
import com.quickfixxMicroservices.PlumberService.repository.PlumberRepo;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class PlumberService implements ServiceInterface{

    private final PlumberRepo plumberRepo;
    @Override
    public List<Plumber> getAllPlumbers() {
        return  plumberRepo.findAll().stream().toList();
    }

    @Override
    public Plumber getById(long id) {
        return plumberRepo.getReferenceById(id);
    }

    @Override
    public Plumber createPlumber(PlumberDto plumberDto) {

        Plumber plumber = new Plumber();
        plumber.setName(plumberDto.getName());
        plumber.setLocation(plumberDto.getLocation());
        plumber.setAddress(plumberDto.getAddress());
        plumber.setContactinfo(plumberDto.getContactinfo());
        plumber.setExperience(plumberDto.getExperience());
        plumber.setQualification(plumberDto.getQualification());

        plumberRepo.save(plumber);
        return plumber;
    }

    @Override
    public Plumber delete(long id) {
        Plumber plumber = plumberRepo.getReferenceById(id);
        plumberRepo.delete(plumber);
        return plumber;
    }

    @Override
    public List<Plumber> getByName(String name) {
        List<Plumber> plumberList = plumberRepo.findAllByName(name).stream().toList();
        return plumberList;
    }

    @Override
    public List<Plumber> getByField(String field) {
        List<Plumber> plumberList = plumberRepo.findByField(field);
        return plumberList;
    }
}
