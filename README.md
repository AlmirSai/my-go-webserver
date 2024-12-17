# My Go Web Server

This is a simple web server application written in Go. It is designed to be used in production environments with Docker and Docker Compose for deployment. The project includes automated CI/CD workflows with GitHub Actions for testing, linting, documentation generation, and deployment.

## Project Structure

```plaintext
my-go-webserver/
├── .github/
│   ├── workflows/
│   │   ├── build.yml
│   │   ├── ci.yml
│   │   ├── cd.yml
│   │   ├── release.yml
│   │   ├── validate.yml
│   │   ├── notify.yml
│   │   ├── docs.yml
│   │   └── monitoring.yml
├── cmd/
│   └── webserver/
│       └── main.go          # Main entry point for the application
├── internal/
│   ├── handler/
│   │   └── handler.go       # HTTP handlers
│   ├── service/
│   │   └── service.go       # Business logic
│   └── storage/
│       └── storage.go       # Data storage interaction
├── pkg/
│   └── config/
│       └── config.go        # Configuration management
├── Dockerfile
├── docker-compose.yml
├── Makefile                 # Makefile for common tasks (build, test, docker, etc.)
├── README.md                # This file
├── go.mod
└── go.sum
```

## Requirements

- Go 1.23.3 or higher
- Docker
- Docker Compose (for local development and testing)

## Setup

### 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/my-go-webserver.git
cd my-go-webserver
```

### 2. Install Go Dependencies

Install all Go dependencies:

```bash
make install-deps
```

### 3. Run the Application Locally (Docker)

You can run the application using Docker Compose:

```bash
make docker-compose-up
```

This will build the Docker image and run it in a container, exposing port 8080.

### 4. Running Tests

You can run tests locally with:

```bash
make test
```

### 5. Lint the Code

To lint the Go code using `golangci-lint`, run:

```bash
make lint
```

### 6. Build the Docker Image

To manually build the Docker image:

```bash
make docker-build
```

### 7. Deploy the Application

To deploy the application, push the Docker image to your Docker registry and follow the steps in the `CI/CD` pipeline.

## CI/CD Workflows

This project includes several GitHub Actions workflows for automating the process of testing, linting, deploying, and more.

- **`build.yml`**: Builds the Go application.
- **`ci.yml`**: Runs tests and lint checks on the `develop` branch.
- **`cd.yml`**: Handles continuous deployment to production when changes are merged into `main`.
- **`release.yml`**: Creates a new GitHub release when a tag is pushed.
- **`validate.yml`**: Runs validation checks on the `develop` branch.
- **`notify.yml`**: Sends notifications on `main` and `develop` branches.
- **`docs.yml`**: Generates project documentation.
- **`monitoring.yml`**: Runs monitoring tasks.

Each workflow is triggered automatically based on branch and tag events.

## Deployment

To deploy the application, follow these steps:

1. Push your changes to the `develop` branch.
2. Once the code is validated and tested, merge into the `main` branch.
3. The deployment will be automatically triggered via the `cd.yml` GitHub Actions workflow, which will push the Docker image to your registry and deploy to production.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to fork this repository, open issues, and submit pull requests. Contributions are welcome!

---

### Notes

- Make sure to replace `your-username` with your actual GitHub username in the `git clone` URL.
- Update the `Dockerfile`, `docker-compose.yml`, and other configurations according to your deployment environment.

