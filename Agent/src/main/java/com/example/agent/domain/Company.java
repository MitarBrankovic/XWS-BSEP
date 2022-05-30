package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;
import org.aspectj.apache.bcel.classfile.Module;

import javax.persistence.*;
import java.util.Set;

@Entity
@Getter
@Setter
public class Company {
    @Id
    @SequenceGenerator(name = "companyIdSeqGen", sequenceName = "companyIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "companyIdSeqGen")
    private Long companyId;

    @Column
    private String contactInfo;

    @Column
    private String description;

    @OneToMany
    private Set<InterviewProcess> interviewProcesses;

    @OneToMany
    private Set<CommentOnCompany> comments;

    @OneToMany
    private Set<OpenPosition> openPositions;

    public Company() {
    }


}
