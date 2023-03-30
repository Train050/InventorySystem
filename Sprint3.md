In sprint 3, our group was able further expand and improve upon the functionality of our application. In the back end, we have increased the documentation of routing so that it is easier for the front end to use them. In addition, we have created a seeding program for users and inventory that generates data to populate the database. We introduced an authorization feature that encrypts the password of users and checks their authorization token before allowing them to get data from the database. Authorization has not been fully implemented into the routing yet but the functions to check the hashed password and token have been created. Additional test cases were created for the previously described features, ensuring that all functions work as intended. Regarding the front end, we were able to figure out how to send user info from registration to the backend as a json object, which can be encrypted with the authorization feature. We also touched up the landing page a bit and made a template that we can expand off of in the future. 

The inventory page Cypress test expects to show the inventory page and that it is directed to 'http://localhost:4200/inventory-home-page' and for the paginator to be working. 
The inventory page unit test tests the creation of the components in the inventory page. 

The login page Cypress test expects for the 'Have an account button' to redirect the user to the sign in page from the register page. 
The unit test for the Register function shold make an http post request with the user info and be redirected to the user home page. 
The unit test for the Login function should make an http post request with the login info to authenitcate the user and redirect them to the user home page.
The login page Cypress test expects for the 'Have an account button' to redirect the user to the sign in page from the register page. 
The unit test for the Register function shold make an http post request with the user info and be redirected to the user home page. 
The unit test for the Login function should make an http post request with the login info to authenitcate the user and redirect them to the user home page.
Unit tests for backend: 

TestMakeUser -- Tests creating a user. First tests function call from http request on an empty user with ID = 1, then, manually tests database entry for a fully populated sample user. If an error is thrown at any point during execution, the test fails. If the the final user in the database does not exactly equal the desired test user, the test fails, else, the test passes. This method is not mirrored in future tests, since it will be better to test actual queries from the client in the next sprint.

TestRemoveItem -- Tests removing an item. Tests function call from http request

TestFindItem -- Tests finding an item (empty item). Tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails.

TestRemoveUser -- Tests removing a user. Tests function call from http request. If an error is thrown at any point during execution, the test fails.

TestFindUser -- Tests find user function. Tests function call from http request on user from ID. If an error is thrown at any point during execution, the test fails.

TestInsertItem -- Tests creating an item. Tests function call from http request to create an item. If an error is thrown at any point during execution, the test fails.

TestUpdateItem -- Tests updating an item from http request. Tests function call from http request on an item with no requested changes. If an error is thrown at any point during execution, the test fails.

TestRemoveItem -- Tests removing an item. Tests function call from http request on an empty item. If an error is thrown at any point during execution, the test fails.

TestFindItem -- Tests finding an item (empty item). Tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails.

TestGetAllItems -- Tests getting all items. Tests function call via testing database from http request. If an error is thrown at any point during execution, the test fails.

TestGetAllUsers -- Tests getting all users. Tests function call via testing database from http request. If an error is thrown at any point during execution, the test fails. 

TestInventorySeeder -- Tests inventory seeding function for product names via mocked database (array). Tests that all elements of array are populated with item names, if not, the test fails.

TestUserSeeder -- Tests user seeding function for product names via mocked database (array). Tests that all elements of array are populated with usernames, if not, the test fails.

TestRemoveUser -- Tests removing a user. Tests function call from http request. If an error is thrown at any point during execution, the test fails.

TestFindUser -- Tests find user function. Tests function call from http request on user from ID. If an error is thrown at any point during execution, the test fails.

TestInsertItem -- Tests creating an item. Tests function call from http request to create an item. If an error is thrown at any point during execution, the test fails.

TestUpdateItem -- Tests updating an item from http request. Tests function call from http request on an item with no requested changes. If an error is thrown at any point during execution, the test fails.

TestRemoveItem -- Tests removing an item. Tests function call from http request on an empty item. If an error is thrown at any point during execution, the test fails.

TestFindItem -- Tests finding an item (empty item). Tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails.

TestGetAllItems -- Tests getting all items. Tests function call via testing database from http request. If an error is thrown at any point during execution, the test fails.

