package com.example.agent.service;

import com.example.agent.domain.*;
import com.example.agent.dtos.CompanyInfoDTO;
import com.example.agent.dtos.CompanyRegistrationRequestDTO;
import com.example.agent.dtos.UserRegistrationDTO;
import com.example.agent.repository.AgentUserRepository;
import com.example.agent.repository.CompanyRegistrationRequestRepository;
import com.example.agent.repository.CompanyRepository;
import com.example.agent.repository.OpenPositionRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AgentServiceImpl implements AgentService {

    @Autowired
    private AgentUserRepository agentUserRepository;

    @Autowired
    private CompanyRegistrationRequestRepository companyRegistrationRequestRepository;

    @Autowired
    private CompanyRepository companyRepository;

    @Autowired
    private OpenPositionRepository openPositionRepository;

    @Override
    public void saveUser(UserRegistrationDTO userRegistrationDTO) {
        agentUserRepository.save(new AgentUser(userRegistrationDTO.getUsername(),
                userRegistrationDTO.getPassword(),
                userRegistrationDTO.getFirstName(),
                userRegistrationDTO.getLastName(),
                userRegistrationDTO.getDateOfBirth(),
                UserRole.Common));
    }

    @Override
    public void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto) {
        companyRegistrationRequestRepository.save(new CompanyRegistrationRequest(dto.getCompanyOwnerUsername(),
                dto.getCompanyOwnerName(),
                dto.getCompanyContactInfo(),
                dto.getCompanyDescription()));
    }

    @Override
    public void saveCompany(CompanyRegistrationRequestDTO dto) {
        AgentUser commonUser = agentUserRepository.findAgentUserByUsername(dto.getCompanyOwnerUsername());
        Company newCompany = new Company(dto.getCompanyContactInfo(), dto.getCompanyDescription());

        commonUser.setCompany(newCompany);
        commonUser.setRole(UserRole.CompanyOwner);

        companyRepository.save(newCompany);
        agentUserRepository.save(commonUser);
    }

    @Override
    public void editCompanyInfo(CompanyInfoDTO dto) {
        Company company = companyRepository.findById(dto.getId()).orElseGet(null);
        company.setContactInfo(dto.getContactInfo());
        company.setDescription(dto.getDescription());
        companyRepository.save(company);
    }

    @Override
    public void addOpenPosition(Long companyId, String positionName) {
        OpenPosition newOpenPosition = new OpenPosition(positionName);
        openPositionRepository.save(newOpenPosition);

        Company company = companyRepository.findById(companyId).orElseGet(null);
        company.getOpenPositions().add(newOpenPosition);

        companyRepository.save(company);
    }
}
