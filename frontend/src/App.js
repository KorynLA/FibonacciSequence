import React, { useRef, useState } from 'react';
//import ViewSequence from './components/FibonacciSequenceDisplay';
import logo from './logo.svg';
import './App.css';

function App() {
  const [digit, setDigit] = useState("");
  const [sequence, setSequence] = useState([]);

  function handleSubmit(event) {
    event.preventDefault();
    fetch("/api/fibonacci/"+digit)
      .then(response => response.json())
      .then(data => setSequence({data}))
      .catch(error => {
        console.log(error);
      });
  }
  
  function viewSequence() {
    if((sequence.data) != undefined) {
      const sequenceValues = sequence.data.map((value) =>
        <li>{value}</li>
      );
    return (
      <ul>{sequenceValues}</ul>
    );
    }
  }

  return (
    <div className="App">
      <h1>Find the Fibonacci Sequence</h1>
      {viewSequence()}
      <form onSubmit={handleSubmit} onChange={e => setDigit(e.target.value)}>
        <label> Digits 
          <input type="text" name="digit" />
        </label>
        <input type="submit" value="Submit" />
      </form>
    </div>
  );
}

export default App;
