import React from 'react';
import './App.css';
import axios from 'axios';

function App() {
  const [data, setData] = React.useState(null);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState(null);

  const [value, setValue] = React.useState('');

  const onSubmit = () => {
    setData(null);
    setError(null);
    setLoading(true);
    axios.get('/api/line-count', {
        params: {
          url: value
        }
    })
      .then(res => {
        setData(res.data);
        setLoading(false);
      })
      .catch(err => {
        setError(err);
        setLoading(false);
      });
  }

  return (
    <div className="App">
      <input type="text" placeholder="git repository..." value={value} onChange={e => setValue(e.target.value)}/>
      <button type="submit" onClick={onSubmit}>Submit</button>
      <pre>
        {loading && 'Loading...'}
        {data && JSON.stringify(data, null, 2)}
        {error && JSON.stringify(error, null, 2)}
      </pre>
    </div>
  );
}

export default App;
