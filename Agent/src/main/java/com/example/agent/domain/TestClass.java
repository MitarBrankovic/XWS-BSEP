package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.Entity;
import javax.persistence.*;

@Entity
@Getter
@Setter
public class TestClass {
    @Id
    @SequenceGenerator(name = "testIdSeqGen", sequenceName = "testIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "testIdSeqGen")
    private int id;

    @Column
    private String name;

    public TestClass() {
    }
}
