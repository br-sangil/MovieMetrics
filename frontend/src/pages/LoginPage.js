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
        <div>
            Login Here
            <form onSubmit={handleLogin}>
                Username: <input value={username} onChange={(e)=>setUsername(e.target.value)} />
                Password: <input type="password" value={password} onChange={(e)=>setPassword(e.target.value)} />
                <button type="submit">Login</button>
            </form>
        </div>
    )
}

export default LoginPage;