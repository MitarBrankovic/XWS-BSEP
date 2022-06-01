package com.example.agent.repository;

import com.example.agent.domain.CommentOnCompany;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CommentOnCompanyRepository extends JpaRepository<CommentOnCompany, Long> {
}
