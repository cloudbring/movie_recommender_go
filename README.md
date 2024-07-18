# Movie Recommender

A personalized movie recommendation system built with Go.

## Overview

This project implements a movie recommendation system based on user-reviewer alignment. It generates personalized movie recommendations by analyzing how often a user's preferences align with movie reviewers' opinions.

## Features

- User registration and login
- Personalized movie recommendations
- Interactive questionnaires after watching movies
- User-reviewer alignment calculation
- Review management

## Technology Stack

- Go programming language
- gorilla/mux for HTTP routing
- gorm for object-relational mapping

## Project Structure

- `main.go`: Entry point of the application
- `recommendation.go`: Handles movie recommendation generation
- `questionnaire.go`: Manages creation and submission of questionnaires
- `user.go`: Handles user-related operations (login, logout, registration)
- `reviewer.go`: Manages reviewer data and operations
- `alignment.go`: Calculates and updates user-reviewer alignment

## Getting Started

1. Clone the repository
2. Install dependencies: `go get -u github.com/gorilla/mux github.com/jinzhu/gorm`
3. Run the application: `go run main.go`

## API Endpoints

- POST `/user/register`: Register a new user
- POST `/user/login`: User login
- GET `/recommendation/{userID}`: Get personalized movie recommendations
- POST `/questionnaire/create`: Create a new questionnaire
- POST `/questionnaire/submit`: Submit a completed questionnaire
- GET `/reviewer/{id}/reviews`: Get reviews from a specific reviewer

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.