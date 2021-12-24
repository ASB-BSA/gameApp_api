import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { ChakraProvider } from "@chakra-ui/react"
import { RecoilRoot } from 'recoil';
import axios from 'axios';
import { BrowserRouter } from 'react-router-dom';
import theme from './config/ChakraConfig';

// https://api.localhost/api/v1/admin/
// https://api.game-boomin.net/api/v1/admin/

// axios.defaults.baseURL = 'https://api.game-boomin.net/api/v1/admin/'
axios.defaults.baseURL = 'https://api.localhost/api/v1/admin/'
axios.defaults.withCredentials = true

ReactDOM.render(
  <React.StrictMode>
    <RecoilRoot>
      <BrowserRouter>
        <ChakraProvider theme={theme}>
          <App />
        </ChakraProvider>
      </BrowserRouter>
    </RecoilRoot>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
