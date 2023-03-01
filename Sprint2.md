Work completed in Sprint 2: 
<<<<<<< HEAD
In sprint 2, our group was able to improve in both the frontend and backend. Regarding the frontend development, we were able to develop a filter function to find specific data entries. We were able to get the table size changing feature working so number of rows shown can be adjusted. In addition, we were able to get a filter click function on the table to work. There is now an add button for "Add Item" on the inventory page and currently working on adding forms to that pop up. For the backend, we redesigned our initial code to be more simplified. We managed to get the SQlite database working with tables inventory and users being created. We were able to get GORM to communicate to the database and allow for the insertion, deletion, and updating of the tuples within the tables. We also started integration of communication between backend and frontend.
=======
In sprint 2, our group was able to improve the functionality of both the frontend and backend. Regarding the frontend development, we were able to develop a filter function to find specific data entries. In addition, we were able to get the table size changing feature working so rows shown can be adjusted. For the backend, we redesigned our initial code to be more simplified. We managed to get the SQlite database working with our tables, inventory and users, allowing for the creation of tuples. We were able to get GORM to communicate to the database and allow for the insertion, deletion, finding of tuples, and updating of the tuples within the tables. We also started integration of communication between backend and frontend. We established the routes that will use JSON to send information between frontend and backend. We lastly updated our code to include documentation for the use of each functions and their appropriate routes. Both frontend and backend also created several test cases to determine if the applications are working as expected.
>>>>>>> d0f1ce639ff18c9299df83294476478ca61d16db

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

