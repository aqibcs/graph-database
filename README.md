# Basketball Stats Neo4j Database

This Docker Compose configuration sets up a Neo4j database for storing basketball player, coach, and team statistics.

## Prerequisites

Before you start, make sure you have Docker Compose installed on your system.

- [Docker Compose Installation](https://docs.docker.com/compose/install/)

## Usage

1. Run the following command to start the Neo4j container:

    ```bash
    docker-compose up -d
    ```

2. Access the Neo4j Browser by opening [http://localhost:7474](http://localhost:7474) in your web browser. Log in with the credentials specified in the `NEO4J_AUTH` environment variable.

## Configuration

- The Neo4j image used is the latest version available on Docker Hub.
- The container is named "neo4j."
- Ports 7474 and 7687 are exposed for Neo4j access.
- The default authentication credentials are `neo4j/mypassword`. Make sure to change the password for production use.

## Cleanup

To stop and remove the Neo4j container, run:

```bash
docker-compose down
