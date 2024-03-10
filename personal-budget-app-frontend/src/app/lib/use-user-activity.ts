'use client'

import { useEffect } from 'react';

export default function useUserActivity(onInactive: () => void) {
  console.log("useUserActivity called");
  useEffect(() => {
    let timeoutId: NodeJS.Timeout;

    function resetTimeout() {
      console.log("user active, resetting timeout");
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
}
