import {Log} from "./../components/Log"
import { Footer } from './../components/Footer';

const RedisLog = () => {
    return(
        <div>
            <p className="mx-10 my-10 text-center text-6xl text-white">Log in / Log out</p>
        <Log />
        <Footer />
        </div>
    )
}

export {RedisLog};