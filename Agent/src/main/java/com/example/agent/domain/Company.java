package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.List;
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
    private String name;

    @Column
    private String contactInfo;

    @Column
    private String description;

    @Column
    private String username;

    @OneToMany(fetch = FetchType.EAGER)
    private Set<InterviewProcess> interviewProcesses;

    @OneToMany(fetch = FetchType.EAGER)
    private Set<CommentOnCompany> comments;

    @OneToMany(fetch = FetchType.EAGER)
    private Set<OpenPosition> openPositions;

    @OneToMany(fetch = FetchType.EAGER)
    private Set<Mark> marks;

    public Company() {
    }

    public Company(String name, String contactInfo, String description, String username) {
        this.name = name;
        this.contactInfo = contactInfo;
        this.description = description;
        this.username = username;
    }


}
