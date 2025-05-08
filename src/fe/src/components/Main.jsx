import { useState } from 'react';
import SearchBar from './SearchBar';
import AlgorithmSelector from './AlgorithmSelector';
import RecipeTree from './RecipeTree';
import axios from 'axios';

function Main() {
  const [element, setElement] = useState(''); 
  const [algorithm, setAlgorithm] = useState('BFS'); 
  const [mode, setMode] = useState('single');
  const [maxRecipes, setMaxRecipes] = useState(1);
  const [recipes, setRecipes] = useState([]);
  const [searchTime, setSearchTime] = useState(0);
  const [nodesVisited, setNodesVisited] = useState(0);

  const dummyRespons = {
    recipes: [
      {
        id: 0,
        name: "Continent",
        children: [
          {
            name: "Land",
            children: [
              { name: "Earth" },
              { name: "Earth" }
            ]
          },
          { name: "Earth" }
        ]
      }
    ],
    time: 0.123,
    nodesVisited: 15
  };

  const handleSearch = async () => {
    try {
      const payload = {
        element,
        algorithm,
        mode,
        maxRecipes: mode === 'multiple' ? maxRecipes : 1,
      };
      // TODO: ke api asli
      // const response = await axios.post('http://localhost:8080/api/search', payload);
      const response = { data: element.toLowerCase() === 'continent' ? dummyRespons : { recipes: [], time: 0, nodesVisited: 0 } };

      setRecipes(response.data.recipes);
      setSearchTime(response.data.time);
      setNodesVisited(response.data.nodesVisited);
    } catch (error) {
      console.error('Error processing search:', error);
      try {
        const localDataModule = await import('../../../scraper/elements.json');
        const localData = localDataModule.default;
        const found = localData.find((e) => e.element.toLowerCase() === element.toLowerCase());
        if (found) {
          const formattedRecipes = found.recipes.slice(0, mode === 'multiple' ? maxRecipes : 1).map((recipe, idx) => ({
            id: idx,
            name: found.element,
            children: recipe.map((child) => ({ name: child })),
          }));
          setRecipes(formattedRecipes);
          setSearchTime(0.1);
          setNodesVisited(10);
        } else {
          setRecipes([]);
          setSearchTime(0);
          setNodesVisited(0);
        }
      } catch (localError) {
        console.error('Error loading local data:', localError);
        setRecipes([]);
        setSearchTime(0);
        setNodesVisited(0);
      }
    }
  };

  return (
    <div>
      <h1 className="text-3xl font-bold text-white mb-6">Little Alchemy 2 Recipe Finder</h1>
      <AlgorithmSelector
        algorithm={algorithm}
        setAlgorithm={setAlgorithm}
        mode={mode}
        setMode={setMode}
        maxRecipes={maxRecipes}
        setMaxRecipes={setMaxRecipes}
      />
      <SearchBar
        element={element}
        setElement={setElement}
        onSearch={handleSearch}
      />
      <div className="stats">
        <p>Search Time: {searchTime.toFixed(3)} seconds</p>
        <p>Nodes Visited: {nodesVisited}</p>
      </div>
      <RecipeTree recipes={recipes} />
    </div>
  );
}

export default Main;