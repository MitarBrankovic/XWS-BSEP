package com.example.agent.repository;

import com.example.agent.domain.OpenPosition;
import org.springframework.data.jpa.repository.JpaRepository;

public interface OpenPositionRepository extends JpaRepository<OpenPosition, Long> {
}
