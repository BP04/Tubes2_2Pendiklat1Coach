import Tree from 'react-d3-tree';

function RecipeTree({ recipes }) {
  if (!recipes || recipes.length === 0) {
    return <p className="text-dark-green text-center">No recipes found.</p>;
  }

  return (
    <div className="recipe-tree">
      {recipes.map((recipe, idx) => (
        <div key={idx} className="mb-8">
          <h3 className="text-lg font-semibold text-dark-green mb-4">Recipe {idx + 1}</h3>
          <div className="w-full h-[500px] overflow-auto">
            <Tree
              data={recipe}
              orientation="vertical"
              translate={{ x: 200, y: 450 }}
              zoomable
              scaleExtent={{ min: 0.5, max: 2 }}
              svgClassName="tree-flip"
              renderCustomNodeElement={(rd3tProps) => (
                <g>a
                  <circle r={10} fill="#F8C249" />
                  <text dy=".31em" x={15} className="tree-node-text">
                    {rd3tProps.nodeDatum.name}
                  </text>
                </g>
              )}
            />
          </div>
        </div>
      ))}
    </div>
  );
}

export default RecipeTree;