# The BalihBank Project

Welcome to the BalihBank Project, a simple banking system written in Go. This project allows users to sign up, log in, and manage their balances through a command-line interface. Below, you'll find an overview of the project's features, how to set it up, and how to use it.

## Features

- **User Signup**: Create a new account with a unique username.
- **User Login**: Log in with your username and password.
- **Deposit and Withdraw**: Manage your account balance by depositing and withdrawing funds.
- **Account Persistence**: User data is stored in local files for persistence.

## Getting Started

### Prerequisites

To run this project, you need to have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/BalihBank.git
   cd BalihBank
   ```

2. Build the project:
   ```sh
   go build -o balihbank main.go
   ```

3. Run the executable:
   ```sh
   ./balihbank
   ```

### Usage

When you run the program, you will be presented with a main menu:

```
~~~~~~~~~~~~~~~~~~~~
#                  #
#   THE BALİHBANK  #
#                  #
~~~~~~~~~~~~~~~~~~~~
 Written in Go 2023 
     By Menesay     
                     
1) Sign Up
2) Login
3) Exit
Option: 
```

#### Sign Up

1. Select the "Sign Up" option by entering `1`.
2. Follow the prompts to enter a unique username and password.

#### Login

1. Select the "Login" option by entering `2`.
2. Enter your username and password.
3. If login is successful, you will be presented with options to deposit, withdraw, or log out.

#### Deposit

1. Select the "Deposit" option by entering `1` after logging in.
2. Enter the amount you wish to deposit.

#### Withdraw

1. Select the "Withdraw" option by entering `2` after logging in.
2. Enter the amount you wish to withdraw.

#### Log Out

1. Select the "Log out" option by entering `3`.

## Project Structure

- **main.go**: The main application file containing all the logic for user interaction and database management.
- **username-pass.db**: A file storing usernames and passwords.
- **username-balance.db**: A file intended for storing usernames and balances (currently not fully implemented).

## To Do

- Implement storing and updating user balances in `username-balance.db`.
- Add more robust error handling and input validation.
- Enhance security measures for storing and handling passwords.

Enjoy using BalihBank! If you encounter any issues or have suggestions for improvement, feel free to open an issue or submit a pull request.
