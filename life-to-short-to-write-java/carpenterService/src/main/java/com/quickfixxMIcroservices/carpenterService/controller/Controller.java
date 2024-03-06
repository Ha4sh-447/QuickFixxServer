package com.quickfixxMIcroservices.carpenterService.controller;

import com.quickfixxMIcroservices.carpenterService.dto.CarpenterDto;
import com.quickfixxMIcroservices.carpenterService.model.Carpenter;
import com.quickfixxMIcroservices.carpenterService.service.CarpenterService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Repository;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("api/carpenter")
@RequiredArgsConstructor
public class Controller {
    private final CarpenterService carpenterService;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<Carpenter> getAll(){
        return carpenterService.getAllCarepenter();
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Carpenter getById(@PathVariable("id") String idStr){
        Long id = Long.parseLong(idStr);
        return carpenterService.getById(id);
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Carpenter create(@RequestBody CarpenterDto carpenterDto){
        return carpenterService.createCarpenter(carpenterDto);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Carpenter remove(@PathVariable("id") String idStr){
        Long id = Long.parseLong(idStr);
        return carpenterService.delete(id);
    }

    @GetMapping("/{name}")
    @ResponseStatus(HttpStatus.OK)
    public List<Carpenter> getByName(@PathVariable("name") String name){
        return carpenterService.getByName(name);
    }


}
