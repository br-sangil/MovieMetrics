import React, { useState, useEffect } from 'react'

export default function Home() {
    const [movie, setMovie] = useState({});

    useEffect(() => {
        fetch('http://localhost:8081/random').then(response => response.json()).then(data => {
          setMovie(data);
          console.log('data', data);
        })
    }, [])

  return (
    <div>
      hi
    </div>
    
  );
}
