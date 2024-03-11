'use client'

// context/SessionContext.js
import { createContext, useContext, useState, useEffect } from 'react';
import { SessionContextType, User } from '@/app/lib/data/definitions';


let SessionContext = createContext<SessionContextType>({
  user: null,
  setUser: () => { } // default value for setUser
});

export function SessionProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [timeLeft, setTimeLeft] = useState<number | null>(null);
  const [alertShown, setAlertShown] = useState<boolean>(false);
  const [isActive, setIsActive] = useState(false);
  
  function signOut() {
    alert('You have been logged out.');
    setTimeLeft(null)
    setIsActive(false)
    setAlertShown(false)
    setUser(null)
  }
  
  async function onInactive(){

  }
  
  useEffect(() => {
    let timeoutId: NodeJS.Timeout;

    function resetTimeout() {

      console.log("user active, resetting timeout");
      setIsActive(true)
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        if (!confirm('You will be logged out in 20 seconds due to inactivity. Do you wish to remain signed in?')) {
          console.log('User is inactive');
          setIsActive(false)
        } else {
          console.log('user chose to remain signed in. Refreshing token...')
          // call refresh token action
        }
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
  }, [user?.expiration]);



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
        signOut()
      }
      if (timeLeft < 10 && isActive) {
        console.log('User is active and 10 seconds remain. Refreshing token...')
      }
    }
  }, [user?.expiration])

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
