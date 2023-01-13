Flight Path Tracker API
This API allows you to track the flight path of a person by providing a list of flights in the form of source and destination airport codes.

Endpoint
POST /calculate

Request
The API expects a JSON body in the following format:

Copy code
[
    {"source": "SFO", "destination": "EWR"},
    {"source": "ATL", "destination": "EWR"},
    {"source": "SFO", "destination": "ATL"},
    {"source": "GSO", "destination": "IND"},
    {"source": "ATL", "destination": "GSO"}
]
Response
The API returns a JSON object in the following format:

Copy code
{
    "path": [
        {"source": "SFO", "destination": "EWR"}
    ]
}
Usage
To use the API, send a POST request to http://localhost:8080/calculate with the list of flights in the request body. The API will return the total flight path in the response.

Example
Copy code
curl -X POST -H "Content-Type: application/json" -d '[{"source": "SFO", "destination": "EWR"},{"source": "ATL", "destination": "EWR"},{"source": "SFO", "destination": "ATL"},{"source": "GSO", "destination": "IND"},{"source": "ATL", "destination": "GSO"}]' http://localhost:8080/calculate
This will return the total flight path [{"source": "SFO", "destination": "EWR"}]

Please make sure that you have the API running on your localhost on port 8080 before making the request.
