package com.example.agent.service;

import com.example.agent.domain.AgentUser;
import com.example.agent.domain.UserRole;
import com.example.agent.dtos.UserRegistrationDTO;
import com.example.agent.repository.AgentUserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AgentServiceImpl implements AgentService {

    @Autowired
    private AgentUserRepository agentUserRepository;

    @Override
    public void saveUser(UserRegistrationDTO userRegistrationDTO) {
        agentUserRepository.save(new AgentUser(userRegistrationDTO.getUsername(),
                userRegistrationDTO.getPassword(),
                userRegistrationDTO.getFirstName(),
                userRegistrationDTO.getLastName(),
                userRegistrationDTO.getDateOfBirth(),
                UserRole.Common));
    }

}
