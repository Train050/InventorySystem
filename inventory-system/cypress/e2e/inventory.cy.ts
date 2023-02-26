describe('expect to show inventory page', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/inventory-home-page')
    cy.get('.mat-mdc-paginator-navigation-next').click()
  })
})
