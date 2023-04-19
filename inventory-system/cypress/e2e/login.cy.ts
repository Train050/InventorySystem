describe('Login Test', () => {
  it('Logs in and redirects to inventory page', () => {
    cy.visit('http://localhost:4200/login-page') // Replace with your login page URL

    // Enter username and password and submit the form
    cy.get('user').type('Username')
    cy.get('#password').type('mypassword')
    cy.get('button[type="submit"]').click()

    // Wait for the login to complete and redirect to inventory page
    cy.url().should('http://localhost:4200/inventory-home-page') // Replace with your inventory page URL
  })
})

