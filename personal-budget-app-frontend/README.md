# Will Lapinel's Personal Budget Web Application

## Created for ITIS 5166, Network-Based Applications, UNCC

### Feature/Layout Requirements (provided by Professor)
https://docs.google.com/presentation/d/121C_A62cB5nhCZLu0aTEuNjg54KBki7RySTbAonfC2o/edit?usp=sharing

- [] Homepage
- [] Login
- [] Logout
- [] Signup
- [] Dashboard
- [] 3 visualizations (graphs and charts, etc) available from dashboard
- [] Ability for user to configure budgets
- [] Enter expense for budget/month

### Technical Requirements (provided by Professor)

- [] API only / completely detached from frontend *
- [] All responses in JSON
- [] Proper status code (200, 400, 500)
- [] Add gzip or some other compression
- [] MySQL database
- [] Auth
- [] Auth: Token needs to expire in 1 minute
- [] Auth: After 40 seconds of inactivity show warning that token will expire in 20 seconds, provide button to reset
- [] Testing: at least 1 Unit Test
- [] Testing: at least 1 E2E Test
- [] Testing: at least 1 Visual Regression test with Applitools

- \* I'm assuming this is in order to provide the potential for developing other frontends as opposed to a backend that is designed exclusively for the NextJS UI

### Possible extra features (by me, the developer)

- [] Account register with ability to download bank transactions via Plaid

### Architectural plan (by me, the developer)

#### NextJS full stack application will handle: 

- UI 
- authentication 

#### Separate Python/Django server will handle 

- authorization 
- business logic 
- database queries


