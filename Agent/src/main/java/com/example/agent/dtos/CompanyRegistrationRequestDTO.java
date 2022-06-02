package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CompanyRegistrationRequestDTO {
    private String companyOwnerUsername;
    private String companyOwnerName;
    private String companyContactInfo;
    private String companyDescription;

    public CompanyRegistrationRequestDTO() {
    }

    public CompanyRegistrationRequestDTO(String companyOwnerUsername, String companyOwnerName, String companyContactInfo, String companyDescription) {
        this.companyOwnerUsername = companyOwnerUsername;
        this.companyOwnerName = companyOwnerName;
        this.companyContactInfo = companyContactInfo;
        this.companyDescription = companyDescription;
    }
}
