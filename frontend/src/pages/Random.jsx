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
      );
    }

    return (
      <div className="bg-movie-posters flex h-screen justify-center items-center">
          <div className="bg-slate-100 w-4/5 h-3/4 drop-shadow-2xl rounded-lg">
            <img className="pt-4 m-auto" src="/movie_random.png" alt="logo" />
            <MovieInfo movie={movie} />
          </div>

      </div>

       
      
    );
}
