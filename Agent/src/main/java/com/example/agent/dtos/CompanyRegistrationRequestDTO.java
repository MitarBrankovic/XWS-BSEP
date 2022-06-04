package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CompanyRegistrationRequestDTO {
    private String companyOwnerUsername;
    private String companyOwnerName;
    private String companyName;
    private String companyContactInfo;
    private String companyDescription;
    private String username;

    public CompanyRegistrationRequestDTO() {
    }

    public CompanyRegistrationRequestDTO(String companyName, String companyOwnerUsername, String companyOwnerName, String companyContactInfo, String companyDescription, String username) {
        this.companyName = companyName;
        this.companyOwnerUsername = companyOwnerUsername;
        this.companyOwnerName = companyOwnerName;
        this.companyContactInfo = companyContactInfo;
        this.companyDescription = companyDescription;
        this.username = username;
    }
}
