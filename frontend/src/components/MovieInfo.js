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
        <div>
            <img src={movie.Poster} alt="Movie poster not available" />
            <h1>Movie title: {movie.Title}</h1>
            {/* <h1>Writer: {movie.Writer}</h1> */}
            {/* <h1>Year released: {movie.Year}</h1> */}
            <h1>Director: {movie.Director}</h1>
            <h1>Actors: {movie.Actors}</h1>
            <h1>Genre: {movie.Genre}</h1>
            {/* <h1>Plot: {movie.Plot}</h1> */}
            {/* <div>
            Ratings: <ul>
                {movie.Ratings.length > 0 ? 
                    movie.Ratings.map(rating => <li key={rating.Value}>{rating.Value}</li>)
                    :
                    <li>No ratings</li>
                }
            </ul>
            </div> */}
        </div>
    )
}

export default MovieInfo;