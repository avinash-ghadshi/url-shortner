import React, { useState } from 'react';
import './../css/home.css';
import titleimage from './../img/title.png';

const Home = () => {
   const [longUrl, setLongUrl] = useState('');
   const [result, setResult] = useState('');
   const [expirationDays, setExpirationDays] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log('Long URL:', longUrl);
        console.log('Expiration Days:', expirationDays);
        // Handle form submission logic here
    };

    const formStyle = {
      display: 'flex',
      alignItems: 'center',
      gap: '20px'
  };

  const fieldContainerStyle = {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'flex-start',
  };

  const inputStyle = {
      padding: '10px',
      fontSize: '16px',
      border: '1px solid #ccc',
      borderRadius: '4px',
      resize: 'none',
  };

    return (
      <div>
      <h2><center><img src={titleimage}></img></center></h2>
      <div className="container">

      <h2>Test Short URL </h2>
      <div className="input-group">
           <input type="text" placeholder="http://localhost:3000/url.short" />
          <button>Visit</button>
      </div>
      <hr />
      <h2>Generate short URL</h2>
      <div className="input-group generate">
      <form onSubmit={handleSubmit} style={formStyle}>
                <div style={fieldContainerStyle}>
                    <label htmlFor="longUrl">Long URL:</label>
                    <textarea
                        id="longUrl"
                        rows={5}
                        cols={50}
                        value={longUrl}
                        onChange={(e) => setLongUrl(e.target.value)}
                        placeholder="Enter the long URL here..."
                        style={{ ...inputStyle, width: '300px', maxHeight: '100px' }}
                        required
                    />
                </div>
                <div style={fieldContainerStyle}>
                    <label htmlFor="expirationDays">Expiry (in min):</label>
                    <input
                        id="expirationDays"
                        type="number"
                        min={1}
                        value={expirationDays}
                        onChange={(e) => setExpirationDays(e.target.value)}
                        placeholder="Expiry in mins"
                        style={{ ...inputStyle, width: '100px' }}
                        required
                    />
                </div>
                <button type="submit"> Generate </button>
                <span className="arrow result">&gt;</span>
                <div style={fieldContainerStyle}>
                    <label htmlFor="result">Result:</label>
                    <textarea
                        id="result"
                        rows={5}
                        cols={50}
                        value={result}
                        style={{ ...inputStyle, width: '300px', maxHeight: '100px' }}
                        required
                    />
                </div>
                <div style={fieldContainerStyle}>
                <button className='btn green'> Copy </button>
                <button className='btn red'> Delete </button>
                </div>
            </form>
        </div>
      </div>
      </div>
    );
};

export default Home;
