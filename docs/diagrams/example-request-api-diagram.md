sequenceDiagram
participant Client
participant NewsAndTopicService
participant Database

    Client->>NewsAndTopicService: GET /news?status=published&topic=sports
    NewsAndTopicService->>Database: Query news where status='published'
    Database-->>NewsAndTopicService: Return news articles
    NewsAndTopicService->>Database: Query topics for news IDs
    Database-->>NewsAndTopicService: Return topics
    NewsAndTopicService-->>Client: Response with news articles and topics
