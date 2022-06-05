package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.Set;

@Entity
@Setter
@Getter
public class OpenPosition {
    @Id
    @SequenceGenerator(name = "positionIdSeqGen", sequenceName = "positionIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "positionIdSeqGen")
    private Long id;

    @Column
    private String positionName;

    @Column
    private String description;

    @Column
    private String criteria;

    @ManyToMany(fetch = FetchType.EAGER)
    private Set<Salary> salaries;

    public OpenPosition() {
    }

    public OpenPosition(String positionName, String description, String criteria) {
        this.positionName = positionName;
        this.description = description;
        this.criteria = criteria;
    }


}
