# challenge-quiz

## Run It

A docker-compose.yml file is given. It includes a redis-setup for initial loads. The Dockerfile for the application is meant to be only for development/debug/testing purposes, as it supports live reloading.

## Overall Architecture

The project was done by with architecture in mind, with `core` being the innermost layer, including use cases and domain entities. Gateways are the way entrypoints (like controllers) can communicate with the `core` use cases, injecting repositories implementations.
`registry` is where all dependencies are handled and put into place.

## Endpoints

The project is a REST API with the following endpoints:

```
GET /quiz/:quizId
DESCRIPTION Retrieves a quiz and its alternatives. A quiz may have less than, but not more than, 4 alternatives
RESPONSE
{
    "id": string,
    "questions": [{
        "text": string,
        "alternatives": [{
            "1": string,
            "2": string,
            "3": string,
            "4": string
        }]
    }, ...]
}
```

```
GET /quiz/:quizid/ranking/:quizzer
DESCRIPTION Retrieves :quizzer ranking
RESPONSE
{
    "ranking": "You were better than X% of all quizzer"
}
```

```
POST /answer
DESCRIPTION Answer a Quiz; Answers are the keys of the alternatives given, i.e. ["1", "2", ...]
BODY
{
    "quizId": string,
    "answers": []string,
    "respondent": string
}
RESPONSE
{
    "score": "N/M"
}
```
