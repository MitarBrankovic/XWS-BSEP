package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.Set;

@Entity
@Getter
@Setter
public class Company {
    @Id
    @SequenceGenerator(name = "companyIdSeqGen", sequenceName = "companyIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "companyIdSeqGen")
    private Long id;

    @Column
    private String contactInfo;

    @Column
    private String description;

    @OneToMany
    private Set<InterviewProcess> interviewProcesses;

    @OneToMany(fetch = FetchType.EAGER)
    private Set<CommentOnCompany> comments;

    @OneToMany(fetch = FetchType.EAGER)
    private Set<OpenPosition> openPositions;

    public Company() {
    }

    public Company(String contactInfo, String description) {
        this.contactInfo = contactInfo;
        this.description = description;
    }


}
