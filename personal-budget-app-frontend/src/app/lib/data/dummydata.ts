import { Account, Transaction } from './definitions'

export const categories = [
    { name: 'Food', id: 1, amount: 0 },
    { name: 'Rent', id: 2, amount: 0 },
    { name: 'Utilities', id: 3, amount: 0 },
    { name: 'Entertainment', id: 4, amount: 0 },
    { name: 'Transportation', id: 5, amount: 0 },
    { name: 'Other', id: 6, amount: 0 },
]


export const transactions: Transaction[] = [
    {
        id: '1',
        account: 'Checking',
        date: new Date(),
        payee: 'Walmart',
        amount: -100,
        memo: 'Groceries',
        category: 'Food',
    },
    {
        id: '2',
        account: 'Checking',
        date: new Date(),
        payee: 'Transfer to Savings',
        amount: -100,
        memo: '',
        category: '',
    },
    {
        id: '3',
        account: 'Savings',
        date: new Date(),
        payee: 'Transfer from Checking',
        amount: 100,
        memo: '',
        category: '',
    },
    {
        id: '4',
        account: 'Credit Card',
        date: new Date(),
        payee: 'Hannah Anderson',
        amount: 100,
        memo: '',
        category: '',
    }
]

export const accounts: Account[] = [
    {
        id: '1',
        name: 'Checking',
        bankName: 'Chase',
        balance: 1000,
        type: 'checking',
    },
    {
        id: '2',
        name: 'Savings',
        bankName: 'Chase',
        balance: 10000,
        type: 'savings',
    },
    {
        id: '3',
        name: 'Credit Card',
        bankName: 'Chase',
        balance: -1000,
        type: 'credit',
    }
]
