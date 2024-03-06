package com.quickfixxMIcroservices.carpenterService.model;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(name = "t_carpenter")
public class Carpenter {


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
