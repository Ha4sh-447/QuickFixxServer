package com.quickfixxMicroservice.electricianService.service;

import com.quickfixxMicroservice.electricianService.dto.ElectricanWithUserDto;
import com.quickfixxMicroservice.electricianService.model.Electrician;
import com.quickfixxMicroservice.electricianService.model.ElectricianSP;
import com.quickfixxMicroservice.electricianService.model.Users;
import com.quickfixxMicroservice.electricianService.repository.ElectricainRepo;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

/**
 * Contains method to perform CRUD operations related to Electrician database
 * Todo- Fix the bug in the createElectricians method which inserts all the fields as null except the auto generated Id field.
 */
@Service
@RequiredArgsConstructor
public class ElectricianServiceImpl implements ElectricianService{

    private final ElectricainRepo electricainRepo;

    @Override
    public List<Object[]> getAllElectriciansWithUsers() {
        return electricainRepo.getAllElectriciansWithUsers();
    }
    @Override
    public List<ElectricianSP> getAllElectrician() {
        List<ElectricianSP> electricianList = electricainRepo.findAll().stream().toList();
        return electricianList;
    }

        @Override
        public Optional<ElectricanWithUserDto> getByID(Long eid) {
            List<Object[]> resultList = electricainRepo.findElectricianWithUserById(eid);

            if (!resultList.isEmpty()) {
                Object[] result = resultList.get(0);
                System.out.println(result);
                ElectricianSP electrician = (ElectricianSP) result[0];
                Users user = (Users) result[1];

                ElectricanWithUserDto electricanWithUserDto = new ElectricanWithUserDto();
                electricanWithUserDto.setElectrician(electrician);
                electricanWithUserDto.setUser(user);
                // Populate ElectricanWithUserDto properties with Electrician and User data

                return Optional.of(electricanWithUserDto);

            } else {
                return Optional.empty();
            }
        }

    @Override
    public List<Electrician> getByName(String name) {
//        List<Electrician> electricianList = electricainRepo.findByName(name).stream().toList();
//        return electricianList;
        return null;
    }

    @Override
    public List<Electrician> getByLocation(String location) {
//        return electricainRepo.findByLocation(location);
            return null;
    }

    @Override
    public void createElectrician(ElectricianSP electriciansp) {
//        Electrician electrician = new Electrician();
        System.out.println(electriciansp.getUId()+" "+ electriciansp.getAddress());
        System.out.println(electriciansp.getEID()+ " "+ electriciansp.getExperience()+" "+ electriciansp.getShopname());
        electricainRepo.save(electriciansp);

    }

    @Override
    public void removeElectrician(Long id) {
        electricainRepo.deleteById(id);
    }

    @Override
    public List<Object[]> getByspecialization(String specialization) {
//        List<ElectricianSP> elecByspecialization = electricainRepo.findByspecz(specialization);
        List<Object[]> resultList = electricainRepo.findBySpecz(specialization);
//
//        if (!resultList.isEmpty()) {
//            Object[] result = resultList.get(0);
//            System.out.println(result);
//            ElectricianSP electrician = (ElectricianSP) result[0];
//            Users user = (Users) result[1];
//
//            ElectricanWithUserDto electricanWithUserDto = new ElectricanWithUserDto();
//            electricanWithUserDto.setElectrician(electrician);
//            electricanWithUserDto.setUser(user);
//            // Populate ElectricanWithUserDto properties with Electrician and User data
//
//            return Optional.of(electricanWithUserDto);
//
//        } else {
//            return Optional.empty();
//        }

        return resultList;
    }
}
