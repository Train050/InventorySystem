describe('redirect to inventory page from registration page', () => {
  it('should redirect to inventory page', () => {
    cy.visit('http://localhost:4200/register')
    cy.get('.mat-mdc-button').click({ multiple: true})
    cy.visit('http://localhost:4200/inventory-home-page')
  });
});

