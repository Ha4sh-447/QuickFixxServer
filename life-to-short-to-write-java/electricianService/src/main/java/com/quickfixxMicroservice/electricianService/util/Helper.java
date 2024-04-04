package com.quickfixxMicroservice.electricianService.util;

import com.quickfixxMicroservice.electricianService.dto.ElectricianDto;
import com.quickfixxMicroservice.electricianService.model.Electrician;

public class Helper {
    public ElectricianDto electricianToDto(Electrician ec){
        ElectricianDto dto = new ElectricianDto();
        dto.setAddress(ec.getAddress());
        dto.setElecid(ec.getId());
        dto.setName(ec.getName());
        dto.setContact(ec.getContact());
//        dto.setLocation(ec.getLocation());
        dto.setExperience(ec.getExperience());
        dto.setQualification(ec.getQualification());

        return dto;
    }

    public Electrician dtoToElectrician(ElectricianDto electricianDto){
        Electrician electrician = new Electrician();
        electrician.setAddress(electricianDto.getAddress());
        electrician.setId(electricianDto.getUserid());
        electrician.setName(electricianDto.getName());
        electrician.setContact(electricianDto.getContact());
//        electrician.setLocation(electricianDto.getLocation());
        electrician.setExperience(electricianDto.getExperience());
        electrician.setQualification(electricianDto.getQualification());

        return electrician;
    }
}
