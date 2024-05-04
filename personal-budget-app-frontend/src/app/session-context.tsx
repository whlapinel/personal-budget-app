'use client'

// context/SessionContext.js
import { createContext, useContext, useState, useEffect } from 'react';
import { SessionContextType, User } from '@/app/lib/data/definitions';
import { refreshToken } from './lib/data/auth';
import {useRouter} from 'next/navigation';
import { signOutAction } from './sign-out-action';
import getSessionAction from './get-session-action';
import { inactiveTimeout } from './constants/session-life-span';


let SessionContext = createContext<SessionContextType>({
  user: null,
  setUser: () => {}, // default value for setUser
  signOut: () => {}
});

export function SessionProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User>(null);
  const [alertShown, setAlertShown] = useState<boolean>(false);
  const [isActive, setIsActive] = useState(false);

  const router = useRouter();

  function getSecondsLeft() {
    return Math.round((user?.expiration! - Date.now()) / 1000);
  }

  function signOut() {
    if (user) {
      setIsActive(false)
      setAlertShown(false)
      setUser(null)
      signOutAction()
      router.push('/sign-in')
      alert('You have been signed out.');
      return;
    } else {
      alert('You are not signed in.');
    }
  }

  function onInactive() {
    setIsActive(false)
    console.log('User is inactive');
    alert('Stay signed in? If you remain inactive you will be signed out.')
  }

async function refreshSession() {
  const newSession = await refreshToken();
  const newExpiration: number = newSession?.expires;
  console.log('newExpiration:', newExpiration);
  setUser({ ...user!, expiration: newExpiration });
}

useEffect(() => {
  let timeoutId: NodeJS.Timeout;
  function resetTimeout() {
      console.log("user active, resetting timeout");
      setIsActive(true)
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        onInactive();
      }, inactiveTimeout); // milliseconds of inactivity
  };

  // Define user activities to monitor
  const events = ['click', 'mousemove', 'keypress', 'scroll', 'touchstart'];

  if (user) {
    events.forEach(event => {
      console.log("adding event listener: " + event);
      window.addEventListener(event, resetTimeout);
    });
    resetTimeout(); // Initialize the activity check
  }
  const interval = setInterval(() => {
    if (user) {
      const secondsLeft = getSecondsLeft();
      console.log('SessionProvider setInterval running! user:', user);
      console.log('Seconds left: ' + secondsLeft);
      if (secondsLeft <= 0) {
        signOut();
      }
      if (secondsLeft < 10 && isActive) {
        console.log('User is active and 10 seconds remain. Refreshing token...')
        refreshSession();
      }
    }
  }, 1000);
  return () => {
    clearInterval(interval);
    clearTimeout(timeoutId); // Clean up on component unmount
    events.forEach(event => {
      window.removeEventListener(event, resetTimeout);
    });
  }
}, [user]);

async function getSession() {
  const session = await getSessionAction();
  console.log('session:', session);
  if (session) {
    const user = {...session.user, expiration: session.expires};
    setUser(user);
  } else {
    console.log('No session found.');
    // signOut();
  }
}

useEffect(() => {
  if (!user) {
    getSession();
  }
}
, []);


return (
  <SessionContext.Provider value={{ user, setUser, signOut }}>
    {children}
  </SessionContext.Provider>
);
}

export function useSession() {
  return useContext(SessionContext);
}
