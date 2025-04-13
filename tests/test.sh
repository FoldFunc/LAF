curl -X POST http://localhost:8080/addLostAndFound \
-H "Content-Type: application/json" \
-d '{
    "name": "Lost Wallet",
    "whoFound": "John Doe",
    "where": "Central Park",
    "notes": "Found near the bench"
}'

