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
        <div className="bg-movie-posters  flex h-screen justify-center items-center">
      <div className="bg-slate-100 w-4/5 h-3/4 drop-shadow-2xl rounded-lg">
      <img className="pt-8 m-auto" src="/loading.png" alt="logo" />
    <svg xmlns="http://www.w3.org/2000/svg" className="animate-spin m-auto h-44 w-60" viewBox="0 0 20 20" fill="red">
<path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd" />
</svg>
</div>
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
