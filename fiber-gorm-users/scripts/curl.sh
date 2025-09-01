#!/usr/bin/env bash
set -euo pipefail

BASE="http://localhost:8085"

echo "== Create user =="
curl -sS -X POST "$BASE/users" -H "Content-Type: application/json" -d '{
  "name":"Ava Shah", "email":"ava@example.com", "password":"secret123", "status":"active"
}' | jq

echo "== Create another user =="
curl -sS -X POST "$BASE/users" -H "Content-Type: application/json" -d '{
  "name":"Rahul Iyer", "email":"rahul@example.com", "password":"password", "status":"active"
}' | jq

echo "== List users (limit=5, skip=0) =="
curl -sS "$BASE/users?limit=5&skip=0" | jq

echo "== Get by ID 1 =="
curl -sS "$BASE/users/1" | jq

echo "== Update ID 1 (block) =="
curl -sS -X PUT "$BASE/users/1" -H "Content-Type: application/json" -d '{
  "status":"blocked"
}' | jq

echo "== Delete ID 2 =="
curl -sS -X DELETE "$BASE/users/2" | jq
