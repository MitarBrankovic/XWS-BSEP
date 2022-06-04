package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
public class CommentOnCompany {
    @Id
    @SequenceGenerator(name = "commentIdSeqGen", sequenceName = "commentIdSeq", initialValue = 2, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "commentIdSeqGen")
    private Long id;

    @Column
    private String comment;

    @Column
    private String userSignature;

    @Column
    private String username;

    public CommentOnCompany() {
    }

    public CommentOnCompany(String comment, String userSignature, String username) {
        this.comment = comment;
        this.userSignature = userSignature;
        this.username = username;
    }
}
