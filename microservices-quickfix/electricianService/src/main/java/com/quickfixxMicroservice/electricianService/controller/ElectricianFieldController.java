package com.quickfixxMicroservice.electricianService.controller;


import com.fasterxml.jackson.databind.ObjectMapper;
import com.quickfixxMicroservice.electricianService.dto.ElectricianDto;
import com.quickfixxMicroservice.electricianService.model.Electrician;
import com.quickfixxMicroservice.electricianService.service.ElectricianService;
import lombok.RequiredArgsConstructor;
import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/electrician")
@RequiredArgsConstructor
public class ElectricianFieldController {


    @Autowired
    ModelMapper modelMapper = new ModelMapper();


    private final ElectricianService electricianService;
    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public String createElectrician(@RequestBody ElectricianDto electricianDto){
        electricianService.createElectrician(electricianDto);
        return "Electrician registered";
    }

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<ElectricianDto> getALl(){
        List<ElectricianDto> electricianList = electricianService.getAllElectrician().stream()
                .map(electrician -> modelMapper.map(electrician, ElectricianDto.class))
                .toList();


        electricianList.stream().map(electrician -> {
            System.out.println(electrician.getId() + " "+ electrician.getName() + electrician.getRating());
            return null;
        }).toList();



        if(electricianList.isEmpty()) System.out.println("No data found");
        return electricianList;
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Optional<Electrician> getByID(@PathVariable("id") String idString){
        Long id= Long.parseLong(idString);
        Optional<Electrician> optionalElectrician = electricianService.getByID(id);

        if(optionalElectrician.isPresent()){
            return electricianService.getByID((id));
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
