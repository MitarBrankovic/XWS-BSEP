package com.example.agent.controller;

import com.example.agent.dtos.UserRegistrationDTO;
import com.example.agent.service.AgentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("/api/agent")
public class AgentController {

    @Autowired
    private AgentService agentService;

    @RequestMapping("/registerUser")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<UserRegistrationDTO> createAdventure(@RequestBody UserRegistrationDTO dto){
        agentService.saveUser(dto);
        return new ResponseEntity<UserRegistrationDTO>(dto, HttpStatus.CREATED);
    }
}
