import Tree from 'react-d3-tree';
import { useRef, useEffect, useState } from 'react';


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
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });
  const containerRef = useRef(null);
  const treeContainerRef = useRef(null);
  
  useEffect(() => {
    if (treeContainerRef.current) {
      const { width, height } = treeContainerRef.current.getBoundingClientRect();
      setDimensions({ width, height });
    }
  }, []);

  if (!recipes || recipes.length === 0) {
    return <p className="text-dark-green">No recipes found.</p>;
  }

  // Custom node renderer with modern styling
  const renderCustomNode = (rd3tProps) => {
    const { nodeDatum } = rd3tProps;
    const imageUrl = getImageUrl(nodeDatum.name);
    
    return (
      <g>
        <circle 
          r="32" 
          fill="#f9fafb"
          stroke="#2F855A" 
          strokeWidth="2"
          filter="drop-shadow(0px 4px 6px rgba(0, 0, 0, 0.1))"
        />
        {imageUrl && (
          <image
            href={imageUrl}
            x="-25"
            y="-25"
            width="50"
            height="50"
            className="tree-flip"
          />
        )}
        <text 
          dy="50px" 
          textAnchor="middle" 
          className=" tree-node-text"
        >
          {nodeDatum.name}
        </text>
      </g>
    );
  };

  return (
    <div className="recipe-tree" ref={containerRef}>
      {recipes.map((recipe, idx) => {
        const centerX = dimensions.width / 2;
        
        return (
          <div key={`${recipe.name}-${idx}`} className="rounded-lg p-6 shadow-sm">
            <h3 >{recipe.name}</h3>
            <div 
              ref={treeContainerRef}
              className="w-full h-[750px] overflow-hidden rounded-lg border border-green-200 bg-gray-50 tree-flip"
            >
              {(
                <Tree
                  data={recipe}
                  orientation="vertical"
                  translate={{ x: centerX, y: 80 }}
                  initialDepth={999}
                  separation={{ siblings: 1.5, nonSiblings: 2 }}
                  zoomable={true}
                  scaleExtent={{ min: 0.5, max: 2 }}
                  nodeSize={{ x: 120, y: 120 }}
                  pathClassFunc={() => 'stroke-green-600 stroke-2'}
                  renderCustomNodeElement={renderCustomNode}
                  enableLegacyTransitions={false}
                  transitionDuration={800}
                />
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
}

export default RecipeTree;