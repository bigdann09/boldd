#!/bin/bash

# WorldBanc Customer ID Generator
# Generates unique customer IDs in the format WB###

CUSTOMER_FILE="../private/customers/customer_details.csv"

# Function to get the last customer ID
get_last_id() {
    if [ ! -f "$CUSTOMER_FILE" ]; then
        echo "0"
        return
    fi
    
    # Get the last customer ID from the file
    LAST_ID=$(tail -n 1 "$CUSTOMER_FILE" | cut -d',' -f1)
    
    # Extract the numeric part
    if [[ $LAST_ID =~ WB([0-9]+) ]]; then
        echo "${BASH_REMATCH[1]}"
    else
        echo "0"
    fi
}

# Function to generate next ID
generate_next_id() {
    LAST_NUM=$(get_last_id)
    NEXT_NUM=$((LAST_NUM + 1))
    
    # Format with leading zeros (3 digits)
    printf "WB%03d" $NEXT_NUM
}

# Main script
echo "WorldBanc Customer ID Generator"
echo "==============================="
echo ""

if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
    echo "Usage: $0 [COUNT]"
    echo ""
    echo "Generate new customer IDs"
    echo ""
    echo "Arguments:"
    echo "  COUNT  - Number of IDs to generate (default: 1)"
    echo ""
    echo "Examples:"
    echo "  $0        # Generate 1 ID"
    echo "  $0 5      # Generate 5 IDs"
    exit 0
fi

COUNT=${1:-1}

if ! [[ "$COUNT" =~ ^[0-9]+$ ]]; then
    echo "Error: COUNT must be a positive number"
    exit 1
fi

echo "Generating $COUNT customer ID(s)..."
echo ""

CURRENT_NUM=$(get_last_id)

for ((i=1; i<=COUNT; i++)); do
    NEXT_NUM=$((CURRENT_NUM + i))
    NEW_ID=$(printf "WB%03d" $NEXT_NUM)
    echo "Generated ID: $NEW_ID"
done

echo ""
echo "ID generation complete."
echo ""
echo "Note: These IDs have been generated but not yet assigned."
echo "To register a new customer, add their details to:"
echo "$CUSTOMER_FILE"
