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
    @SequenceGenerator(name = "positonIdSeqGen", sequenceName = "positonIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "positonIdSeqGen")
    private Long id;

    @Column
    private String positionName;

    @ManyToMany
    private Set<Sallary> sallarys;

    public OpenPosition() {
    }


}
