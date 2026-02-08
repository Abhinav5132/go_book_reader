import {useState} from 'react';
import './App.css';
import {GetBookPath, OpenBookFile} from "../wailsjs/go/main/App"

function App() {

    const [bookData, setBookData] = useState<String | null>(null)
    const txtPath = async () => {
        const data = await GetBookPath()

        if (!data || data == "") return 
        setBookData(data)
    }

    if (bookData) {
        return (
            <div className='p-8 text-black font-bold whitespace-pre-line leading-relaxed text-left'>
                <div>{bookData}</div>
            </div>
        )
    }

    else {
        return <div>
            <button onClick={txtPath}>
                Open Book
            </button>
        </div >
    }
}

export default App
