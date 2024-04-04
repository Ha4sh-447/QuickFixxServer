package com.quickfixxMicroservice.electricianService.controller;


import com.quickfixxMicroservice.electricianService.dto.ElectricanWithUserDto;
import com.quickfixxMicroservice.electricianService.dto.ElectricianDto;
import com.quickfixxMicroservice.electricianService.model.ElectricianSP;
import com.quickfixxMicroservice.electricianService.model.Users;
import com.quickfixxMicroservice.electricianService.service.ElectricianService;
import lombok.RequiredArgsConstructor;
import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/api/electrician")
@RequiredArgsConstructor
public class ElectricianFieldController {


    @Autowired
    ModelMapper modelMapper = new ModelMapper();


    private final ElectricianService electricianService;
    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public String createElectrician(@RequestBody ElectricianSP electricianDto){
        System.out.println(electricianDto.getUId()+ " "+ electricianDto.getExperience()+" "+ electricianDto.getRating()+" "+electricianDto.getAddress()+" "+electricianDto.getEID());
        electricianService.createElectrician(electricianDto);
        return "Electrician registered";
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<ElectricianDto> getALl(){
        List<Object[]> electricianAndUserList = electricianService.getAllElectriciansWithUsers();

        List<ElectricianDto> electricianDtoList = electricianAndUserList.stream().map(objects -> {
            ElectricianSP electrician = (ElectricianSP) objects[0];
            Users user = (Users) objects[1];

            ElectricianDto electricianDto = new ElectricianDto();
            electricianDto.setUserid((long) user.getId());
            electricianDto.setElecid(electrician.getEID());
            electricianDto.setName(electrician.getShopname());
            // Assuming contact and location are fields from User entity
            electricianDto.setContact(Long.valueOf(user.getContact()));
//            electricianDto.setLocation(user.getLocation());
            electricianDto.setAddress(electrician.getAddress());
            electricianDto.setExperience(electrician.getExperience());
            electricianDto.setSpecialization(electrician.getSpecz());
            // Assuming qualification is not available in the current data
            electricianDto.setQualification(null);
            electricianDto.setRating(electrician.getRating());
            electricianDto.setShopname(electrician.getShopname());
            electricianDto.setImage(user.getImage());
            return electricianDto;
        }).collect(Collectors.toList());

        if (electricianDtoList.isEmpty()) {
            System.out.println("No data found");
        }

        return electricianDtoList;
    }

    @GetMapping("/field/{field}")
    @ResponseStatus(HttpStatus.OK)
    public List<ElectricianDto>getElectricianByField(@PathVariable("field") String field){
//        List<ElectricianDto> dto = electricianService.getByspecialization(field).stream()
//                .map(electrician -> modelMapper.map(electrician, ElectricianDto.class))
//                .toList();
        List<Object[]>byFieldList = electricianService.getByspecialization(field);

        List<ElectricianDto> electricianDtoList = byFieldList.stream().map(objects -> {
            ElectricianSP electrician = (ElectricianSP) objects[0];
            Users user = (Users) objects[1];

            ElectricianDto electricianDto = new ElectricianDto();
            electricianDto.setUserid((long) user.getId());
            electricianDto.setElecid(electrician.getEID());
            electricianDto.setName(user.getName());
            // Assuming contact and location are fields from User entity
            electricianDto.setContact(Long.valueOf(user.getContact()));
//            electricianDto.setLocation(user.getLocation());
            electricianDto.setAddress(electrician.getAddress());
            electricianDto.setExperience(electrician.getExperience());
            electricianDto.setSpecialization(electrician.getSpecz());
            // Assuming qualification is not available in the current data
            electricianDto.setQualification(null);
            electricianDto.setRating(electrician.getRating());

            return electricianDto;
        }).collect(Collectors.toList());

        if (electricianDtoList.isEmpty()) {
            System.out.println("No data found");
        }

        return electricianDtoList;
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Optional<ElectricianDto> getByID(@PathVariable("id") String idString){
        Long id= Long.parseLong(idString);
        Optional<ElectricanWithUserDto> optionalElectrician = electricianService.getByID(id);

        Users user = optionalElectrician.get().getUser();
        ElectricianSP electrician = optionalElectrician.get().getElectrician();

        ElectricianDto electricianDto = new ElectricianDto();
        electricianDto.setUserid((long) user.getId());
        electricianDto.setElecid(electrician.getEID());
        electricianDto.setName(user.getName());
        electricianDto.setContact(Long.valueOf(user.getContact()));
        electricianDto.setAddress(electrician.getAddress());
        electricianDto.setExperience(electrician.getExperience());
        electricianDto.setSpecialization(electrician.getSpecz());
        electricianDto.setQualification(null);
        electricianDto.setRating(electrician.getRating());
        electricianDto.setShopname(electrician.getShopname());
        electricianDto.setImage(user.getImage());

        if(optionalElectrician.isPresent()){
            return Optional.of(electricianDto);
        }else {
            return null;
        }
    }

    @DeleteMapping("/delete/{id}")
    @ResponseStatus(HttpStatus.OK)
    public String deleteID(@PathVariable("id") Long id){
        electricianService.removeElectrician(id);
        return "Electrician data removed successfully";
    }




}
