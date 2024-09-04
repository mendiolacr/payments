# Online Payment Platform 💳

Payment platform API and bank simulator.

## Installation 🔧

You have two option:
### 1 Option
1. **Download the Project**
   - Clone this repository to your local machine:
     ```bash
     git clone https://github.com/mendiolacr/payments.git
     ```

2. **Configure the Database**
   - Create a database in your preferred SQL management tool.
   - Update the database credentials in the project (db.go).

3. **Build and Run with Docker**
   - In the root directory of the project, run the following command to build and start the containers:
     ```bash
     docker-compose up --build
     ```
   - This will download the necessary images and start the containers for both the application and the database.

### 2 Option
**Docker Repository**

For download payment platform run
 ```bash
docker pull mendiolacr/payment_platform:latest
 ```

For download simulator run
 ```bash
docker pull mendiolacr/bank_simulator:latest
 ```

Download the two services in the same folder and add docker-compose.yml file and run 

     ```bash
     docker-compose up --build
     ```

## Running the Solution 🚀

**Access the API**
   - The payment platform API will be available at `http://localhost:8080` .
   - The bank simulator API will be available at `http://bank_simulator:8081` .
   - You can test the endpoints using tools like Postman.
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/5205075-a0612109-f5a2-4919-aa48-37cca8af6b78?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D5205075-a0612109-f5a2-4919-aa48-37cca8af6b78%26entityType%3Dcollection%26workspaceId%3Df326ae33-2888-45af-b1fe-04ddc953c1e1)

### Dependencies and Prerequisites

- **Docker**: Ensure Docker and Docker Compose are installed on your machine.
- **SQL Database**: You'll need a compatible SQL management tool..

## Assumptions 🤔

During the design and development of the platform, the following assumptions were made:

- The database would be accessible from the Docker container.
- The development environment would be compatible with the Docker and Docker Compose versions used.
- SQL and transactions were chosen for payment processing, as they offer robust data management and consistency for financial transactions.

## Areas for Improvement 🔍

Here are some potential areas for improving the platform:

- **Scalability**: Implement horizontal scalability solutions, such as Kubernetes, to handle higher transaction volumes.
- **Security**: Integrate advanced security mechanisms, like authentication and data encryption.
- **Golang**: Explore optimizations in Golang, such as performance tuning and better concurrency management, to enhance the efficiency and speed of the application.

## Cloud Technologies ☁️

- **Docker**: Used for containerization and orchestration of the development environment, ensuring consistency across all stages of the software lifecycle.

**Justification**: Docker simplifies deployment and dependency management, making the application more portable and easier to maintain.

## Built With 🛠️

- **Golang**: The primary programming language for API development.
- **SQL**: Database management for storing and retrieving transaction data.
- **Docker**: Containerization of the environment to facilitate installation and execution.

---

⌨️ with ❤️ by MendiolaCR 😊
