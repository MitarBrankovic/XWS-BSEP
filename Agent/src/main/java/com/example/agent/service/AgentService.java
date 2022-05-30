package com.example.agent.service;

import com.example.agent.dtos.UserRegistrationDTO;

public interface AgentService {
    void saveUser(UserRegistrationDTO userRegistrationDTO);
}
