#!/bin/bash

echo -e "\n=== ADD LOST & FOUND ==="
curl -s -X POST http://localhost:8080/addLostAndFound \
-H "Content-Type: application/json" \
-d '{
  "name": "Lost Wallet",
  "whoFound": "John Doe",
  "where": "Central Park",
  "notes": "Found near the bench"
}' | jq .

echo -e "\n=== GET LOST & FOUND ==="
curl -s -X POST http://localhost:8080/getLostAndFound \
-H "Content-Type: application/json" \
-d '{
  "name": "Lost Wallet",
  "whoFound": "John Doe",
  "where": "Central Park"
}' | jq .

echo -e "\n=== CLAIM LOST & FOUND ==="
curl -s -X POST http://localhost:8080/claimLostAndFound \
-H "Content-Type: application/json" \
-d '{
  "name": "Lost Wallet",
  "claimingName": "Alice Smith",
  "dateClaimed": "2025-04-13"
}' | jq .

echo -e "\n=== SHOW ALL LOST & FOUND ==="
curl -s http://localhost:8080/showAllLostAndFound | jq .

