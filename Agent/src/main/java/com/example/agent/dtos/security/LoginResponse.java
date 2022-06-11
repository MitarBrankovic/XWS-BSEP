package com.example.agent.dtos.security;

import com.example.agent.domain.AgentUser;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class LoginResponse {
    private String token;
    private AgentUser user;
}
