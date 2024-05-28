import { createConnectTransport } from '@connectrpc/connect-web';
import { createContext } from 'react';

export const TransportContext = createContext(
  createConnectTransport({
    baseUrl: 'http://localhost:21337',
  })
);
