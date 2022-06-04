package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Setter
@Getter
public class Sallary {

    @Id
    @SequenceGenerator(name = "sallaryIdSeqGen", sequenceName = "sallaryIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "sallaryIdSeqGen")
    private Long id;

    @Column
    private Double sallaryValue;

    @Column
    private Long userId;

    public Sallary() {
    }

    public Sallary(Double sallaryValue, Long userId) {
        this.sallaryValue = sallaryValue;
        this.userId = userId;
    }
}
