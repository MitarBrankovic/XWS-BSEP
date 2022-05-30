package com.example.agent.repository;

import com.example.agent.domain.AgentUser;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AgentUserRepository extends JpaRepository<AgentUser, Long> {
}
