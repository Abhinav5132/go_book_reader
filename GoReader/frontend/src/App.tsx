import {useContext, useEffect, useRef, useState} from 'react';
import { Routes, Route } from "react-router-dom"
import './App.css';
import TitleBar from './TitleBar';
import HomePage from './HomePage';
import { BookDataContext } from './BookDataContext';
function App() {
    const [bookData, setBookData] = useState<string | null>(null)
    return (
    <BookDataContext.Provider value={{bookData, setBookData}}>
        <Routes>
            <Route element={<TitleBar />}>
                    <Route path = "/" element={<HomePage />}/>
                    <Route path = "/open-txt" element={<OpenTxt />} />
            </Route>
        </Routes>
    </BookDataContext.Provider>
    
    )
}


export function OpenTxt() {
    
    const ctx = useContext(BookDataContext)

    if (!ctx || !ctx.bookData) return (
        <div>
            Unable to Open File
        </div>
    )
    
    const {bookData} = ctx;
    return (
        <div className='p-8 text-black font-bold whitespace-pre-line 
        leading-relaxed text-left'>
            <div>{bookData}</div>
        </div>
    )
}

export default App