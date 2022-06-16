import React, {useState, useEffect} from 'react'

const MovieInfo = ({movie}) => {

    if (!movie || movie == undefined || movie.length === 0) {
        return (
            <div>
                Loading
            </div>
        )
    }


    return (
        <div className="pt-10 flex flex-col justify-center items-center text-xl font-sans">
            <img src={movie.Poster} alt="Movie poster not available" />
            <p className="pt-6 pb-4 font-bold text-3xl">{movie.Title}</p>
            <p>Director: {movie.Director}</p>
            <p>Actors: {movie.Actors}</p>
            <p className="pb-8">Genre: {movie.Genre}</p>
           
        </div>
    )
}

export default MovieInfo;