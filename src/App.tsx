import React, {useEffect} from 'react';
import './App.css';
import axios from 'axios';

function App() {
  // axios request on /api/line-count
  const [data, setData] = React.useState(null);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState(null);

  useEffect(() => {
    axios.get('/api/line-count')
      .then(res => {
        setData(res.data);
        setLoading(false);
      })
      .catch(err => {
        setError(err);
        setLoading(false);
      });
  }, []);

  return (
    <div className="App">
      <pre>
        {loading && 'Loading...'}
        {data && JSON.stringify(data, null, 2)}
        {error && JSON.stringify(error, null, 2)}
      </pre>
    </div>
  );
}

export default App;
