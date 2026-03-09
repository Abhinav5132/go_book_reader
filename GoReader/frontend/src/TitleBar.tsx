import { useState } from "react";
import {Outlet, useNavigate} from "react-router-dom"

export default function TitleBar() {
    const navigator = useNavigate();
    const [SideBarState, setSideBarState] = useState(false);

    return (
        <div className="">
            <div className="flex flex-row bg-blue-400 p-2">
                <button className="font-bold bg-blue-500 rounded-2xl p-4 drop-shadow-2xl" 
                    onClick={() => {
                        setSideBarState(SideBarState => !SideBarState);
                    }
                }>
                    SideBar
                </button>

                <div >
                    <div className="font-bold text-2xl p-4" onClick={() => {navigator("/")}}>
                        GoReader
                    </div>
                </div>
            </div>
            
            <div className="flex flex-row">
                <SideBar isActive={SideBarState} />
                <div className="flex-1 overflow-hidden pt-16 bg-white">
                    <Outlet />
                </div>
                
            </div>
        </div>
    )
}

type SideBarProps = {
    isActive: Boolean
}

function SideBar({isActive}: SideBarProps) {

    if (!isActive) {
        return null
    }
    
    
    return (
        <div className="bg-blue-400">
            This is the sidebar    
        </div>
    )
}