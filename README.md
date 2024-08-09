# Library-Management System

there are several services : 
- [auth service](https://github.com/storyofhis/backend-syn/blob/main/user-svc/README.md) : to handle users and authentication
- [book service](https://github.com/storyofhis/backend-syn/blob/main/book-svc/README.md) : to manage books
- [authors service](https://github.com/storyofhis/backend-syn/blob/main/author-svc/README.md) : to manage authors
- [category service](https://github.com/storyofhis/backend-syn/blob/main/category-svc/README.md) : to manage book categories

## Entity-Relationship Diagram (ERD) : 
this is the schema of database design systems
![Screenshot 2024-08-09 at 20 33 12](https://github.com/user-attachments/assets/05b54daf-d877-4b32-b5ef-14c9632d8f6c)

- book_db: database for `book-svc` only have book table
-------------------
| book            |
|-----------------|
| id              |
| title           |       
| author_id       |
| description     |      
| category_id     |
| stock           |   
| created_at      |   
| updated_at      |   
-------------------   
                    
- author_db: database for `author-svc` only have author table
-------------------   
| author          |   
|-----------------|   
| id              |
| name            |      
| biography       |      
| created_at      |     
| updated_at      |      
-------------------     
                         
- category_db: database for `category-svc` only have category table       
-------------------    
| category        |      
|-----------------|     
| id              |
| name            |        
| description     |        
| created_at      |        
| updated_at      |        
-------------------       
                        
- user_db: database for `user-svc` have two tables (borrowing and user) 
-------------------    
| borrowing       |
|-----------------|
| id              |
| user_id         |
| book_id         |                   
| borrow_date     |                   
| due_date        |                  
| return_date     |                   
| status          |                   
-------------------                   
                                      
- UserService                           
-------------------                   
| user            |                   
|-----------------|                   
| id              |
| username        |
| password_hash   |
| email           |
| created_at      |
| updated_at      |
-------------------

