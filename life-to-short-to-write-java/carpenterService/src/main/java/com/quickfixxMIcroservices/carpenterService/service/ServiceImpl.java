package com.quickfixxMIcroservices.carpenterService.service;

import com.quickfixxMIcroservices.carpenterService.dto.CarpenterDto;
import com.quickfixxMIcroservices.carpenterService.model.Carpenter;

import java.util.List;

public interface ServiceImpl {

    List<Carpenter> getAllCarepenter();
    Carpenter getById(long id);
    Carpenter createCarpenter(CarpenterDto plumberDto);
    Carpenter delete(long id);
    List<Carpenter> getByName(String name);
}
