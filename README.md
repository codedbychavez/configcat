# ConfigCat

## Paint Buddy API - v1.0

### Core Features
1. Returns the total liters of paint needed to cover a square or rectangular space given the **width** and **height**.

2. Returns the mix result of two colors.

### Beta Features (Feature flags)
1. Returns a random color - This can be interpreted as a random color we recommend for your next paint project.
2. Returns a list of paint companies - This can be interpreted as paint companies that chooses to advertise with us. and we can suggest those to users.

## API Endpoints

### Core features

#### Total liters of paint.
- Endpoint: `http://127.0.0.1:3000/api/v1/paint/liters`
- Method: `POST`
- Content-Type: `application/json`
- Body:

        {
            "width": 1.5,
            "height": 3.0
        }


#### Total liters of paint.
- Endpoint: `http://127.0.0.1:3000/api/v1/paint/mix`
- Method: `POST`
- Content-Type: `application/json`
- Body:

        {
            "color1": "blue",
            "height": "yellow"
        }


### Beta features (Feature Flags)

#### Return a random color.
- Endpoint: `http://127.0.0.1:3000/api/v1/paint/random-color`
- Method: `GET`

- **Configcat**  
Feature Flag: "randomColor"

#### Return a list of paint companies.
- Endpoint: `http://127.0.0.1:3000/api/v1/paint/companies`
- Method: `GET`

- **Configcat**  
Feature Flag: "companies"

## Setup 

### Without Docker
1. Clone this repo
2. Open this repo in your terminal
3. Run `make env` to create a `.env` file or `cp .env.example .env`
4. Add your ConfigCat API key to the `.env` file
5. Install go packages, run `make install` or `go mod download`
5. Run `make run` to start the app or `go run .`

### With Docker
1. Clone this repo
2. Open this repo in your terminal
3. Run `make env` to create a `.env` file or `cp .env.example .env`
4. Add your ConfigCat API key to the `.env` file
5. Run `make docker` or `docker-compose up --build`

## Interacting with the API (Using Postman)
1. import the file `ConfigCat.postman_collection.json` into postman