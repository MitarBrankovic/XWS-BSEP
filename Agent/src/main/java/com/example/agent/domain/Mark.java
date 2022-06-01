package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
public class Mark {
    @Id
    @SequenceGenerator(name = "markIdSeqGen", sequenceName = "markIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "markIdSeqGen")
    private Long id;

    @Column
    private Integer mark;

    public Mark() {
    }

    public Mark(Integer mark) {
        this.mark = mark;
    }
}