TestGetAllUsers -- Tests getting all users. Tests function call via testing database from http request. If an error is thrown at any point during execution, the test fails. 

TestInventorySeeder -- Tests inventory seeding function for product names via mocked database (array). Tests that all elements of array are populated with item names, if not, the test fails.

TestUserSeeder -- Tests user seeding function for product names via mocked database (array). Tests that all elements of array are populated with usernames, if not, the test fails.

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

userSeeder -- Mock users can be generated and inserted into the database based on the requested number of entries

Inventory API

makeItem -- the function creates a new item tuple in the database through GORM

getItemWithID -- the function searches for the item in the inventory table based on ID and returns it if found

getItemWithName -- the function searches for the item in the inventory table based on the name and returns it if found

getItemWithDate -- the function searches for one or multiple items in inventory table based on acquired date and returns all that are found

getFirstItemWithDate -- the function searches for the first item that has the input date in the inventory table and returns it

getAllItems -- the function returns all of the tuples in the inventory table

removeItemByID -- the function removes the item based on the passed in unique ID in the inventory table

removeItemByName -- the function removes the item based on the passed in Name in the inventory table

updateItemByID -- the function updates the item based on the ID passed through JSON ojbect

updateItemByName -- the function updates the item based on Name passed through JSON object

inventorySeeder -- Mock items can be generated and inserted into the database based on the requested number of entries

Routing API

router.HandleFunc("/login/{ID}", getUserWithID).Methods("GET") -- returns a json object of a user based on input ID
router.HandleFunc("/login/{Username}", getUserWithUsername).Methods("GET") -- retrieves a user json object from their input Username
router.HandleFunc("/login/{Email}", getUserWithEmail).Methods("GET") -- Finds and returns the user json object based on their email
router.HandleFunc("/login", getAllUsers).Methods("GET") -- Returns all json user objects within the user database

router.HandleFunc("/login/{ID}", updateUserById).Methods("PUT") -- Register for updating the user ID for an existing user
router.HandleFunc("/login/{Username}", updateUserByUsername).Methods("PUT") -- Register for updating an existing user Username

router.HandleFunc("/login/{ID}", removeUserByID).Methods("DELETE") -- Allows the finding and deletion of a user based on their ID
router.HandleFunc("/login/{Username}", removeUserByUsername).Methods("DELETE") -- Allows a user to be deleted by their Username
router.HandleFunc("/login/{Email}", removeUserByEmail).Methods("DELETE") -- A user can be deleted through their Email in this route

router.HandleFunc("/registration", makeUser).Methods("POST") -- Allows the creation of a new user within the user database

router.HandleFunc("/inventory", makeItem).Methods("POST") -- Allows the creation of a new Item within the inventory database

router.HandleFunc("/inventory/{ID}", getItemWithID).Methods("GET") -- returns an item based on its ID passed in
router.HandleFunc("/inventory/{ProductName}", getItemWithName).Methods("GET") -- locates and returns an item based on input product name
router.HandleFunc("/inventory/{DateAcquired}", getFirstItemWithDate).Methods("GET") -- finds the first item matching the input date
router.HandleFunc("/inventory/{DateAcquired}", getItemsWithDate).Methods("GET") -- finds all the items matching the input date
router.HandleFunc("/inventory", getAllItems).Methods("GET") -- returns all of the items within the inventory database

router.HandleFunc("/inventory/{ID}", updateItemById).Methods("PUT") -- allows for an item to be updated based on the input ID
router.HandleFunc("/inventory/{ProductName}", updateItemByName).Methods("PUT") -- through an inserted product name, an item info can be fixed

router.HandleFunc("/inventory/{ID}", removeItemByID).Methods("DELETE") -- based on a passed in ID, an item can be removed in the database
router.HandleFunc("/inventory/{ProductName}", removeItemByName).Methods("DELETE") -- Through the product name, an item can be deleted

Front End testing

1. AddItemComponent -- testing component template logic
2. Registration -- test sending the http post request to the backend with the user registration info for encryption and redirect to user homepage.
3. Login -- test sending the http post request to the backend with user login info for authentication and redirect to user homepage if access is granted.
4. LoginF -- tests to make sure that the login does not authenticate the user when they input the wrong information (should fail the test).