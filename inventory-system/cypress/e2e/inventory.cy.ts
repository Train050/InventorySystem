describe('inventory page', () => {
it('passes', () => {
  cy.visit('http://localhost:4200/inventory-home-page')
  cy.get('.mat-mdc-paginator-navigation-next').click()
});

it('should display the table with inventory items', () => {
  cy.visit('http://localhost:4200/inventory-home-page')
  cy.url().should('eq', 'http://localhost:4200/inventory-home-page')
  cy.get('table').should('be.visible')
});
});

describe('add item dialog', () => {
it('should display the add item button', () => {
  cy.visit('http://localhost:4200/inventory-home-page')
  cy.get('.mat-mdc-button').should('be.visible')
});

it('should display the add item dialog', () => {
  cy.visit('http://localhost:4200/inventory-home-page')
  cy.get('.mat-mdc-button').click({ multiple: true })
});
});

describe("Post Method", function(){
  it("Posting to the database", function(){
     cy.request('http://localhost:8080/inventory')
     .its('body') // yields the first element of the returned list
     // make a new post on behalf of the user
     cy.request('POST', 'http://localhost:8080/inventory', {
        productName: 'test',
        dateAcquired: '2021-05-05',
        quantity: 1,
     })
  })
});
