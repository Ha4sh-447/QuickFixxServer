package com.quickfixxMicroservices.PlumberService.controller;

import com.quickfixxMicroservices.PlumberService.dto.PlumberDto;
import com.quickfixxMicroservices.PlumberService.model.Plumber;
import com.quickfixxMicroservices.PlumberService.service.PlumberService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("api/plumbing")
@RequiredArgsConstructor
public class Controller {
    private final PlumberService plumberService;

    @GetMapping
    @ResponseStatus(HttpStatus.OK)
    public List<Plumber> getAll(){
        return plumberService.getAllPlumbers();
    }

    @GetMapping("/{Id}")
    @ResponseStatus(HttpStatus.OK)
    public Plumber getById(@PathVariable("Id") String idString){
        Long id= Long.parseLong(idString);
        Plumber plumber = plumberService.getById(id);
        return plumber;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Plumber create(@RequestBody PlumberDto plumberDto){
       return plumberService.createPlumber(plumberDto);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Plumber remove(@PathVariable("id") String delId){
        Long id = Long.parseLong(delId);
        return plumberService.delete(id);
    }

    @GetMapping("/{name}")
    @ResponseStatus(HttpStatus.OK)
    public List<Plumber> getByName(@PathVariable("name") String name){
        return plumberService.getByName(name);
    }

}
