package com.example.agent.domain;

import lombok.Getter;
import lombok.Setter;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import javax.persistence.*;
import javax.validation.constraints.NotEmpty;
import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Date;
import java.util.List;

@Entity
@Getter
@Setter
public class AgentUser implements Serializable, UserDetails {

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

    @ManyToOne()
    @JoinColumn(name = "role_id")
    private Role role;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        List<Role> collection = new ArrayList<>();
        collection.add(this.role);
        return collection;
    }

    @Column
    private String apiToken;

    @OneToOne
    @JoinColumn(name="companyId", referencedColumnName="id")
    private Company company;

    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    public boolean isAccountNonLocked() {
        return true;
    }

    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return true;
    }

    public AgentUser() {}

    public AgentUser(String username, String password, String firstName, String lastName, LocalDateTime dateOfBirth, Role role) {
        this.username = username;
        this.password = password;
        this.firstName = firstName;
        this.lastName = lastName;
        this.dateOfBirth = dateOfBirth;
        this.role = role;
    }
}
