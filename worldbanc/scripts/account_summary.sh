#!/bin/bash

# WorldBanc Account Summary Script
# Displays account summary for a specific customer

CUSTOMER_FILE="../private/customers/customer_details.csv"
ACCOUNT_FILE="../private/accounts/account_balances.csv"
TRANSACTION_FILE="../private/transactions/transactions_2026_01.csv"
EMERGENCY_FILE="../private/customers/emergency_contacts.txt"

# Function to display usage
usage() {
    echo "Usage: $0 CUSTOMER_ID"
    echo ""
    echo "Display complete account summary for a customer"
    echo ""
    echo "Example:"
    echo "  $0 WB001"
    exit 1
}

# Check if customer ID is provided
if [ -z "$1" ]; then
    usage
fi

if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
    usage
fi

CUSTOMER_ID="$1"

# Check if files exist
if [ ! -f "$CUSTOMER_FILE" ] || [ ! -f "$ACCOUNT_FILE" ] || [ ! -f "$TRANSACTION_FILE" ]; then
    echo "Error: Required data files not found!"
    exit 1
fi

echo "╔════════════════════════════════════════════════════════════╗"
echo "║           WORLDBANC ACCOUNT SUMMARY REPORT                 ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# Get customer details
CUSTOMER_INFO=$(grep "^$CUSTOMER_ID," "$CUSTOMER_FILE")

if [ -z "$CUSTOMER_INFO" ]; then
    echo "Error: Customer ID $CUSTOMER_ID not found!"
    exit 1
fi

# Parse customer information
IFS=',' read -r cust_id full_name email phone address city state zip dob ssn_last4 account_opened <<< "$CUSTOMER_INFO"

echo "CUSTOMER INFORMATION"
echo "━━━━━━━━━━━━━━━━━━━━"
echo "Customer ID:      $cust_id"
echo "Name:             $full_name"
echo "Email:            $email"
echo "Phone:            $phone"
echo "Address:          $address"
echo "City, State ZIP:  $city, $state $zip"
echo "Date of Birth:    $dob"
echo "SSN (Last 4):     ****-**-$ssn_last4"
echo "Account Opened:   $account_opened"
echo ""

# Get account details
ACCOUNT_INFO=$(grep "^$CUSTOMER_ID," "$ACCOUNT_FILE")

if [ -n "$ACCOUNT_INFO" ]; then
    IFS=',' read -r cust_id account_num account_type balance available last_txn status <<< "$ACCOUNT_INFO"
    
    echo "ACCOUNT DETAILS"
    echo "━━━━━━━━━━━━━━━"
    echo "Account Number:   ****-****-****-${account_num: -4}"
    echo "Account Type:     $account_type"
    echo "Status:           $status"
    echo "Current Balance:  \$$balance"
    echo "Available:        \$$available"
    echo "Last Transaction: $last_txn"
    echo ""
fi

# Get transaction count and total
TRANSACTION_COUNT=$(grep ",$CUSTOMER_ID," "$TRANSACTION_FILE" | wc -l)
TOTAL_DEPOSITS=$(grep ",$CUSTOMER_ID," "$TRANSACTION_FILE" | grep ",DEPOSIT," | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')
TOTAL_WITHDRAWALS=$(grep ",$CUSTOMER_ID," "$TRANSACTION_FILE" | grep ",WITHDRAWAL," | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')

echo "TRANSACTION SUMMARY"
echo "━━━━━━━━━━━━━━━━━━━"
echo "Total Transactions: $TRANSACTION_COUNT"
echo "Total Deposits:     \$$TOTAL_DEPOSITS"
echo "Total Withdrawals:  \$$TOTAL_WITHDRAWALS"
echo ""

# Get recent transactions
echo "RECENT TRANSACTIONS (Last 5)"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
grep ",$CUSTOMER_ID," "$TRANSACTION_FILE" | tail -n 5 | while IFS=',' read -r txn_id cust_id acc_num date time type amount desc balance_after status; do
    echo "Date: $date $time"
    echo "Type: $type | Amount: \$$amount"
    echo "Description: $desc"
    echo "Balance After: \$$balance_after"
    echo "---"
done

# Get emergency contact
echo ""
echo "EMERGENCY CONTACT"
echo "━━━━━━━━━━━━━━━━━"
if [ -f "$EMERGENCY_FILE" ]; then
    # Extract emergency contact for this customer
    awk "/Customer ID: $CUSTOMER_ID/,/---/" "$EMERGENCY_FILE" | grep -v "Customer ID:" | grep -v "^---$" | grep -v "^$"
else
    echo "No emergency contact information available"
fi

echo ""
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                    END OF REPORT                           ║"
echo "╚════════════════════════════════════════════════════════════╝"
