package com.example.agent.repository;

import com.example.agent.domain.InterviewProcess;
import org.springframework.data.jpa.repository.JpaRepository;

public interface InterviewProcessRepository extends JpaRepository<InterviewProcess, Long> {
}
