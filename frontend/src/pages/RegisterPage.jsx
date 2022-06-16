import React, {useState}from 'react';
import { useNavigate } from 'react-router-dom';

const RegisterPage = () => {
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    // TODO : When backend has endpoint, send post request with user info
    const handleRegister = (event) => {
        event.preventDefault();
        setName('');
        setUsername('');
        setPassword('');
        navigate('/login');
    }
    return (
        <div className="bg-movie-posters flex h-screen justify-center items-center flex-row">
            <div className="bg-white w-auto h-96 drop-shadow-2xl rounded-xl justify-center items-center flex-row">
            <form onSubmit={handleRegister} className="p-10 flex flex-col text-sm font-sans text-slate-600 justify-center items-center">
                
                <div className="p-3">Name: <input className="w-40 bg-gray-100" value={name} onChange={(e)=>setName(e.target.value)}/></div>
                <div className="p-3">  Username: <input className="w-40 bg-gray-100" value={username} onChange={(e)=>setUsername(e.target.value)} /></div>
                <div className="p-3">  Password: <input className="w-40 bg-gray-100" value={password} type="password" onChange={(e)=>setPassword(e.target.value)} /></div>
                    <button className="font-bold p-3 my-4 w-40 rounded-full bg-red-600 hover:bg-red-700 text-white" type="submit">Register</button>
                
            </form>
            </div>
        </div>
    )
}

export default RegisterPage;