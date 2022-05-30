package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;
import java.util.Date;

@Getter
@Setter
public class UserRegistrationDTO {
    private String username;

    private String password;

    private String firstName;

    private String lastName;

    private LocalDateTime dateOfBirth;

    public UserRegistrationDTO(){}

    public UserRegistrationDTO(String username, String password, String firstName, String lastName, LocalDateTime dateOfBirth) {
        this.username = username;
        this.password = password;
        this.firstName = firstName;
        this.lastName = lastName;
        this.dateOfBirth = dateOfBirth;
    }
}
