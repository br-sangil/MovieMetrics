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
        <a href="/random">
        <img className="w-40" src="/movie_random.png" alt="logo" />
      </a>
      )
    }

    return (
      <div>
       
        <MovieInfo movie={movie} />
      </div>
      
    );
}
