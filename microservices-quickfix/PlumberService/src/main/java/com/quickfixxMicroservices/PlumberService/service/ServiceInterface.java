package com.quickfixxMicroservices.PlumberService.service;

import com.quickfixxMicroservices.PlumberService.dto.PlumberDto;
import com.quickfixxMicroservices.PlumberService.model.Plumber;

import java.util.List;

public interface ServiceInterface {

    List<Plumber> getAllPlumbers();
    Plumber getById(long id);
    Plumber createPlumber(PlumberDto plumberDto);
    Plumber delete(long id);
    List<Plumber> getByName(String name);
}
