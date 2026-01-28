#!/bin/bash

# WorldBanc Transaction Filter Script
# Usage: ./filter_transactions.sh [TYPE] [CUSTOMER_ID]
# Example: ./filter_transactions.sh DEPOSIT WB001

TRANSACTION_FILE="../private/transactions/transactions_2026_01.csv"

if [ ! -f "$TRANSACTION_FILE" ]; then
    echo "Error: Transaction file not found!"
    exit 1
fi

# Function to display usage
usage() {
    echo "Usage: $0 [TYPE] [CUSTOMER_ID]"
    echo ""
    echo "Filter transactions by type and/or customer ID"
    echo ""
    echo "Transaction Types:"
    echo "  DEPOSIT    - Show all deposits"
    echo "  WITHDRAWAL - Show all withdrawals"
    echo "  PAYMENT    - Show all payments"
    echo "  TRANSFER   - Show all transfers"
    echo ""
    echo "Examples:"
    echo "  $0 DEPOSIT              # Show all deposits"
    echo "  $0 DEPOSIT WB001        # Show deposits for customer WB001"
    echo "  $0 \"\" WB001            # Show all transactions for customer WB001"
    exit 1
}

# Check if help is requested
if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
    usage
fi

TYPE="$1"
CUSTOMER_ID="$2"

echo "WorldBanc Transaction Filter"
echo "============================"
echo ""

# Filter based on parameters
if [ -z "$TYPE" ] && [ -z "$CUSTOMER_ID" ]; then
    # No filters - show all transactions
    echo "Showing all transactions:"
    echo ""
    cat "$TRANSACTION_FILE"
elif [ -n "$TYPE" ] && [ -z "$CUSTOMER_ID" ]; then
    # Filter by type only
    echo "Filtering by type: $TYPE"
    echo ""
    head -n 1 "$TRANSACTION_FILE"
    grep ",$TYPE," "$TRANSACTION_FILE"
elif [ -z "$TYPE" ] && [ -n "$CUSTOMER_ID" ]; then
    # Filter by customer ID only
    echo "Filtering by customer: $CUSTOMER_ID"
    echo ""
    head -n 1 "$TRANSACTION_FILE"
    grep ",$CUSTOMER_ID," "$TRANSACTION_FILE"
else
    # Filter by both type and customer ID
    echo "Filtering by type: $TYPE and customer: $CUSTOMER_ID"
    echo ""
    head -n 1 "$TRANSACTION_FILE"
    grep ",$CUSTOMER_ID," "$TRANSACTION_FILE" | grep ",$TYPE,"
fi

echo ""
echo "Filter complete."
