import { useState, useEffect } from 'react';
// import elementsData from '../../public/elements.json';

async function fetchElements() {
  try {
    const response = await fetch('http://localhost:8080/elements'); 
    if (!response.ok) {
      throw new Error('Failed to fetch elements');
    }
    const data = await response.json();
    console.log('Fetched elements:', data);
    return data;
  } catch (error) {
    console.error('Error fetching elements:', error);
    return [];
  }
}

function SearchBar({ element, setElement, onSearch }) {
  const [input, setInput] = useState(element);
  const [elementOptions, setElementOptions] = useState([]);

  useEffect(() => {
    const loadElements = async () => {
      const data = await fetchElements();
      const sorted = data
        .map((item) => item.element)
        .sort((a, b) => a.localeCompare(b));
      setElementOptions(sorted);
    };

    loadElements();
  }, []);

  const handleSubmit = async (e) => {
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