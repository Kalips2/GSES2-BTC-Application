# GSES2 test task
My test task to GSES 2.

## Description:
The Bitcoin Rate Service allows users to fetch the current exchange rate of Bitcoin in UAH and subscribe to receive updates the rate changes. It provides a RESTful API for interacting with the service.
# Use cases:
* Get current exchange rate of Bitcoin to UAH.
* Subscription to receive email notifications about the exchange rate of Bitcoin to UAH.
* Send emails with corresponding rates to subscribed emails.

## URL paths:
```
GET  -> http://localhost:8080/rate
POST -> http://localhost:8080/subscribe             
POST -> http://localhost:8080/sendEmails
```
## Getting started
* Run the project using Docker Compose:
```
docker compose up
```
* Run the project explicitly with Docker:
```
docker build -t btc-app .
docker run -p 8080:8080 btc-app
```
## Architecture
* **Handlers**  - handle HTTP requests, send the required parameters to the appropriate service, and return a response to the user. Define the API endpoints.
* **Services** - Contain the core business logic of the application, such as fetching the current rate and managing subscriptions.
* **Repositories** - Provide data persistence for subscriptions. In this implementation, a CSV file is used as a storage mechanism.
* **Utils** -  Include utility functions and helpers, such as email validation and constants.

For getting current rate API from https://min-api.cryptocompare.com/ is used.
If the type of currency is wanted to be changed you get replace default currency (BTC, UAH) in the constants file to what you need.


 
