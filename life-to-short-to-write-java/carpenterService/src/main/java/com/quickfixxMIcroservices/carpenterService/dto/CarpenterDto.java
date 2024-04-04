package com.quickfixxMIcroservices.carpenterService.dto;

import lombok.*;

import java.util.List;

@Getter
@Setter
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class CarpenterDto {

    private String name;
    private long contactinfo;
    private String location;
    private String address;
    private String experience;
    private List<String> qualification;
    private String specialization;
    private int rating;
}
