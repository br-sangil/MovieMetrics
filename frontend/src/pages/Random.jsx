import React, { useState, useEffect } from 'react'
import MovieInfo from '../components/MovieInfo';

export default function Home() {
    const [movie, setMovie] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8081/random').then(response => response.json()).then(data => {
          setMovie(data);
          console.log('data', data);
        })
    }, [])

    
    if (!movie || movie === undefined || movie.length === 0) {
      return (
        <div>
          Loading
        </div>
      )
    }

    return (
      <div>
        <h1>Your Random Movie:</h1>
        <MovieInfo movie={movie} />
      </div>
      
    );
}
