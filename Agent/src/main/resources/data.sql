insert into salary (id, salary_value, user_id) values (1, 500, 3)

insert into open_position (id, position_name) values (1, 'React developer')

insert into open_position_salaries (open_position_id, salaries_id) values (1, 1)

insert into interview_process (id, description, user_signature, username) values (1, 'Interview process', 'Vue.js developer (Junior)', 'petaragent')

insert into comment_on_company (id, comment, user_signature, username) values (1, 'Comment on company', 'Software developer (Medior)', 'petaragent')

insert into mark (id, mark) values (1, 5)

insert into company (id, name, contact_info, description, username) values (1, 'Majkrosoft za sirotinju', 'Company contact info', 'Company description', 'stevanagent')

insert into company_comments (company_id, comments_id) values (1, 1)

insert into company_marks (company_id, marks_id) values (1, 1)

insert into company_interview_processes (company_id, interview_processes_id) values (1, 1)

insert into company_open_positions (company_id, open_positions_id) values (1, 1)

insert into role(id, name) values (0, 'ROLE_ADMIN')
insert into role(id, name) values (1, 'ROLE_COMMON')
insert into role(id, name) values (2, 'ROLE_COMPANY_OWNER')

insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role_id, company_id) values (1, 'Marko', 'Markovic', 'markoagent', '$2a$12$aER/Nl0mc8hfjAqCyg79CeF/E2lhbSTESZi95RTxPX8KJTQut17Ju', '2000-01-01', 0, null )
insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role_id, company_id) values (2, 'Petar', 'Petrovic', 'petaragent', '$2a$12$aER/Nl0mc8hfjAqCyg79CeF/E2lhbSTESZi95RTxPX8KJTQut17Ju', '2001-01-01', 1, null )
insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role_id, company_id) values (3, 'Sima', 'Kesic', 'simakesic', '$2a$12$aER/Nl0mc8hfjAqCyg79CeF/E2lhbSTESZi95RTxPX8KJTQut17Ju', '2001-01-01', 1, null )
insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role_id, company_id) values (4, 'Stevan', 'Stevanovic', 'stevanagent', '$2a$12$aER/Nl0mc8hfjAqCyg79CeF/E2lhbSTESZi95RTxPX8KJTQut17Ju', '2002-01-01', 2, 1)


insert into company_registration_request (id, company_owner_username, company_owner_name, company_name, company_contact_info, company_description, username) values (1, 'simakesic', 'Sima Kesic', 'Microsoft', 'Company contact info', 'Company description', 'simakesic')