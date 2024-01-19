
type BudgetCategory = {
    id: string,
    name: string,
    available: number,
    assigned: number,
    transactions: Transaction[]
}

type Transaction = {
    id: string,
    date: Date,
    amount: number,
    description: string,
    category: BudgetCategory | Transaction[] // for split transactions?  Not sure
}

type Account = {
    id: string,
    name: string,
    bankName: string,
    transactions: Transaction[]
    balance: number,
}