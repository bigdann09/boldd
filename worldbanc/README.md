# WorldBanc

A comprehensive banking filesystem project designed for learning terminal commands, shell scripting, and file system operations.

## Overview

WorldBanc is an educational banking simulation that provides a realistic file structure with customer data, transaction records, account information, and operational scripts. This project is perfect for practicing:

- Terminal navigation and file system operations
- Shell scripting and data processing
- CSV and text file manipulation
- Understanding file permissions and access control
- Working with structured data

## Project Structure

```
worldbanc/
├── private/                    # Confidential data (restricted access)
│   ├── customers/
│   │   ├── customer_details.csv       # Full customer information
│   │   └── emergency_contacts.txt     # Emergency contact database
│   ├── accounts/
│   │   └── account_balances.csv       # Account balances and details
│   ├── transactions/
│   │   └── transactions_2026_01.csv   # Transaction history
│   └── reports/
│       └── monthly_report_2026_01.txt # Internal financial reports
│
├── public/                     # Public-facing information
│   ├── customers/
│   │   └── customer_list.txt          # Public customer directory
│   ├── accounts/
│   │   └── account_types.txt          # Account type descriptions
│   ├── transactions/
│   │   └── transaction_summary.txt    # Transaction statistics
│   └── reports/
│       └── branch_locations.txt       # Branch location information
│
├── scripts/                    # Utility scripts for data operations
│   ├── filter_transactions.sh         # Filter transactions by type/customer
│   ├── generate_customer_id.sh        # Generate new customer IDs
│   ├── account_summary.sh             # Display customer account summary
│   └── calculate_totals.sh            # Calculate financial totals
│
└── README.md                   # This file
```

## Features

### Data Files

- **Customer Information**: 20 sample customers with realistic data
- **Account Balances**: Current balances for Premium, Standard, and Business accounts
- **Transaction History**: 40+ transactions including deposits, withdrawals, payments, and transfers
- **Emergency Contacts**: Contact information for customer emergencies
- **Branch Locations**: 18 branch locations across the United States

### Shell Scripts

All scripts are located in the `scripts/` directory and are executable.

#### 1. Filter Transactions (`filter_transactions.sh`)

Filter transactions by type and/or customer ID.

```bash
cd scripts
./filter_transactions.sh DEPOSIT              # Show all deposits
./filter_transactions.sh DEPOSIT WB001        # Show deposits for customer WB001
./filter_transactions.sh "" WB001             # Show all transactions for WB001
```

**Transaction Types**: DEPOSIT, WITHDRAWAL, PAYMENT, TRANSFER

#### 2. Generate Customer ID (`generate_customer_id.sh`)

