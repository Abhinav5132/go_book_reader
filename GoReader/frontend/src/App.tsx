import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <div id='main' className="text-shadow-amber-100 p-8 rounded-2xl bg-amber-800">
            <h1 id='Title'>
                "Welcome to GoReader"
            </h1>

            <button className='border-green-500 text-2xl text-shadow-fuchsia-200 p-8'>
                "Get Started"
            </button>
        </div>
    )
}

export default App
