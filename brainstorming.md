# Backend

- communicate with the DB
- API
  - CRUD/GraphQL
  - quizzes management
  - generate PDF with stats
- users management
- authentication and authorization

# Db Structure

- users
- quizes content
- answers

# Tables

Users

- id
- name
- username/email
- salt
- token/oauth2

Quizes

- id
- category (work, nature, pets, personal, others)
- question (text + answers + points/importance)
- audience
-
