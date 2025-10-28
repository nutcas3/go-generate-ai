#!/bin/bash

# Speed Running REST APIs - Example API Calls
# Make sure the server is running: make run

BASE_URL="http://localhost:8080"

echo "=== Speed Running REST APIs - Demo ==="
echo ""

# Create a user
echo "1. Creating a user..."
curl -X POST $BASE_URL/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Johnson",
    "email": "alice@example.com"
  }' | jq
echo ""

# Create another user
echo "2. Creating another user..."
curl -X POST $BASE_URL/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bob Smith",
    "email": "bob@company.com"
  }' | jq
echo ""

# Get user by ID
echo "3. Getting user by ID (1)..."
curl $BASE_URL/users/1 | jq
echo ""

# List all users
echo "4. Listing all users..."
curl "$BASE_URL/users?limit=10&offset=0" | jq
echo ""

# Update a user
echo "5. Updating user 1..."
curl -X PUT $BASE_URL/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Johnson-Smith",
    "email": "alice.smith@example.com"
  }' | jq
echo ""

# Try to create duplicate email
echo "6. Trying to create user with duplicate email (should fail)..."
curl -X POST $BASE_URL/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Charlie Brown",
    "email": "alice.smith@example.com"
  }' | jq
echo ""

# Get non-existent user
echo "7. Getting non-existent user (should return 404)..."
curl -i $BASE_URL/users/999 2>/dev/null | head -n 1
echo ""

# Create user with invalid input
echo "8. Creating user with invalid input (should fail)..."
curl -X POST $BASE_URL/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "email": "invalid"
  }' | jq
echo ""

# List users with pagination
echo "9. Listing users with pagination (limit=1, offset=0)..."
curl "$BASE_URL/users?limit=1&offset=0" | jq
echo ""

# Delete a user
echo "10. Deleting user 2..."
curl -X DELETE $BASE_URL/users/2 -i 2>/dev/null | head -n 1
echo ""

# Verify deletion
echo "11. Verifying user 2 is deleted..."
curl -i $BASE_URL/users/2 2>/dev/null | head -n 1
echo ""

echo "=== Demo Complete ==="
