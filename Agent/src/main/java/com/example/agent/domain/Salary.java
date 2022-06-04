package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Setter
@Getter
public class Salary {

    @Id
    @SequenceGenerator(name = "salaryIdSeqGen", sequenceName = "salaryIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "salaryIdSeqGen")
    private Long id;

    @Column
    private Double salaryValue;

    @Column
    private Long userId;

    public Salary() {
    }

    public Salary(Double salaryValue, Long userId) {
        this.salaryValue = salaryValue;
        this.userId = userId;
    }
}
