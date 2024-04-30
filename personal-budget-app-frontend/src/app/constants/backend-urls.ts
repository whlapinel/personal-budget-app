const production = process.env.PRODUCTION
let host: string = ''
if (production === 'true') {
    host = 'backend'
} else {
    host = 'localhost'
}
const baseUrl = `http://${host}:8080`;

console.log('baseUrl:', baseUrl);

export const backendUrls = {
    signin: `${baseUrl}/signin`,
    categories: `${baseUrl}/categories`,
    assignments: `${baseUrl}/assignments`,
    transactions: `${baseUrl}/transactions`,
    accounts: `${baseUrl}/accounts`,
    goals: `${baseUrl}/goals`,
    users: `${baseUrl}/users`,
    monthlyBudgets: `${baseUrl}/monthly-budgets`,
    budgetPageData: `${baseUrl}/budget-page-data`,
}