import Chat from '../components/Chat';
import {useState, useEffect} from 'react'
import axios from 'axios';
import {useNavigate} from 'react-router-dom'
import MovieInfo from '../components/MovieInfo';

export default function Home() {
  const navigate = useNavigate();
  const [search, setSearch] = useState("");
  const [movies, setMovies] = useState(null);

  const handleSearch = async (event) => {
    event.preventDefault();
    const movie = await axios.get(`http://localhost:8081/?t=${encodeURI(search)}`);
    const data = movie.data;
    setSearch('');
    console.log(data);
    setMovies(data);
    // console.log(movies);
    const element = document.querySelector('#movielist')
    const topPos = element.getBoundingClientRect().top + window.pageYOffset
    
    window.scrollTo({
      top: topPos, // scroll so that the element is at the top of the view
      behavior: 'smooth' // smooth scroll
    })
    
  }


  return (
    <div className="bg-movie-posters flex h-screen justify-center items-center flex flex-col">
        

          <form onSubmit={handleSearch} className="flex justify-center" action="#movieList">
            <input value={search} onChange={(e)=>setSearch(e.target.value)} type="text" className=" px-4 inline-flex rounded-l w-40 md:w-64 lg:w-96 h-10 bg-gray-200" placeholder="Search Movie by Title..."></input>
            <button type="submit" className ="h-10 w-12 rounded-r bg-red-600 hover:bg-red-700">
              <svg xmlns="http://www.w3.org/2000/svg" className="m-auto h-5 w-5" viewBox="0 0 20 20" fill="white">
                <path fillRule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clipRule="evenodd" />
              </svg>
            </button>
          </form>
        

        <div className="p-10 absolute bottom-0 right-0">
          <Chat />
        </div>

        {/* Search results */}
    <div className="flex flex-row pt-40">
      <a id="movielist"></a>
      {movies &&
        movies.map(movie =>
          <div className='bg-slate-100 drop-shadow-2xl rounded-md mr-10 p-4'>
            <MovieInfo movie={movie} key={movie.Title} />
            </div>
          )
      }
    </div>
    </div>

    
  
  );
}
