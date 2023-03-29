Work completed in Sprint 2: 
<<<<<<< HEAD
In sprint 2, our group was able to improve in both the frontend and backend. Regarding the frontend development, we were able to develop a filter function to find specific data entries. We were able to get the table size changing feature working so number of rows shown can be adjusted. In addition, we were able to get a filter click function on the table to work. There is now an add button for "Add Item" on the inventory page and currently working on adding forms to that pop up. For the backend, we redesigned our initial code to be more simplified. We managed to get the SQlite database working with tables inventory and users being created. We were able to get GORM to communicate to the database and allow for the insertion, deletion, and updating of the tuples within the tables. We also started integration of communication between backend and frontend.
=======
In sprint 2, our group was able to improve the functionality of both the frontend and backend. Regarding the frontend development, we were able to develop a filter function to find specific data entries. In addition, we were able to get the table size changing feature working so rows shown can be adjusted. For the backend, we redesigned our initial code to be more simplified. We managed to get the SQlite database working with our tables, inventory and users, allowing for the creation of tuples. We were able to get GORM to communicate to the database and allow for the insertion, deletion, finding of tuples, and updating of the tuples within the tables. We also started integration of communication between backend and frontend. We established the routes that will use JSON to send information between frontend and backend. We lastly updated our code to include documentation for the use of each functions and their appropriate routes. Both frontend and backend also created several test cases to determine if the applications are working as expected.
>>>>>>> d0f1ce639ff18c9299df83294476478ca61d16db

Unit tests and Cypress tests for frontend:
   The inventory page Cypress test expects to show the inventory page and that it is directed to 'http://localhost:4200/inventory-home-page' and for the paginator to be working. 
   The inventory page unit test tests the creation of the components in the inventory page. 
   
Unit tests for backend: 

TestMakeUser -- Tests creating a user. First tests function call from http request on an empty user with ID = 1, then, manually tests database entry for a fully populated sample user. If an error is thrown at any point during execution, the test fails. If the the final user in the database does not exactly equal the desired test user, the test fails, else, the test passes. This method is not mirrored in future tests, since it will be better to test actuall queries from the client in the next sprint.

TestUpdateUser -- Tests updating a user. Firsts tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails. Does not test actual update / alteration of data since that process is tied into the client side requests in our code, will be more efficient to test later.

TestRemoveUser -- Tests removing a user. Tests function call from http request. If an error is thrown at any point during execution, the test fails.

TestFindUser -- Tests find user function. Tests function call from http request on user from ID. If an error is thrown at any point during execution, the test fails.

TestInsertItem -- Tests creating an item. Tests function call from http request to create an item. If an error is thrown at any point during execution, the test fails.

TestUpdateItem -- Tests updating an item from http request. Tests function call from http request on an item with no requested changes. If an error is thrown at any point during execution, the test fails.

TestRemoveItem -- Tests removing an item. Tests function call from http request on an empty item. If an error is thrown at any point during execution, the test fails.

TestFindItem -- Tests finding an item (empty item). Tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails.

API Documentation
User API

makeUser -- creates the user based on the passed in JSON containing user information

getUserWithID -- searches for a user based on the input ID in the database through GORM

getUserWithUsername -- searches for a user based on the input Username in the database through GORM

getUserWithEmail -- searches for a user based on the input Email in the database through GORM

getAllUsers -- returns all of the tuples of all of the users in the database

removeByUserID -- removes a user from the database by searching for the user by ID

removeByUserEmail -- removes a user from the database by searching for the user by Email

removeByUserUsername -- removes a user from the database by searching for the user by Username

updateUserByID -- updates the information of the user by ID

updateUserByUsername -- updates the user information based on the Username and JSON information

Routing API

makeItem -- the function creates a new item tuple in the database through GORM

getItemWithID -- the function searches for the item in the inventory table based on ID and returns it if found

getItemWithName -- the function searches for the item in the inventory table based on the name and returns it if found

getItemWithDate -- the function searches for one or multiple items in inventory table based on acquired date and returns all that are found

getFirstItemWithDate -- the function searches for the first item that has the input date in the inventory table and returns it

getAllItems -- the function returns all of the tuples in the inventory table

removeItemByID -- the function removes the item based on the passed in unique ID in the inventory table

removeItemByName -- the function removes the item based on the passed in Name in the inventory table

updateItemByID -- the function updates the item based on the ID passed through JSON

updateItemByName -- the function updates the item based on Name passed through JSON
