
# Bed and Breakfast Application
  This is a bookings and reservation web application for a bed and breakfast establishment. It is built using the Go programming language.
  
## Link to the website
- In order to access the website directly, visit http://ec2-54-176-175-249.us-west-1.compute.amazonaws.com:8089

## Features
-  __User Authentication__: Users are required to log in with a username before they can access the application.
- __Make Reservation__: Users can make a reservation from a specified start date to end date.
- __Check Availability__: Before they make a reservation, users can enter in a start to end date range to check availability.
- __Admin Panel__: The application provides property owners (admin) with a dashboard to manage user reservations.
- __Email__: The application uses a SMTP mail server called Mailhog to permit property owners to send emails to their clients.
- __Logout__: Users can log out from the application.

## Tech Stack
The application incorporates the following technologies:

- __Go__: A programming language used for building the backend server and handling business logic.
- __HTML/CSS__: Implement frontend webpages.
- __AWS EC2__: Cloud platform used to host web application.
- __Docker__: Containerization technology used to automate deployment of application.
- 
## Setup Instructions
To run the application locally, follow these steps:

1. __Clone the repository__
2. __Navigate to the project directory__ 
3. __Install Go and dependencies__
4. __Install PostgreSQL and create a database called bookings__
5. __Install soda__
   - Instructions to install: https://github.com/elsevierlabs-os/soda/blob/master/docs/installation.md
6.  __Generate the database schema using migration files__
   - Run the command: "soda migrate"
7.  __Build the project__
   - Run the command: "go build -o bookings cmd/web/*.go"
8. __Run the program__
   - Run the shell script called run.sh by typing in "./run.sh"
9. __Open your web browser__
   - Visit ```http://localhost:8080``` to access the application.












