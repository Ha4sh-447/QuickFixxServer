package com.quickfixxMicroservice.electricianService.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class ElectricianDto {
    private Long id;
    private String name;
    private Long contact;
    private String location;
    private String address;
    private String experience;
    private List<String> qualification;
    private float rating;
}
