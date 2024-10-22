# Secret Sharing

This project implements secret sharing schemes, including Shamir's Secret Sharing (SSS) and Verifiable Secret Sharing (VSS). The algorithms allow a secret to be divided into parts, where a certain number of parts are required to reconstruct the original secret.

## Features

-   Shamir's Secret Sharing:

    -   Divides a secret into multiple shares.
    -   Any specified number of shares (threshold) can be used to reconstruct the secret.

-   Verifiable Secret Sharing:
    -   Extends Shamir's scheme to allow verification of shares.
    -   Ensures integrity and authenticity of the shares before reconstruction.

## Prerequisites

-   Go 1.16 or higher
-   A working Go environment set up on your machine

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/YourUsername/secret-sharing.git
    cd secret-sharing
    ```

2. Install dependencies (if any):
    ```bash
    git mod tidy
    ```

## Running the Application

-   To build the application for your current platform, run:

    ```bash
     make build
    ```

-   Run the Application
    To run the application directly, you can use:

    ```bash
     make run
    ```

-   Testing
    To run the tests for the secret sharing schemes, use:

    ```bash
    make test
    ```

-   Clean Build Artifacts
    To clean up any built binaries, use:

    ```bash
    make clean
    ```
