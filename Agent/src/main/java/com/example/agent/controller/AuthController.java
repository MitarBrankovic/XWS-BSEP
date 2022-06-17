package com.example.agent.controller;

import com.example.agent.dtos.security.LoginRequest;
import com.example.agent.dtos.security.LoginResponse;
import com.example.agent.service.AgentServiceImpl;
import com.example.agent.service.AuthService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

import javax.naming.AuthenticationException;

@Controller
@CrossOrigin(origins = "https://localhost:4201")
@RequestMapping("/api/auth")
public class AuthController {

    @Autowired
    private AuthService authService;

    @RequestMapping("/login")
    public ResponseEntity<LoginResponse> login(@RequestBody LoginRequest request) throws AuthenticationException {
        return new ResponseEntity<>(authService.login(request), HttpStatus.OK);
    }
}
