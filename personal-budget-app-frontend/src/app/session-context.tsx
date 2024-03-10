'use client'

// context/SessionContext.js
import { createContext, useContext, useState, useEffect } from 'react';
import { SessionContextType, User } from '@/app/lib/data/definitions';
import useUserActivity from '@/app/lib/use-user-activity';
import { updateSession } from './lib/data/auth';


let SessionContext = createContext<SessionContextType>({
  user: null,
  setUser: () => { }, // default value for setUser
  isActive: false
});

export function SessionProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [timeLeft, setTimeLeft] = useState<number | null>(null);
  const [alertShown, setAlertShown] = useState<boolean>(false);
  const [isActive, setIsActive] = useState<boolean>(false);

  function signOut() {
    alert('You have been logged out.');
    setIsActive(false);
    setTimeLeft(null)
    setAlertShown(false)
    setUser(null)
  }

  function onInactive() {
    console.log('User is inactive');
  }

  async function refreshToken() {
    console.log('refreshing token');
    const didUpdate = await updateSession();
    if (didUpdate) { console.log('token refreshed') }
    else { console.log('token not refreshed')}      
  }


  useEffect(() => {
    if (timeLeft) {
      console.log("timeleft: ", timeLeft)
      if ((timeLeft < 20) && (!alertShown)) {
        console.log('timeLeft: ', timeLeft);
        console.log('timeLeft < 20');
        const response = confirm('Your session will expire in 20 seconds. Do you want to remain signed in?');
        if (response) {
          console.log('User wants to remain signed in');
          // send request to server to extend session
          // if successful, set new expiration time
          // if not successful, sign out
        } else {
          signOut();
        }
        setAlertShown(true);
      }
      if ((timeLeft < 0) && (!isActive)) {
        signOut();
      }
      if ((timeLeft < 10) && (isActive)) {
        // refresh token
        refreshToken();
  
      }
    }
  }, [timeLeft])

  useEffect(() => { console.log("user: ", user) }, [user])

  console.log("useUserActivity called");
  useEffect(() => {
    let timeoutId: NodeJS.Timeout;

    function resetTimeout() {
      console.log("user active, resetting timeout");
      setIsActive(true);
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        onInactive(); // Trigger the callback when the user is inactive
      }, 40000); // 40 seconds of inactivity
    };

    // Define user activities to monitor
    const events = ['click', 'mousemove', 'keypress', 'scroll', 'touchstart'];

    events.forEach(event => {
      console.log("adding event listener: " + event);
      window.addEventListener(event, resetTimeout);
    });

    resetTimeout(); // Initialize the activity check

    return () => {
      clearTimeout(timeoutId); // Clean up on component unmount
      events.forEach(event => {
        window.removeEventListener(event, resetTimeout);
      });
    };
  }, [onInactive]);


  useEffect(() => {
    // check every second

    const interval = setInterval(() => {
      if (user?.expiration) {
        console.log('setting timeLeft');
        console.log('user.expiration:', user?.expiration);
        const now = Date.now();
        console.log('now:', now);
        const timeLeftInSeconds = (user?.expiration - now) / 1000;
        console.log(timeLeftInSeconds);
        setTimeLeft(timeLeftInSeconds);
        console.log(new Date(user?.expiration ? user?.expiration : 0));
        console.log("timeLeft < 20: ", (timeLeft ? timeLeft < 20 : null));

        console.log('SessionProvider setInterval running! user:', user);
      }
    }, 1000);
    return () => clearInterval(interval);
  }, [user]);

  return (
    <SessionContext.Provider value={{ user, setUser, isActive }}>
      {children}
    </SessionContext.Provider>
  );
}

export function useSession() {
  return useContext(SessionContext);
}
