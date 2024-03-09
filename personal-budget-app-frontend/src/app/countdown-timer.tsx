import { useState, useEffect, createContext, useContext } from 'react';


export default function CountdownTimer ({ startInSeconds, children }: { startInSeconds: number, children: React.ReactNode}) {
    const [timeLeft, setTimeLeft] = useState(startInSeconds);

    const SessionContext = createContext({ timeLeft, setTimeLeft });
  
    useEffect(() => {
      if (timeLeft <= 0) {
        return;
      }
      const intervalId = setInterval(() => {
        setTimeLeft(timeLeft => timeLeft - 1);
      }, 1000);
  
      return () => clearInterval(intervalId);
    }, [timeLeft]);
  
    // Function to reset the timer
    const resetTimer = () => {
      setTimeLeft(startInSeconds);
    };
  
    return (
      <div>
        <h1>Countdown Timer</h1>
        <div>Time Left: {timeLeft} seconds</div>
        <button onClick={resetTimer}>Reset Timer</button>
        {children}
      </div>
    );
  };
  