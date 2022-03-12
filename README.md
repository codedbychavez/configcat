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

