
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