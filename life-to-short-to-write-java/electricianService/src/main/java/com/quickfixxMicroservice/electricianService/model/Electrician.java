package com.quickfixxMicroservice.electricianService.model;

import jakarta.persistence.*;
import lombok.*;


import java.util.List;

@Entity
@Table(name = "t_electricianDB")
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class Electrician {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private Long contact;
    private String location;
    private String address;
    @Column(name = "experience")
    private String experience;
    private List<String> qualification;
    private int rating;
}
