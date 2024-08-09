# Library-Management System

there are several services : 
- [auth service](https://github.com/storyofhis/backend-syn/blob/main/user-svc/README.md) : to handle users and authentication
- [book service](https://github.com/storyofhis/backend-syn/blob/main/book-svc/README.md) : to manage books
- [authors service](https://github.com/storyofhis/backend-syn/blob/main/author-svc/README.md) : to manage authors
- [category service](https://github.com/storyofhis/backend-syn/blob/main/category-svc/README.md) : to manage book categories

schema for Entity-Relationship Diagram (ERD) : 
- BookService
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
                    
- AuthorService         
-------------------   
| author          |   
|-----------------|   
| id              |
| name            |      
| biography       |      
| created_at      |     
| updated_at      |      
-------------------     
                         
- CategoryService          
-------------------    
| category        |      
|-----------------|     
| id              |
| name            |        
| description     |        
| created_at      |        
| updated_at      |        
-------------------       
                        
- BorrowingService         
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

