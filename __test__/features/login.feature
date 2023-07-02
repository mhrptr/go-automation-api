@login

Feature: Login API
  As a user
  I want to be able to login to the application
  So that I can access my account

  Scenario: Successful login
    Given base url with endpoint "/login/email"
    When hit POST request with the following data:
      """
      {
        "username": "myusername",
        "password": "mypassword"
      }
      """
    Then validate POST response