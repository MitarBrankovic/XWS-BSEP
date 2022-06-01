package com.example.agent.repository;

import com.example.agent.domain.Sallary;
import org.springframework.data.jpa.repository.JpaRepository;

public interface SallaryRepository extends JpaRepository<Sallary, Long> {
}
