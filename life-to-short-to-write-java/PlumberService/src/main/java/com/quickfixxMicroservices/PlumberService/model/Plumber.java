package com.quickfixxMicroservices.PlumberService.model;


import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@Entity
@Table(name = "t_plumbers")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class Plumber {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;
    private String name;
    private long contactinfo;
    private String location;
    private String address;
    @Column(name = "experience")
    private String experience;
    private List<String> qualification;
}
