import { Link } from 'react-router-dom';

function Home() {
  return (
    <div className="home flex flex-col md:flex-row items-center justify-center min-h-[calc(100vh-80px)]">
      <div className="home-content md:w-6/7 text-center md:text-left">
        <h1 className="home-title text-5xl md:text-6xl leading-[4rem] md:leading-[5rem] font-bold text-dark-green uppercase leading-tight">
          Little Alchemy Element Finder
        </h1>
        <p className="text-dark-green text-lg mt-2">
          Made with love by 13523057, 13523067, 13523094
        </p>
        <p className="text-dark-green text-base md:text-lg mt-4 max-w-md mx-auto md:mx-0">
          A simple Little Alchemy 2 elements recipes visualizer using BFS, DFS, and bidirectional search. Made in Go & React js.
        </p>
        <Link to="/main">
          <button className="get-started mt-6 px-6 py-3 bg-dark-green text-beige rounded-full text-lg font-medium hover:bg-dark-green-dark transition-all">
            Get started!
          </button>
        </Link>
      </div>
      <div className="home-image md:w-1/2 mt-8 md:mt-0 flex justify-center">
        
            <img
            src="/beaker.svg"
            alt="Beaker Illustration"
            className="w-64 h-64 md:w-80 md:h-80"
            />
        
      </div>
    </div>
  );
}

export default Home;