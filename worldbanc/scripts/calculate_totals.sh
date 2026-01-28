#!/bin/bash

# WorldBanc Financial Totals Calculator
# Calculates various financial totals from account and transaction data

ACCOUNT_FILE="../private/accounts/account_balances.csv"
TRANSACTION_FILE="../private/transactions/transactions_2026_01.csv"

# Function to display usage
usage() {
    echo "Usage: $0 [OPTION]"
    echo ""
    echo "Calculate financial totals and statistics"
    echo ""
    echo "Options:"
    echo "  -a, --accounts     Calculate total account balances"
    echo "  -t, --transactions Calculate transaction totals"
    echo "  -s, --summary      Display complete financial summary"
    echo "  -h, --help         Display this help message"
    echo ""
    echo "Examples:"
    echo "  $0 --accounts"
    echo "  $0 --transactions"
    echo "  $0 --summary"
    exit 0
}

# Function to calculate account totals
calculate_accounts() {
    echo "ACCOUNT BALANCE TOTALS"
    echo "======================"
    echo ""
    
    # Total balance across all accounts
    TOTAL_BALANCE=$(tail -n +2 "$ACCOUNT_FILE" | cut -d',' -f4 | awk '{sum+=$1} END {printf "%.2f", sum}')
    echo "Total Balance (All Accounts): \$$TOTAL_BALANCE"
    echo ""
    
    # Balance by account type
    echo "By Account Type:"
    echo "----------------"
    
    PREMIUM_TOTAL=$(grep ",Premium," "$ACCOUNT_FILE" | cut -d',' -f4 | awk '{sum+=$1} END {printf "%.2f", sum}')
    PREMIUM_COUNT=$(grep -c ",Premium," "$ACCOUNT_FILE")
    echo "Premium:  $PREMIUM_COUNT accounts | \$$PREMIUM_TOTAL"
    
    STANDARD_TOTAL=$(grep ",Standard," "$ACCOUNT_FILE" | cut -d',' -f4 | awk '{sum+=$1} END {printf "%.2f", sum}')
    STANDARD_COUNT=$(grep -c ",Standard," "$ACCOUNT_FILE")
    echo "Standard: $STANDARD_COUNT accounts | \$$STANDARD_TOTAL"
    
    BUSINESS_TOTAL=$(grep ",Business," "$ACCOUNT_FILE" | cut -d',' -f4 | awk '{sum+=$1} END {printf "%.2f", sum}')
    BUSINESS_COUNT=$(grep -c ",Business," "$ACCOUNT_FILE")
    echo "Business: $BUSINESS_COUNT accounts | \$$BUSINESS_TOTAL"
    
    echo ""
    
    # Average balance
    TOTAL_ACCOUNTS=$(tail -n +2 "$ACCOUNT_FILE" | wc -l)
    AVG_BALANCE=$(echo "$TOTAL_BALANCE $TOTAL_ACCOUNTS" | awk '{printf "%.2f", $1 / $2}')
    echo "Average Account Balance: \$$AVG_BALANCE"
    echo ""
}

# Function to calculate transaction totals
calculate_transactions() {
    echo "TRANSACTION TOTALS"
    echo "=================="
    echo ""
    
    # Total transactions
    TOTAL_TXN=$(tail -n +2 "$TRANSACTION_FILE" | wc -l)
    echo "Total Transactions: $TOTAL_TXN"
    echo ""
    
    # By transaction type
    echo "By Transaction Type:"
    echo "--------------------"
    
    DEPOSIT_COUNT=$(grep -c ",DEPOSIT," "$TRANSACTION_FILE")
    DEPOSIT_TOTAL=$(grep ",DEPOSIT," "$TRANSACTION_FILE" | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')
    echo "Deposits:    $DEPOSIT_COUNT transactions | \$$DEPOSIT_TOTAL"
    
    WITHDRAWAL_COUNT=$(grep -c ",WITHDRAWAL," "$TRANSACTION_FILE")
    WITHDRAWAL_TOTAL=$(grep ",WITHDRAWAL," "$TRANSACTION_FILE" | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')
    echo "Withdrawals: $WITHDRAWAL_COUNT transactions | \$$WITHDRAWAL_TOTAL"
    
    PAYMENT_COUNT=$(grep -c ",PAYMENT," "$TRANSACTION_FILE")
    PAYMENT_TOTAL=$(grep ",PAYMENT," "$TRANSACTION_FILE" | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')
    echo "Payments:    $PAYMENT_COUNT transactions | \$$PAYMENT_TOTAL"
    
    TRANSFER_COUNT=$(grep -c ",TRANSFER," "$TRANSACTION_FILE")
    TRANSFER_TOTAL=$(grep ",TRANSFER," "$TRANSACTION_FILE" | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')
    echo "Transfers:   $TRANSFER_COUNT transactions | \$$TRANSFER_TOTAL"
    
    echo ""
    
    # Total transaction volume
    TOTAL_VOLUME=$(tail -n +2 "$TRANSACTION_FILE" | cut -d',' -f7 | awk '{sum+=$1} END {printf "%.2f", sum}')
    echo "Total Transaction Volume: \$$TOTAL_VOLUME"
    
    # Average transaction amount
    AVG_TXN=$(echo "$TOTAL_VOLUME $TOTAL_TXN" | awk '{printf "%.2f", $1 / $2}')
    echo "Average Transaction Amount: \$$AVG_TXN"
    echo ""
}

# Function to display complete summary
display_summary() {
    echo "╔════════════════════════════════════════════════════════════╗"
    echo "║         WORLDBANC FINANCIAL SUMMARY REPORT                 ║"
    echo "║              Generated: $(date '+%Y-%m-%d %H:%M:%S')                  ║"
    echo "╚════════════════════════════════════════════════════════════╝"
    echo ""
    
    calculate_accounts
    echo ""
    calculate_transactions
    
    echo "╔════════════════════════════════════════════════════════════╗"
    echo "║                    END OF REPORT                           ║"
    echo "╚════════════════════════════════════════════════════════════╝"
}

# Check if files exist
if [ ! -f "$ACCOUNT_FILE" ] || [ ! -f "$TRANSACTION_FILE" ]; then
    echo "Error: Required data files not found!"
    exit 1
fi

# Parse command line arguments
case "$1" in
    -a|--accounts)
        calculate_accounts
        ;;
    -t|--transactions)
        calculate_transactions
        ;;
    -s|--summary)
        display_summary
        ;;
    -h|--help|"")
        usage
        ;;
    *)
        echo "Error: Unknown option '$1'"
        echo ""
        usage
        ;;
esac
