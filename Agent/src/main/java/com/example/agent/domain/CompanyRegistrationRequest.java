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
    private String companyOwner;

    @Column
    private String companyContactInfo;

    @Column
    private String companyDescription;

    public CompanyRegistrationRequest() {
    }

    public CompanyRegistrationRequest(String companyOwner, String companyContactInfo, String companyDescription) {
        this.companyOwner = companyOwner;
        this.companyContactInfo = companyContactInfo;
        this.companyDescription = companyDescription;
    }
}
