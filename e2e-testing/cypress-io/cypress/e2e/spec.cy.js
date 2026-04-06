describe(
  "TEEEST",
  () => {
    it("passes", () => {
      cy.viewport(2560, 1440);
      cy.visit("http://web.nextreleaseplease.com/api/v1/auth/e2e-test-token");
    });
  }
)