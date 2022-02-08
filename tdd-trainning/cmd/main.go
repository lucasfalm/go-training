package main

import "fmt"

/* Types of testes
---------------------------

1) Unit tests -> just test the components logic, mocking the functions/methods called inside block
As much information as possible about where the failure happened

e.g: Create new user

Unit test on the controller
Unit test on the service
Unit test on the dto and dao
Integration test between controller and service
Functional test between all (POST /new-user with json, and test the result)


---------------------------

2) Integration tests -> test the integration (interaction) between all the components, not mocking the functions/methods, make sure all diferent layers, artefacts and components are all integrated well


---------------------------

3) Functional Test -> is needed to start and run the application (make requests against the application and check the result), eg: test a POST request against the route, then check if the response is what we expect
We need all running, the database, the routes, the services.

---------------------------

Pyramid of tests:
Functional
Integration
Unit

---------------------------

Concepts:
White Box
Access to private components (such as varriables, methods, functions, and so on)
(e.g our test is in the same package)

Black Box
Access to public components

---------------------------

Fases of tests:
Init - setup
Execution - call the method
Validation - check if the result is the same as expected
Tear Down - clear

---------------------------

Each clausule return from a method or function in production code must have an test case
*/

func main() {
	something := saySomethingNow()
	fmt.Println(something)
}

func saySomethingNow() string {
	return "hey"
}
