'use client'

// context/SessionContext.js
import { createContext, useContext, useState, useEffect } from 'react';
import { SessionContextType, User } from '@/app/lib/data/definitions';
import useUserActivity from '@/app/lib/use-user-activity';


let SessionContext = createContext<SessionContextType>({
  user: null,
  setUser: () => { } // default value for setUser
});

export function SessionProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [timeLeft, setTimeLeft] = useState<number | null>(null);
  const [alertShown, setAlertShown] = useState<boolean>(false);
  
  function signOut() {
    alert('You have been logged out.');
    setTimeLeft(null)
    setAlertShown(false)
    setUser(null)
  }

  useUserActivity(() => {
    console.log('User is inactive');
  });


  useEffect(() => {
    if (timeLeft) {
      console.log("timeleft: ", timeLeft)
      if ((timeLeft < 20) && (!alertShown)) {
        console.log('timeLeft: ', timeLeft);
        console.log('timeLeft < 20');
        alert('Your session will expire in 20 seconds');
        setAlertShown(true);
      }
      if (timeLeft < 0) {
        alert('You have been logged out.');
        setTimeLeft(null)
        setAlertShown(false)
        setUser(null)
      }
    }
  }, [timeLeft])

  useEffect(() => { console.log("user: ", user) }, [user])

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

      // if user is active and 20 seconds remain before expiration, alert user
    }, 1000);
    return () => clearInterval(interval);
  }, [user]);

  return (
    <SessionContext.Provider value={{ user, setUser }}>
      {children}
    </SessionContext.Provider>
  );
}

export function useSession() {
  return useContext(SessionContext);
}
