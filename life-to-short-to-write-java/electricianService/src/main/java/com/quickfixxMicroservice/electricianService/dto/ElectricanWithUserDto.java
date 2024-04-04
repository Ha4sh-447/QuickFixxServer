package com.quickfixxMicroservice.electricianService.dto;

import com.quickfixxMicroservice.electricianService.model.ElectricianSP;
import com.quickfixxMicroservice.electricianService.model.Users;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;


@Data
@AllArgsConstructor
@NoArgsConstructor
public class ElectricanWithUserDto {
    private ElectricianSP electrician;
    private Users user;
}
