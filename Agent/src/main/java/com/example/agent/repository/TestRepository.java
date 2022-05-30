package com.example.agent.repository;

import com.example.agent.domain.TestClass;
import org.springframework.data.jpa.repository.JpaRepository;

public interface TestRepository extends JpaRepository<TestClass, Integer> {
}
