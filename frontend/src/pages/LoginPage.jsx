import React, {useState} from 'react';
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    // TODO: When backend endpoint is created, send info to backend to verify login and log the user in
    const handleLogin = (event) => {
        event.preventDefault();
        setUsername('');
        setPassword('');
        navigate('/')
    }

    return (
        <div className="bg-movie-posters flex h-screen justify-center items-center flex-row">
             <div className="bg-white w-auto h-96 drop-shadow-2xl rounded-xl justify-center items-center flex-row">
            <form onSubmit={handleLogin} className="p-10 flex flex-col text-sm font-sans text-slate-600 justify-center items-center">
                <div className="p-3">Username: <input className="w-40 bg-gray-100" value={username} onChange={(e)=>setUsername(e.target.value)} /> </div>
                <div className="p-3">Password: <input className="w-40 bg-gray-100" type="password" value={password} onChange={(e)=>setPassword(e.target.value)} /></div>
                <button className="font-bold p-3 my-4 w-40 rounded-full bg-red-600 hover:bg-red-700 text-white" type="submit">Login</button>
            </form>
            </div>
        </div>

    )
}

export default LoginPage;