package com.example.agent.service;

import com.example.agent.domain.AgentUser;
import com.example.agent.dtos.security.LoginRequest;
import com.example.agent.dtos.security.LoginResponse;
import com.example.agent.security.jwt.JwtUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import javax.naming.AuthenticationException;

@Service
public class AuthService {

    @Autowired
    private AgentServiceImpl userService;
    @Autowired
    private AuthenticationManager authenticationManager;
    @Autowired
    private JwtUtils jwtUtils;

    public LoginResponse login(LoginRequest request) throws AuthenticationException {
        authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(request.getUsername(), request.getPassword()));
        AgentUser user = userService.loadUserByUsername(request.getUsername());
        String token = jwtUtils.generateToken(user);

        return new LoginResponse(token, user);
    }

    public AgentUser getCurrentUser() {
        return (AgentUser) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
    }
}
