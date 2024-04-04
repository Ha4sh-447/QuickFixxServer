package com.quickfixxMicroservice.electricianService.model;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name = "t_electriciandb_2")
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class ElectricianSP{
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long eID;
    private Long uId;
    private String specz;
    private String experience;
    private String address;
    private float rating;
    private String shopname;
}
