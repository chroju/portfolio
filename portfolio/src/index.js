import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Blogs from './Blogs';
import * as serviceWorker from './serviceWorker';

ReactDOM.render(
  <React.StrictMode>
    <Blogs rssURL="https://chroju.github.io/atom.xml" />
  </React.StrictMode>,
  document.getElementById('techblog')
);
ReactDOM.render(
  <React.StrictMode>
    <Blogs rssURL="https://chroju.hatenablog.jp/rss" />
  </React.StrictMode>,
  document.getElementById('hobbyblog')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
