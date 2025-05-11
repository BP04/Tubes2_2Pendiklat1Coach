import { useState } from 'react';
import elementsData from '../../public/elements.json';

function SearchBar({ element, setElement, onSearch }) {
  const [input, setInput] = useState(element);

  // Sort element names alphabetically
  const elementOptions = elementsData
    .map((item) => item.element)
    .sort((a, b) => a.localeCompare(b));

  const handleSubmit = (e) => {
    e.preventDefault();
    setElement(input);
    onSearch();
  };

  const handleSelect = (e) => {
    const selectedElement = e.target.value;
    setInput(selectedElement);
    setElement(selectedElement);
  };

  return (
    <div className="search-bar">
      <select
        value={input}
        onChange={handleSelect}
        className="w-48 p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all"
      >
        <option value="">Select an element</option>
        {elementOptions.map((elem, idx) => (
          <option key={idx} value={elem}>
            {elem}
          </option>
        ))}
      </select>
      <input
        type="text"
        value={input}
        onChange={(e) => {
          setInput(e.target.value);
          setElement(e.target.value);
        }}
        className="w-64 p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-all"
        placeholder="Or type element (e.g., Land)"
      />
      <button
        onClick={handleSubmit}
        className="px-4 py-2 bg-yellow-600 text-dark-green rounded-md hover:bg-yellow-700 transition-all"
      >
        Search
      </button>
    </div>
  );
}

export default SearchBar;