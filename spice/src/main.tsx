import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { ThemeProvider } from '@emotion/react';
import theme from './theme';
import { CssBaseline } from '@mui/material';
import { BrowserRouter } from 'react-router-dom';
import { createConnectTransport } from '@connectrpc/connect-web';
import { TransportContext } from 'TransportContext';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <TransportContext.Provider
          value={createConnectTransport({
            baseUrl: 'http://localhost:21337',
            useBinaryFormat: false,
            credentials: 'same-origin',
            useHttpGet: false,
          })}
        >
          <App />
        </TransportContext.Provider>
      </ThemeProvider>
    </BrowserRouter>
  </React.StrictMode>
);
