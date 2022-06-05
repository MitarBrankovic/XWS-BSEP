package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDateTime;
import java.util.Date;

@Entity
@Getter
@Setter
public class AgentUser {

    @Id
    @SequenceGenerator(name = "agentUserIdSeqGen", sequenceName = "agentUserIdSeq", initialValue = 5, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "agentUserIdSeqGen")
    private Long id;

    @Column(unique = true)
    private String username;

    @Column
    private String password;

    @Column
    private String firstName;

    @Column
    private String lastName;

    @Column
    private LocalDateTime dateOfBirth;

    @Column
    private UserRole role;

    @Column
    private String apiToken;

    @OneToOne
    @JoinColumn(name="companyId", referencedColumnName="id")
    private Company company;

    public AgentUser() {}

    public AgentUser(String username, String password, String firstName, String lastName, LocalDateTime dateOfBirth, UserRole role) {
        this.username = username;
        this.password = password;
        this.firstName = firstName;
        this.lastName = lastName;
        this.dateOfBirth = dateOfBirth;
        this.role = role;
    }
}