Generate new customer IDs in the WorldBanc format (WB###).

```bash
cd scripts
./generate_customer_id.sh        # Generate 1 ID
./generate_customer_id.sh 5      # Generate 5 IDs
```

#### 3. Account Summary (`account_summary.sh`)

Display a complete account summary for a specific customer.

```bash
cd scripts
./account_summary.sh WB001       # Show summary for customer WB001
```

Displays:
- Customer information
- Account details
- Transaction summary
- Recent transactions
- Emergency contact

#### 4. Calculate Totals (`calculate_totals.sh`)

Calculate various financial totals and statistics.

```bash
cd scripts
./calculate_totals.sh --accounts      # Calculate account balance totals
./calculate_totals.sh --transactions  # Calculate transaction totals
./calculate_totals.sh --summary       # Display complete financial summary
```

## Getting Started

### Prerequisites

- Unix-like operating system (Linux, macOS, WSL)
- Bash shell
- Basic command-line tools (grep, awk, cut, etc.)

### Installation

1. Clone or download this repository:
```bash
git clone <repository-url>
cd worldbanc
```

2. Verify the directory structure:
```bash
ls -la
tree  # If tree is installed
```

3. Navigate to the scripts directory:
```bash
cd scripts
```

4. Verify scripts are executable:
```bash
ls -l *.sh
```

If not executable, run:
```bash
chmod +x *.sh
```

## Learning Exercises

### Beginner Level

1. **Navigation Practice**
   - Navigate through the directory structure
   - List files in different directories
   - Display file contents using `cat`, `less`, or `more`

2. **File Operations**
   - Count the number of customers: `wc -l private/customers/customer_details.csv`
   - View the first 5 transactions: `head -n 6 private/transactions/transactions_2026_01.csv`
   - Search for a specific customer: `grep "WB001" private/customers/customer_details.csv`

3. **Basic Filtering**
   - Find all Premium accounts: `grep "Premium" private/accounts/account_balances.csv`
   - Count deposits: `grep -c "DEPOSIT" private/transactions/transactions_2026_01.csv`

### Intermediate Level

1. **Using Scripts**
   - Run all provided scripts with different parameters
   - Analyze the output and understand the data

2. **Data Analysis**
   - Calculate average account balance
   - Find the customer with the highest balance
   - Count transactions per customer

3. **Custom Filtering**
   - Find all transactions over $1000
   - List customers from a specific state
   - Find accounts with balance below minimum

### Advanced Level

1. **Script Modification**
   - Modify scripts to add new features
   - Create new scripts for custom operations
   - Add error handling and validation

2. **Data Processing**
   - Generate monthly reports
   - Calculate interest on accounts
   - Identify suspicious transactions

3. **Automation**
   - Create a script to process new transactions
   - Automate customer onboarding
   - Generate automated reports

## Sample Commands

### View Customer Information
```bash
# View all customers
cat public/customers/customer_list.txt

# View detailed customer information
cat private/customers/customer_details.csv

# Search for a specific customer
grep "Sarah Johnson" private/customers/customer_details.csv
```

### Analyze Transactions
```bash
# View all transactions
cat private/transactions/transactions_2026_01.csv

# Count transactions by type
grep -c "DEPOSIT" private/transactions/transactions_2026_01.csv
grep -c "WITHDRAWAL" private/transactions/transactions_2026_01.csv

# Find large transactions (over $1000)
awk -F',' '$7 > 1000 {print}' private/transactions/transactions_2026_01.csv
```

### Account Analysis
```bash
# View account types
cat public/accounts/account_types.txt

# Find Premium accounts
grep "Premium" private/accounts/account_balances.csv

# Calculate total balance (requires awk)
awk -F',' 'NR>1 {sum+=$4} END {print "Total: $"sum}' private/accounts/account_balances.csv
```

### Using Pipes and Redirection
```bash
# Save filtered results to a file
grep "WB001" private/transactions/transactions_2026_01.csv > wb001_transactions.txt

# Count and sort transaction types
cut -d',' -f6 private/transactions/transactions_2026_01.csv | sort | uniq -c

# Find customers in California
grep ",CA," private/customers/customer_details.csv | cut -d',' -f2
```

## Data Format Reference

### Customer Details CSV
```
customer_id,full_name,email,phone,address,city,state,zip_code,date_of_birth,ssn_last4,account_opened
```

### Account Balances CSV
```
customer_id,account_number,account_type,balance,available_balance,last_transaction_date,status
```

### Transactions CSV
```
transaction_id,customer_id,account_number,date,time,type,amount,description,balance_after,status
```

## Security Notes

- The `private/` directory contains sensitive information and should have restricted access in a real-world scenario
- The `public/` directory contains information that can be shared with customers
- Customer SSN data is truncated to last 4 digits only
- Account numbers are displayed with masking in reports

## Educational Use

This project is designed for educational purposes to help learn:

- **File System Navigation**: Understanding directory structures and file organization
- **Shell Commands**: Practicing grep, awk, sed, cut, sort, and other utilities
- **Data Processing**: Working with CSV and text files
- **Shell Scripting**: Writing and executing bash scripts
- **Access Control**: Understanding public vs. private data separation

## Contributing

This is an educational project. Feel free to:
- Add more customer data
- Create additional scripts
- Enhance existing functionality
- Add new features and reports

## License

This project is provided for educational purposes.

## Support

For questions or issues:
- Review the script help messages: `./script_name.sh --help`
- Check the data format reference above
- Examine the sample commands section

## Acknowledgments

Inspired by the Boot.dev "Learn Terminals and Shells" course and designed to provide hands-on experience with real-world banking data structures.

---

**WorldBanc** - Where Learning Meets Banking
