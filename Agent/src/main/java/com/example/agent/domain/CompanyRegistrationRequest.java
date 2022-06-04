package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
public class CompanyRegistrationRequest {
    @Id
    @SequenceGenerator(name = "companyRegistrationRequestIdSeqGen", sequenceName = "companyRegistrationRequestIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "companyRegistrationRequestIdSeqGen")
    private Long id;

    @Column
    private String companyOwnerUsername;

    @Column
    private String companyOwnerName;

    @Column
    private String companyName;

    @Column
    private String companyContactInfo;

    @Column
    private String companyDescription;

    @Column
    private String username;

    public CompanyRegistrationRequest() {
    }

    public CompanyRegistrationRequest(String companyOwnerUsername, String companyOwnerName, String companyName, String companyContactInfo, String companyDescription, String username) {
        this.companyOwnerUsername = companyOwnerUsername;
        this.companyOwnerName = companyOwnerName;
        this.companyName = companyName;
        this.companyContactInfo = companyContactInfo;
        this.companyDescription = companyDescription;
        this.username = username;
    }
}
