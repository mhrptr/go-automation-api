@login

Feature: Login API
  As a user
  I want to be able to login to the application
  So that I can access my account

  Scenario: Successful login
    Given base url with endpoint "/login/email"
    When hit POST request with email "muharram.samsuddin@kompas.com" and password "kompas1234"
    Then validate POST response