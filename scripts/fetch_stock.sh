
#!/bin/bash

BASE_URL="http://localhost:8000/api/stocks"
COMPANY=$1

if [ -z "$COMPANY" ]; then
    echo "Usage: $0 <company>"
    exit 1
fi

curl "$BASE_URL/$COMPANY"
