package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
public class InterviewProcess {
    @Id
    @SequenceGenerator(name = "proccessIdSeqGen", sequenceName = "proccessIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "proccessIdSeqGen")
    private Long id;

    @Column
    private String description;

    @Column
    private String userSignature;

    @Column
    private String username;

    public InterviewProcess() {
    }

    public InterviewProcess(String description, String userSignature, String username) {
        this.description = description;
        this.userSignature = userSignature;
        this.username = username;
    }
}
