package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Setter
@Getter
public class Sallary {

    @Id
    @SequenceGenerator(name = "sallaryIddSeqGen", sequenceName = "sallaryIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "sallaryIdSeqGen")
    private Long id;

    @Column
    private Double sallaryValue;


    public Sallary() {
    }
}
