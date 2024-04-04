package com.quickfixxMicroservices.PlumberService.dto;


import jakarta.persistence.Column;
import lombok.*;

import java.util.List;

@Getter
@Setter
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class PlumberDto {

    private String name;
    private long contactinfo;
    private String location;
    private String address;
    private String experience;
    private List<String> qualification;
    private String field;
    private int rating;
}
