package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
public class InterviewProcess {
    @Id
    @SequenceGenerator(name = "proccessIdSeqGen", sequenceName = "proccessIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "proccessIdSeqGen")
    private Long proccessId;

    @Column
    private String description;

    public InterviewProcess() {
    }
}
