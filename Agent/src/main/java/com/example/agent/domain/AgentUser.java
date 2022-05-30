package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.Date;

@Entity
@Getter
@Setter
public class AgentUser {

    @Id
    @SequenceGenerator(name = "agentUserIdSeqGen", sequenceName = "agentUserIdSeq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "agentUserIdSeqGen")
    private Long id;

    @Column
    private String username;

    @Column
    private String password;

    @Column
    private String firstName;

    @Column
    private String lastName;

    @Column
    private Date dateOfBirth;

    @Column
    private UserRole role;

    @OneToOne
    @JoinColumn(name="companyId", referencedColumnName="id")
    private Company company;

    public AgentUser() {}
}
