package com.example.agent.repository;

import com.example.agent.domain.Salary;
import org.springframework.data.jpa.repository.JpaRepository;

public interface SallaryRepository extends JpaRepository<Salary, Long> {
}
