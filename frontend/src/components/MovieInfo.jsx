import React, {useState, useEffect} from 'react'
import Loading from './Loading'
const MovieInfo = ({movie}) => {

    if (!movie || movie == undefined || movie.length === 0) {
        return (
            <Loading />
        );
    }


    return (
        <div className="pt-10 flex flex-col justify-center items-center text-xl font-sans">
            <img src={movie.Poster} alt="Movie poster not available" />
            <p className="flex pt-6 pb-4 font-bold md:text-3xl">{movie.Title}<svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 m-auto" fill="red" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
  <path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
</svg></p>
            <p>Director: {movie.Director}</p>
            <p>Actors: {movie.Actors}</p>
            <p className="pb-8">Genre: {movie.Genre}</p>
           
        </div>
    )
}

export default MovieInfo;