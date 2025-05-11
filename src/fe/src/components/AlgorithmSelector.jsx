function AlgorithmSelector({ algorithm, setAlgorithm, mode, setMode, maxRecipes, setMaxRecipes }) {
  return (
    <div className="algorithm-selector">
      <div className="flex items-center gap-2">
        <label className="text-dark-green font-medium">Algorithm:</label>
        <select
          value={algorithm}
          onChange={(e) => setAlgorithm(e.target.value)}
          className="p-2 border border-dark-green rounded-md bg-beige focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all"
        >
          <option value="BFS">BFS</option>
          <option value="DFS">DFS</option>
        </select>
      </div>
      <div className="flex items-center gap-2">
        <label className="text-dark-green font-medium">Mode:</label>
        <select
          value={mode} 
          onChange={(e) => setMode(e.target.value)}
          className="p-2 border border-dark-green rounded-md bg-beige focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all"
        >
          <option value="single">Shortest Recipe</option>
          <option value="multiple">Multiple Recipes</option>
        </select>
      </div>
      {mode === 'multiple' && (
        <div className="flex items-center gap-2">
          <label className="text-dark-green font-medium">Max Recipes:</label>
          <input
            type="number"
            value={maxRecipes}
            onChange={(e) => setMaxRecipes(Number(e.target.value))}
            min="1"
            className="w-20 p-2 border border-dark-green rounded-md bg-beige focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all"
          />
        </div>
      )}
    </div>
  );
}

export default AlgorithmSelector;