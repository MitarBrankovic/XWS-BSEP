package com.example.agent.repository;

import com.example.agent.domain.CompanyRegistrationRequest;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CompanyRegistrationRequestRepository extends JpaRepository<CompanyRegistrationRequest, Long> {
    void removeAllByCompanyOwnerUsername(String username);
}
