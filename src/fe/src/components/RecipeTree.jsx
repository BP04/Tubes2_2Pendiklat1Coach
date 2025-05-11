import Tree from 'react-d3-tree';

const imageModules = import.meta.glob('../../public/icons/*.svg', { eager: true });

console.log('imageModules:', Object.keys(imageModules));
console.log(
  'imageModules details:',
  Object.fromEntries(
    Object.entries(imageModules).map(([path, module]) => [path, module.default])
  )
);

const getImageUrl = (name) => {
  const imagePath = Object.keys(imageModules).find((path) =>
    path.includes(`${name}.svg`)
  );
  return imagePath ? imageModules[imagePath].default : null;
};


function RecipeTree({ recipes }) {
  if (!recipes || recipes.length === 0) {
    return <p className="text-dark-green text-center">No recipes found.</p>;
  }

  return (
    <div className="recipe-tree">
      {recipes.map((recipe, idx) => (
         <div key={`${recipe.name}-${idx}`} className="mb-8">
          <h3 className="text-lg text-dark-green mb-4">{recipe.name}</h3>
          <div className="w-full h-[500px] overflow-auto">
            <Tree
              data={recipe}
              orientation="vertical"
              translate={{ x: 200, y: 450 }}
              zoomable
              scaleExtent={{ min: 0.5, max: 2 }}
              svgClassName="tree-flip"
              renderCustomNodeElement={(rd3tProps) => {
                const imageUrl = getImageUrl(rd3tProps.nodeDatum.name);
                return (
                  <g>
                    {imageUrl && (
                      <image
                        className="tree-flip"
                        href={imageUrl}
                        x="-30"
                        y="-30"
                        width="60"
                        height="60"
                      />
                    )}
                    <text dy="40px" textAnchor='middle' className="tree-node-text">
                      {rd3tProps.nodeDatum.name}
                    </text>
                  </g>
                );
              }}
            />
          </div>
        </div>
      ))}
    </div>
  );
}

export default RecipeTree;