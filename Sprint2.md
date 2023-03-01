Work completed in Sprint 2: 

In sprint 2, our group was able to improve the functionality of both the frontend and backend. Regarding the frontend development, we were able to develop a filter function to find specific data entries. In addition, we were able to get the table size changing feature working so rows shown can be adjusted. There is now an add button for "Add Item" on the inventory page and currently working on adding forms to that pop up. For the backend, we redesigned our initial code to be more simplified. We managed to get the SQlite database working with our tables, inventory and users, allowing for the creation of tuples. We were able to get GORM to communicate to the database and allow for the insertion, deletion, finding of tuples, and updating of the tuples within the tables. We also started integration of communication between backend and frontend. We established the routes that will use JSON to send information between frontend and backend. We lastly updated our code to include documentation for the use of each functions and their appropriate routes. Both frontend and backend also created several test cases to determine if the applications are working as expected.


Unit tests and Cypress tests for frontend:
    1. Inventory cypress:  
    describe('expect to show inventory page', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/inventory-home-page')
    cy.get('.mat-mdc-paginator-navigation-next').click()
  })
})

    2. Login Cypress:
    
    describe('Button position test', () => {
    it('Button should change position when another button is clicked', () => {
      cy.visit('http://localhost:53597/login-page');
      cy.get('.login__submit').eq(1).then(($button) => {
        const initialPosition = $button.position();
        cy.get('.login__submit').eq(1).click();
        cy.wait(1000); // wait for 1 second for the position change to happen
        cy.get('.login__submit').eq(1).then(($button) => {
          const newPosition = $button.position();
          expect(newPosition).to.not.equal(initialPosition);
        });
      });
    });
  });
    
    3. Inventory unit tests:  
     it('should create', () => {
    expect(component).toBeTruthy();
    expect(component.dataSource).toBeTruthy();
    expect(component.paginator).toBeInstanceOf(MatPaginator);
    expect(component.applyFilter).toBeTruthy();
    expect(component).toBeDefined();
    expect(component.openAddItem).toBeTruthy();
    expect(component.openAddItem).toBeInstanceOf(InventoryHomePageComponent);
  });

Unit tests for backend: 

TestMakeUser -- Tests creating a user. First tests function call from http request on an empty user with ID = 1, then, manually tests database entry for a fully populated sample user. If an error is thrown at any point during execution, the test fails. If the the final user in the database does not exactly equal the desired test user, the test fails, else, the test passes. This method is not mirrored in future tests, since it will be better to test actuall queries from the client in the next sprint.

TestUpdateUser -- Tests updating a user. Firsts tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails. Does not test actual update / alteration of data since that process is tied into the client side requests in our code, will be more efficient to test later.

TestRemoveUser -- Tests removing a user. Tests function call from http request. If an error is thrown at any point during execution, the test fails.

TestFindUser -- Tests find user function. Tests function call from http request on user from ID. If an error is thrown at any point during execution, the test fails.

TestInsertItem -- Tests creating an item. Tests function call from http request to create an item. If an error is thrown at any point during execution, the test fails.

TestUpdateItem -- Tests updating an item from http request. Tests function call from http request on an item with no requested changes. If an error is thrown at any point during execution, the test fails.

TestRemoveItem -- Tests removing an item. Tests function call from http request on an empty item. If an error is thrown at any point during execution, the test fails.

TestFindItem -- Tests finding an item (empty item). Tests function call from http request on a user with no requested changes. If an error is thrown at any point during execution, the test fails.

