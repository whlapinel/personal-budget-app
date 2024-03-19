## 3/18/24

- FIXME: email cookie should not be stored separately from session, right??  Why did I do that?
- FIXME: transactions page not showing
- FIXME: other random issues with session management 

## 3/14/24

- FIXME: add category not working, somewhere is logging: Error 1054 (42S22): Unknown column 'user_id' in 'field list' DONE

## 3/9/24

- implement session management
- 1 minute token expiration with warning at 20 seconds left
- still need to add ability for user to refresh token ("remain signed in")
- still need to verify that user does not have dashboard access without a token (stored in cookie)
- need to somehow register "activity" so that token will be refreshed after 40 seconds of inactivity
- need to implement automatic refresh of token when 30 seconds remain and user is active 

## 3/3/24

- implement basic signup and signin functionality (email and password)
- implement JWT auth and session management
- sessions stored in cookies
- storing encrypted passwords in db
- removed "username"

## 2/29/24

- began implementing jwt authentication on Go backend. Persisted on client via local storage.
- implemented jwt authentication on POST /categories route.
- Next up: implement refresh token route on client and backend, automatically make API call from client prior to token expiration.

## 2/27/24

- removed parallel routes in dashboard and 'views' variable, changed each view to regular route segments

## 2/25/24

- Created Go backend.  Super excited to be learning Go. Successfully read dummy data from Go Backend into NextJS frontend.
- Next up: 
- [x] set up MariaDB or MySQL database locally 
- [x] implement JSON web tokens (backend)

