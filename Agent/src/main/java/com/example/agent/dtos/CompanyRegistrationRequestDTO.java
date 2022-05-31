package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CompanyRegistrationRequestDTO {
    private String companyOwner;
    private String companyContactInfo;
    private String companyDescription;

    public CompanyRegistrationRequestDTO() {
    }
}
