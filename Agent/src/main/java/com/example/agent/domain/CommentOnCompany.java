package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
public class CommentOnCompany {
    @Id
    @SequenceGenerator(name = "commentIdSeqGen", sequenceName = "commentIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "commentIdSeqGen")
    private Long id;

    @Column
    private String comment;

    @Column
    private String userSignature;

    public CommentOnCompany() {
    }
}