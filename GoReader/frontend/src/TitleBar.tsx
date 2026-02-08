import {Outlet, useNavigate} from "react-router-dom"

export default function TitleBar() {
    const navigator = useNavigate();
    return (
        <div className="flex flex-col">
            <div className="z-1000 flex justify-start top-0 fixed bg-blue-400 right-0 left-0">
                <div className="font-bold text-2xl p-4" onClick={() => {navigator("/")}}>
                    GoReader
                    </div>
            </div>

            <div className="flex-1 overflow-hidden pt-16 bg-white">
                <Outlet />
            </div>

            
        </div>
    )
}

